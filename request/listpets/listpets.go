package listpets

import (
	"petstore/model"
)

func WithLimit(limit int) RequestOption {
	return func(r *Request) {
		r.Limit = &limit
	}
}

type RequestOption = func(l *Request)
type Request struct {
	Limit *int
}
type Response model.Pets
