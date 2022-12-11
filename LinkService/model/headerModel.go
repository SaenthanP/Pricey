package model

type Header struct {
    UserId int `header:"UserId" binding:"required"`
}