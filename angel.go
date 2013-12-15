//Package angel provides functions for interacting with the AngelList API
package angel

import (
	"encoding/json"
	"fmt"
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
	GET    = iota
	HEAD   = iota
	POST   = iota
	PUT    = iota
	DELETE = iota
)

const API_BASE = "https://api.angel.co/1"

//SECONDS_PER_QUERY sets the number of seconds that must elapse between queries
//By default, execute at most one query every ten seconds
//Set to 0 to turn off throttling
var SECONDS_PER_QUERY = time.Duration(10 * time.Second)

var queryQueue = make(chan queryChan, 10)

type queryChan struct {
	endpoint_path string
	method        int
	keyVals       map[string]string
	response_ch   chan queryResponse
}

type queryResponse struct {
	result []byte
	err    error
}

func init() {
	go throttledQuery(queryQueue)
}

//GetQuery issues a GET request to the specified endpoint
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

//PostQuery issues a POST request to the specified endpoint
func PostQuery(endpoint_path string, keyVals map[string]string) ([]byte, error) {
	//TODO refactor this
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

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("received http status code %d", resp.StatusCode)
	}
	log.Printf("Received http status code of %d", resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body, err
}

//Query issues a query to the AngelList API
//Like GetQuery and PostQuery, this function exists for the convenience of the automatically-throttled queries
func Query(endpoint_path string, method int, keyVals map[string]string) (bts []byte, err error) {
	switch method {
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
func throttledQuery(queryQueue chan queryChan) {
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

		<-time.After(SECONDS_PER_QUERY)
	}
}

func execQueryThrottled(endpoint string, method int, vals map[string]string, result interface{}) error {
	resp_ch := make(chan queryResponse)
	queryQueue <- queryChan{endpoint, method, vals, resp_ch}
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

//This requires separate two functions because AngelList will return an error if an incorrect or empty access_token is provided even for unauthenticated queries

func (c AngelClient) execAuthenticatedQueryThrottled(endpoint string, method int, vals map[string]string, result interface{}) error {
	vals["access_token"] = c.Access_token
	return execQueryThrottled(endpoint, method, vals, result)

}

func (c AngelClient) execQueryThrottled(endpoint string, method int, vals map[string]string, result interface{}) error {
	return execQueryThrottled(endpoint, method, vals, result)
}
