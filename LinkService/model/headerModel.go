package model

type Header struct {
    UserId string `header:"UserId" binding:"required"`
}