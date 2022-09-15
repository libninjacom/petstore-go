package main

import (
	"fmt"
	"petstore"
	"petstore/request/listpets"
)

func main() {
	client := petstore.NewClientFromEnv()
	res, err := client.ListPets(listpets.WithLimit(1))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
