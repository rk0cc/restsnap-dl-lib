package resthttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
func (reqI requestInfo) doRequest(a *[]Auth) string {
	_ = SetContentType(reqI.Content)
	_ = SetCORS(reqI.CrossOrigin)
	if a != nil { /* If authorization is provided */
		for _, authmeta := range *a {
			if b, isBA := authmeta.(basicAuth); isBA {
				b.getBAURL(reqI)
			} else if t, isTA := authmeta.(tokenAuth); isTA {
				t.assignTokenHeader(false)
			} else {
				panic("This is not a valid Auth interface")
			}
		}
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", reqI.URL, nil)
	if err != nil {
		panic("Unable do GET request")
	}
	for _, h := range headerSet {
		req.Header.Set(h.name, h.value)
	}
	res, err := client.Do(req)
	if err != nil {
		panic("No response of target website")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Fail to read context of the website")
	}
	return string(body)
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

func getReqInfo(url, content, cors string) requestInfo {
	return requestInfo{
		URL:         url,
		Content:     content,
		CrossOrigin: cors,
	}
}

//SendRequest : Send request with no authorize method
func SendRequest(url, content, cors string) string {
	i := getReqInfo(url, content, cors)
	return i.doRequest(nil)
}

//SendRequestWithBasicAuth : Send request with username and password
func SendRequestWithBasicAuth(url, content, cors, uname, pwd string) string {
	ba := basicAuth{
		username: uname,
		password: pwd,
	}
	reqBa := &[]Auth{ba}
	i := getReqInfo(url, content, cors)
	return i.doRequest(reqBa)
}

//SendRequestWithTokenAuth : Send request with token
func SendRequestWithTokenAuth(url, content, cors, tokenHeader, tokenStr string) string {
	ta := tokenAuth{
		headerName: tokenHeader,
		token:      tokenStr,
	}
	reqTa := &[]Auth{ta}
	i := getReqInfo(url, content, cors)
	return i.doRequest(reqTa)
}

//SendRequestWithBasicAndToken : Send request with username, password and token
func SendRequestWithBasicAndToken(url, content, cors, uname, pwd, tokenHeader, tokenStr string) string {
	ba := basicAuth{
		username: uname,
		password: pwd,
	}
	ta := tokenAuth{
		headerName: tokenHeader,
		token:      tokenStr,
	}
	req := &[]Auth{ba, ta}
	i := getReqInfo(url, content, cors)
	return i.doRequest(req)
}
