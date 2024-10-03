package router

type MathBData struct {
	Class string
	Code  string
	Title string
	Name  string
	Error string
	Date  string
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
