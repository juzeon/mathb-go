package router

import "html/template"

type MathBData struct {
	Class template.HTMLAttr
	Code  string
	Title template.HTML
	Name  template.HTML
	Error template.HTML
	Date  template.HTML
}

type CreatePostData struct {
	Code  string `form:"code" binding:"required"`
	Title string `form:"title"`
	Name  string `form:"name"`
}

type MathBError struct {
	StatusCode   int
	ReasonPhrase string
}
