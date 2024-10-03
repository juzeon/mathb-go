package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"html"
	"mathb-go/db"
	"net/http"
	"time"
)

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
		paste := db.Paste{
			UUID:      uuid.New().String(),
			Title:     html.EscapeString(postData.Title),
			Name:      html.EscapeString(postData.Name),
			Code:      html.EscapeString(postData.Code),
			CreatedAt: time.Time{},
		}
		db.PasteTx.MustCreate(&paste)
		c.Redirect(http.StatusTemporaryRedirect, "/"+paste.UUID)
	})
	engine.GET("/", func(c *gin.Context) {
		c.HTML(200, "mathb.html", MathBData{})
	})
	engine.GET("/:uuid", getPaste)
	engine.POST("/:uuid", getPaste)
}
func getPaste(c *gin.Context) {
	paste, err := db.PasteTx.FindOne("uuid=?", c.Param("uuid"))
	if err != nil {
		c.HTML(500, "error.html", MathBError{
			StatusCode:   500,
			ReasonPhrase: err.Error(),
		})
		return
	}
	if paste == nil {
		c.HTML(404, "error.html", MathBError{
			StatusCode:   404,
			ReasonPhrase: "Not Found",
		})
		return
	}
	c.HTML(200, "mathb.html", MathBData{
		Class: "class='post'",
		Code:  paste.Code,
		Title: paste.Title,
		Name:  paste.Name,
		Error: "",
		Date:  paste.CreatedAt.Format(time.DateTime),
	})
}
