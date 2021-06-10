package response

type responseImpl struct {
	flag         int
	statusCode   int
	contentType  string
	body         []byte
}

func (r *responseImpl) Flag() int {
	return r.flag
}

func (r *responseImpl) StatusCode() int {
	return r.statusCode
}

func (r *responseImpl) ContentType() string {
	return r.contentType
}

func (r *responseImpl) Body() []byte {
	return r.body
}
