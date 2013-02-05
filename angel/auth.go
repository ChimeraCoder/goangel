package angel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AngelClient struct {
	Client_id     string
	Client_secret string
	Access_token  string
}

//Generate a URI which users should be redirected to in order to authorize the application
func (c AngelClient) AuthorizeUri() string {
	return fmt.Sprintf("https://angel.co/api/oauth/authorize?client_id=%s&response_type=code", c.Client_id)
}

func (c AngelClient) RequestAccessToken(code string) (access_token string, err error) {
	v := url.Values{}
	v.Set("client_id", c.Client_id)
	v.Set("client_secret", c.Client_secret)
	v.Set("code", code)
	v.Set("grant_type", "authorization_code")
	response, err := http.PostForm("https://angel.co/api/oauth/token", v)
	if err != nil {
		return
	}
	bts, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	result := make(map[string]interface{})
	json.Unmarshal(bts, &result)
	//Some Oauth2 implementations return this as a non-string
	access_token, ok := result["access_token"].(string)
	if !ok {
		err = fmt.Errorf("error converting access token %v to string", result["access_token"])
	}
	return
}
