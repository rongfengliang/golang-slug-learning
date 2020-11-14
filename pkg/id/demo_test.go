// this library just for test slug genreate id

package id_test

import "log"

//  example for generate id
func Example_GenerateSlugID() {
	slug := GenerateSlugID("荣锋亮的测试博客")
	log.Println(slug)
}

//  example for generate id
func Example_GenerateSlugIDLang() {
	slug := GenerateSlugIDLang("荣锋亮的测试博客", "en")
	log.Println(slug)
}
