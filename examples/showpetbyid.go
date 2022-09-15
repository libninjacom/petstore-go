package main

import (
	"fmt"
	"petstore"
)

func main() {
	client := petstore.NewClientFromEnv()
	petId := 1
	res, err := client.ShowPetById(petId)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
