package main

import (
	"appdemo/pkg/id"
	"log"
)

func main() {
	slug := id.GenerateSlugID("荣锋亮的测试博客")
	log.Println(slug)
}
