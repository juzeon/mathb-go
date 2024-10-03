package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/imroc/req/v3"
	"html"
	"html/template"
	"log/slog"
	"mathb-go/conf"
	"mathb-go/db"
	"time"
)

var client *req.Client

func init() {
	client = req.NewClient().SetTimeout(10 * time.Second)
}
func RegisterRouters(engine *gin.Engine) {
	engine.Static("/css", "web/css")
	engine.Static("/js", "web/js")
	engine.Static("/img", "web/img")
	engine.StaticFile("/favicon.ico", "web/img/favicon.ico")
	engine.StaticFile("/favicon.png", "web/img/favicon.png")
	engine.LoadHTMLGlob("web/html/*")
	engine.POST("/", func(c *gin.Context) {
		var postData CreatePostData
		if err := c.ShouldBind(&postData); err != nil {
			c.HTML(400, "error.html", MathBError{
				StatusCode:   400,
				ReasonPhrase: err.Error(),
			})
			return
		}
		var recaptchaResult RecaptchaResult
		_, err := client.R().SetSuccessResult(&recaptchaResult).SetFormData(map[string]string{
			"secret":   conf.Conf.RecaptchaSecret,
			"response": postData.RecaptchaResponse,
		}).Post("https://www.google.com/recaptcha/api/siteverify")
		if err != nil {
			c.HTML(500, "error.html", MathBError{
				StatusCode:   500,
				ReasonPhrase: err.Error(),
			})
			return
		}
		if !recaptchaResult.Success || recaptchaResult.Score < 0.5 {
			c.HTML(400, "error.html", MathBError{
				StatusCode:   400,
				ReasonPhrase: "reCAPTCHA Error; Score: " + fmt.Sprintf("%v", recaptchaResult.Score),
			})
			return
		}
		slog.Info("recaptcha", "score", recaptchaResult.Score)
		paste := db.Paste{
			UUID:      uuid.New().String(),
			Title:     html.EscapeString(postData.Title),
			Name:      html.EscapeString(postData.Name),
			Code:      postData.Code,
			CreatedAt: time.Time{},
		}
		db.PasteTx.MustCreate(&paste)
		c.Redirect(302, "/"+paste.UUID)
	})
	engine.GET("/", func(c *gin.Context) {
		c.HTML(200, "mathb.html", MathBData{
			Sitekey: conf.Conf.RecaptchaSitekey,
		})
	})
	engine.GET("/:uuid", getPaste)
	engine.POST("/:uuid", getPaste)
}
func getPaste(c *gin.Context) {
	paste := db.PasteTx.MustFindOne("uuid=?", c.Param("uuid"))
	if paste == nil {
		c.HTML(404, "error.html", MathBError{
			StatusCode:   404,
			ReasonPhrase: "Not Found",
		})
		return
	}
	c.HTML(200, "mathb.html", MathBData{
		Class:   "class='post'",
		Code:    paste.Code,
		Title:   template.HTML(paste.Title),
		Name:    template.HTML(paste.Name),
		Error:   "",
		Date:    template.HTML(paste.CreatedAt.Format(time.DateTime)),
		Sitekey: conf.Conf.RecaptchaSitekey,
	})
}
