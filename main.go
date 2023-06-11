package main

import (
	"clean-app/config"
	"clean-app/rest"
)

func main() { 
    config.InitConfig()
    s := rest.NewServer()
    s.Run()
}