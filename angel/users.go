package angel
import (
    "log"
    "fmt"
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

	var roles []StartupRole
	roles_bts, err := json.Marshal(res["startup_roles"])
	if err != nil {
	}
	if err := json.Unmarshal(roles_bts, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}
