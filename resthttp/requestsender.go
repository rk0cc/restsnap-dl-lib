package resthttp

import (
	"encoding/json"
	"net/url"
)

//Auth : Getter of authorization
//
//Authorization interface
type Auth interface {
	GetJSON() string //Get JSON formatted authorize data
}

//Authorize by entering username and password
type basicAuth struct {
	username string //Username
	password string //Password
}

//Authorize by using token
type tokenAuth struct {
	headerName string //HTTP header for sending token to API servers
	token      string //Token string
}

//A type that storing current requests
type requestInfo struct {
	URL         string //Target URL
	Content     string //MIME content type
	CrossOrigin string //Set Access-Cross-Origin-Allow
}

func (ba basicAuth) GetJSON() string {
	j, _ := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: ba.username,
		Password: ba.password,
	})
	return string(j)
}

func (ta tokenAuth) GetJSON() string {
	j, _ := json.Marshal(struct {
		Token string `json:"token"`
	}{
		Token: ta.token,
	})
	return string(j)
}

//Making request
func (reqI requestInfo) doRequest(a *Auth) {
	_ = SetContentType(reqI.Content)
	_ = SetCORS(reqI.CrossOrigin)
	if a != nil { /* If authorization is provided */
		var authmeta Auth = *a
		if b, isBA := authmeta.(basicAuth); isBA {
			b.getBAURL(reqI)
		} else if t, isTA := authmeta.(tokenAuth); isTA {
			t.assignTokenHeader(false)
		} else {
			panic("This is not a valid Auth interface")
		}
	}
	//res, err := http.Get(reqI.URL)
}

func (ba basicAuth) getBAURL(reqI requestInfo) {
	urlObj, err := url.Parse(reqI.URL)
	if err != nil {
		panic("This may not a valid URL")
	}
	urlObj.User = url.UserPassword(ba.username, ba.password)
	reqI.URL = urlObj.String()
}

func (ta tokenAuth) assignTokenHeader(allowOverRide bool) {
	_ = SetCustomHeader(ta.headerName, ta.token, allowOverRide)
}
