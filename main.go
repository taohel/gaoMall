package main

import (
	"gaoMall/app"
	"gaoMall/app/api"
	"gaoMall/app/papi"
)

func main() {
	err := app.Init()
	if err != nil {
		panic(err)
	}

	go api.NewServer()
	go papi.NewServer()
	select {}
}
