// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

/*
Package slug generate slug from unicode string, URL-friendly slugify with
multiple languages support.

Example:

	package main

	import (
		"log"

		"github.com/rongfengliang/golang-slug-learning/pkg/id"
	)
	func main() {
		slug := id.GenerateSlugID("荣锋亮的测试博客")
		log.Println(slug)
	}

*/
package main
