package main

import (
	"fmt"
	"petstore"
)

func main() {
	client := petstore.NewClientFromEnv()
	res, err := client.CreatePets()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
