package main

import (
	"fmt"
	"go_recipes/repository"
	"log"
)

func main() {
	err := repository.Migrate()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("is it good ?")
}
