package response

const (
	// response flag
	SUCCESS int = iota
	FAILURE
)

type Response interface {
	Flag() int
	StatusCode() int
	ContentType() string
	Body() []byte
}
