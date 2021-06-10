package middleware

import "github.com/bcowtech/structproto"

var _ structproto.TagResolver = NoneTagResolver

func UrlTagResolver(fieldname, token string) (*structproto.Tag, error) {
	var tag *structproto.Tag
	if token != "-" {
		tag = &structproto.Tag{
			Name: token,
		}
	}
	return tag, nil
}
