package main

import (
	"gotify/gotify"
	"gotify/testService"
)

func main() {
	s := new(testService.TestService)
	gotify.NewService("msg", "proA", s)
}
