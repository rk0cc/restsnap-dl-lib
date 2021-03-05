package resthttp

import "errors"

//Contains HTTP(S) header's name and value
type httpHeader struct {
	name  string
	value string
}

var (
	headerSet []httpHeader
)

//Define HTTP(S) header
func SetHeader(n, v string) error {
	if n == "" {
		//When header's name is ignored
		return errors.New("Header name must be provided")
	}
	headerSet = append(headerSet, httpHeader{n, v})
	return nil //normally will return this
}

//Define content type for making request
func SetContentType(mime string) error {
	return SetHeader("Content-Type", mime)
}
