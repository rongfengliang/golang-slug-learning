package id

import (
	"github.com/gosimple/slug"
)

// GenerateSlugID generate one slug id
func GenerateSlugID(titile string) string {
	return slug.Make(titile)
}

// GenerateSlugIDLang generate one slug id
func GenerateSlugIDLang(titile string, lang string) string {
	return slug.MakeLang(titile, lang)
}
