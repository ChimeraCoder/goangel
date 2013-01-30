package angel

import (
    "log"
    "encoding/json"
)

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


