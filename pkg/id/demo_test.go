// this library just for test slug genreate id

package id_test

import (
	"log"

	"github.com/rongfengliang/golang-slug-learning/v2/pkg/id"
)

//  example for generate id
func ExampleGenerateSlugID() {
	slug := id.GenerateSlugID("荣锋亮的测试博客")
	log.Println(slug)
}

//  example for generate id
func ExampleGenerateSlugIDLang() {
	slug := id.GenerateSlugIDLang("荣锋亮的测试博客", "en")
	log.Println(slug)
}
