package main

import (
	"fmt"

	"github.com/cairomedeiros/todo-list/config"
	"github.com/cairomedeiros/todo-list/router"
)

func main() {

	//initialize configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	router.Initialize()
}
