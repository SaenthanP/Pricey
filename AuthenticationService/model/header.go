package model

type Header struct {
	Token string `header:"token" binding:"required"`
}
