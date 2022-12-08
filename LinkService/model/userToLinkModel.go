package model
import (
	"github.com/google/uuid"
)

type UserToLink struct{
	LinkId	 		  uuid.UUID
	UserId	 		  uuid.UUID 
}