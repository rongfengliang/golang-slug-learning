package main

import (
	"log"

	"github.com/rongfengliang/golang-slug-learning/pkg/id"
)

func main() {
	slug := id.GenerateSlugID("荣锋亮的测试博客")
	log.Println(slug)
}
