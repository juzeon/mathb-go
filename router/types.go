package router

import "html/template"

type MathBData struct {
	Class   template.HTMLAttr
	Code    string
	Title   template.HTML
	Name    template.HTML
	Error   template.HTML
	Date    template.HTML
	Sitekey string
}

type CreatePostData struct {
	Code              string `form:"code" binding:"required"`
	Title             string `form:"title"`
	Name              string `form:"name"`
	RecaptchaResponse string `form:"g-recaptcha-response" binding:"required"`
}

type MathBError struct {
	StatusCode   int
	ReasonPhrase string
}
type RecaptchaResult struct {
	Success     bool    `json:"success"`
	Score       float64 `json:"score"`
	Action      string  `json:"action"`
	ChallengeTs string  `json:"challenge_ts"`
	Hostname    string  `json:"hostname"`
}
