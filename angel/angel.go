package angel

import (
    "fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	UserId    = iota
	StartupId = iota
)

const (
    GET = iota
    HEAD = iota
    POST = iota
    PUT = iota
)

//For now, assume we already have the access token somehow 
const API_BASE = "https://api.angel.co/1"

const SECONDS_PER_QUERY = 10 //By default, execute at most one query every ten seconds
//Set to 0 to turn off throttling

var queryQueue = make(chan QueryChan, 10)

type QueryChan struct {
	endpoint_path string
    method int
	keyVals       map[string]string
	response_ch   chan QueryResponse
}

type QueryResponse struct {
	result []byte
	err    error
}

func init() {
	go throttledQuery(queryQueue)
}

//Issue a GET request to the specified endpoint
func GetQuery(endpoint_path string, keyVals map[string]string) ([]byte, error) {

	endpoint_url := API_BASE + endpoint_path

	v := url.Values{}

	for key, val := range keyVals {
		v.Set(key, val)
	}

	log.Printf("Querying %s", endpoint_url+"?"+v.Encode())
	resp, err := http.Get(endpoint_url + "?" + v.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body, err
}

//Issue a POST request to the specified endpoint
func PostQuery(endpoint_path string, keyVals map[string]string) ([]byte, error) {

	endpoint_url := API_BASE + endpoint_path

	v := url.Values{}

	for key, val := range keyVals {
		v.Set(key, val)
	}

	log.Printf("Querying %s", endpoint_url+"?"+v.Encode())
	resp, err := http.PostForm(endpoint_url, v)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body, err
}

func Query(endpoint_path string, method int, keyVals map[string]string) (bts []byte, err error) {
    switch method{
    case GET:
        bts, err = GetQuery(endpoint_path, keyVals)
    case POST:
        bts, err = PostQuery(endpoint_path, keyVals)
    default:
        err = fmt.Errorf("method not supported or not yet implemented")
    }
    return
}

//Execute a query that will automatically be throttled
func throttledQuery(queryQueue chan QueryChan) {
	for q := range queryQueue {

		endpoint_path := q.endpoint_path
        method := q.method
		keyVals := q.keyVals
		response_ch := q.response_ch
		result, err := Query(endpoint_path, method, keyVals)
		response_ch <- struct {
			result []byte
			err    error
		}{result, err}

		time.Sleep(SECONDS_PER_QUERY)
	}
}

func execQueryThrottled(endpoint string, method int, vals map[string]string, result interface{}) error {
	resp_ch := make(chan QueryResponse)
	queryQueue <- QueryChan{endpoint, method, vals, resp_ch}
	r := <-resp_ch
	res := r.result
	if err := r.err; err != nil {
		return err
	}

	//Result should be a pointer to the desired struct
	if err := json.Unmarshal(res, result); err != nil {
		return err
	}
	return nil
}


func (c AngelClient) execQueryThrottled(endpoint string, method int, vals map[string]string, result interface{}) error {
    vals["access_token"] = c.Access_token
    return execQueryThrottled(endpoint, method, vals, result)

}
