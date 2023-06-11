package main

import "clean-app/rest"

func main() { 
    s := rest.NewServer()
    s.Run()
}