package angel

import (
    "fmt"
    "encoding/json"
)


//Query /users/:id for a user's information given an id


//Query /users/:id/startups for a user's startup roles
func QueryUsersStartups(id int64) ([]StartupRole, error){
    endpoint := fmt.Sprintf("/users/%d/startups", id)
	resp_ch := make(chan QueryResponse)
	queryQueue <- QueryChan{endpoint, map[string]string{}, resp_ch}
	r := <-resp_ch
	res := r.result
	if err := r.err; err != nil {
		return nil, err
	}

    var tmp struct{
        Startup_roles []StartupRole
    }

	if err := json.Unmarshal(res, &tmp); err != nil {
		return nil, err
	}
	return tmp.Startup_roles, nil
}

//Query /users/batch for up to 50 users at a time, givne a comma-separated list of IDs



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
	if err := json.Unmarshal(res, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

