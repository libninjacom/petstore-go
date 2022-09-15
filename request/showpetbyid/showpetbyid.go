package showpetbyid

import (
	"petstore/model"
)

type RequestOption = func(l *Request)
type Request struct {
	PetId int
}
type Response model.Pet
