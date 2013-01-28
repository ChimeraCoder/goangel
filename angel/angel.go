package angel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
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
	result map[string]interface{}
	err    error
}

func init() {
	go throttledQuery(queryQueue)
}

func Query(endpoint_path string, keyVals map[string]string) (map[string]interface{}, error) {

	endpoint_url := API_BASE + endpoint_path

	v := url.Values{}

	for key, val := range keyVals {
		v.Set(key, val)
	}

	resp, err := http.Get(endpoint_url + "?" + v.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	return result, err
}

//Execute a query that will automatically be throttled
func throttledQuery(queryQueue chan QueryChan) {
	for q := range queryQueue {

		endpoint_path := q.endpoint_path
		keyVals := q.keyVals
		response_ch := q.response_ch
		result, err := Query(endpoint_path, keyVals)
		response_ch <- struct {
			result map[string]interface{}
			err    error
		}{result, err}

		time.Sleep(SECONDS_PER_QUERY)
	}
}

//Query /users/search for a user with the specified slug
//TODO fix this to search for emails as well
func QueryUsersSearch(slug string) (*AngelUser, error) {
	resp_ch := make(chan QueryResponse)
	queryQueue <- QueryChan{"/users/search", map[string]string{"slug": slug}, resp_ch}
	r := <-resp_ch
	res := r.result
	if err := r.err; err != nil {
		return nil, err
	}

	var user AngelUser
	users_bts, err := json.Marshal(res)
	if err != nil {
	}
	if err := json.Unmarshal(users_bts, &user); err != nil {
		log.Print(string(users_bts))
		return nil, err
	}
	return &user, nil
}

//Query /startup_roles for all startup roles associated with the user with the given id
func QueryStartupRoles(id int64, id_type int) ([]StartupRole, error) {
	resp_ch := make(chan QueryResponse)

	switch id_type {
	case UserId:
		{
			queryQueue <- QueryChan{"/startup_roles", map[string]string{"user_id": strconv.FormatInt(id, 10)}, resp_ch}
		}
	case StartupId:
		{
			queryQueue <- QueryChan{"/startup_roles", map[string]string{"startup_id": strconv.FormatInt(id, 10)}, resp_ch}
		}
	default:
		return nil, fmt.Errorf("invalid id_type provided")
	}
	r := <-resp_ch
	res := r.result
	if err := r.err; err != nil {
		return nil, err
	}

	roles_array := res["startup_roles"].([]interface{})

	var roles []StartupRole
	roles_bts, err := json.Marshal(roles_array)
	if err != nil {
		log.Print("Woah, error occured while marshalling")
		panic(err)
	}
	if err := json.Unmarshal(roles_bts, &roles); err != nil {
		return nil, err
	}
	return roles, err
}
