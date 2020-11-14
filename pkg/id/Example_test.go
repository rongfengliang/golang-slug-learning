// apache license

// this library just for test slug genreate id

package id

import "log"

// generate id

//   slug := id.GenerateSlugID("荣锋亮的测试博客")
//	 log.Println(slug)

// ExampleidGenerateSlugID example for generate id
func ExampleidGenerateSlugID() {
	slug := GenerateSlugID("荣锋亮的测试博客")
	log.Println(slug)
}
