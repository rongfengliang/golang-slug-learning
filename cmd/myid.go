package main

import (
	"log"

	"github.com/rongfengliang/golang-slug-learning/v2/pkg/id"
)

func main() {
	slug := id.GenerateSlugID("荣锋亮的测试博客")
	log.Println(slug)
	slug2 := id.GenerateSlugIDLang("荣锋亮的测试博客", "en")
	log.Println(slug2)
}
