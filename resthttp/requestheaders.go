package resthttp

import "errors"

//Contains HTTP(S) header's name and value
type httpHeader struct {
	name  string
	value string
}

//Check the value is defined
type valDefined struct {
	contentType bool //Check 'Content-Type'
}

var (
	headerSet []httpHeader                     //Array for containing HTTP headers
	vd        valDefined   = valDefined{false} //Check common header is defined or not
)

//Define HTTP(S) header
func setHeader(n, v string) error {
	if n == "" {
		//When header's name is ignored
		return errors.New("Header name must be provided")
	}
	headerSet = append(headerSet, httpHeader{n, v})
	return nil //normally will return this
}

//SetContentType : Define content type for making request
func SetContentType(mime string) error {
	if !vd.contentType {
		return errors.New("Content type is defined already")
	}
	vd.contentType = true
	return setHeader("Content-Type", mime)
}

//SetCustomHeader : Allow user define custom headers by theirselves
//
//THIS ACTION MAY CAUSE REQUEST FAILED IF DEFINED INCORRECTLY
func SetCustomHeader(n, v string, override bool) error {
	for _, headers := range headerSet {
		if n == headers.name {
			if override {
				headers.value = v
				return nil
			}
			return errors.New("The header " + n + "has been defined already")
		}
	}
	/* When no header defined */
	return setHeader(n, v)
}
