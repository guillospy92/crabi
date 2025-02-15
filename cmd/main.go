package main

import "github.com/guillospy92/crabi/bootstrap"

func main() {
	app := bootstrap.LoadApplication("resources/application.yaml")
	bootstrap.StartApplication(app)
}
