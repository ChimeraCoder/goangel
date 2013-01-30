package angel

import (
	"io/ioutil"
    "encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	UserId    = iota
	StartupId = iota
)

//For now, assume we already have the access token somehow 
const API_BASE = "https://api.angel.co/1"

const SECONDS_PER_QUERY = 10 //By default, execute at most one query every ten seconds
//Set to 0 to turn off throttling

var queryQueue = make(chan QueryChan, 10)

type QueryChan struct {
	endpoint_path string
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
func Query(endpoint_path string, keyVals map[string]string) ([]byte, error) {

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

//Execute a query that will automatically be throttled
func throttledQuery(queryQueue chan QueryChan) {
	for q := range queryQueue {

		endpoint_path := q.endpoint_path
		keyVals := q.keyVals
		response_ch := q.response_ch
		result, err := Query(endpoint_path, keyVals)
		response_ch <- struct {
			result []byte
			err    error
		}{result, err}

		time.Sleep(SECONDS_PER_QUERY)
	}
}

func execQueryThrottled(endpoint string, vals map[string]string, result interface{}) error {
	resp_ch := make(chan QueryResponse)
	queryQueue <- QueryChan{endpoint, vals, resp_ch}
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
