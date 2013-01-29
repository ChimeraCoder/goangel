/* Implement the endpoints documented at 
https://angel.co/api/spec/follows
*/
package angel

import (
    "fmt"
    "encoding/json"
)


//Query /users/:id/followers for a user's followers
//TODO implement proper pagination
func QueryUsersFollowers(user_id int64) ([]AngelUser, error) {
	return queryFollowersAux(user_id, UserId)
}

//Query /startups/:id/followers for a user's followers
//TODO implement proper pagination
func QueryStartupsFollowers(user_id int64) ([]AngelUser, error) {
	return queryFollowersAux(user_id, StartupId)
}

//Auxiliary function used for /users/:id/followers and /startups/:id/followers
func queryFollowersAux(angel_id int64, entity int) ([]AngelUser, error) {
	var endpoint string
	switch entity {
	case UserId:
		endpoint = fmt.Sprintf("/users/%d/followers", angel_id)
	case StartupId:
		endpoint = fmt.Sprintf("/startups/%d/followers", angel_id)
	default:
		return nil, fmt.Errorf("invalid entity provided")
	}
	resp_ch := make(chan QueryResponse)
	queryQueue <- QueryChan{endpoint, map[string]string{}, resp_ch}
	r := <-resp_ch
	res := r.result
	if err := r.err; err != nil {
		return nil, err
	}

	var batch_response UsersBatchResponse
	resp_bts, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(resp_bts, &batch_response); err != nil {
		return nil, err
	}
	return batch_response.Users, nil
}

//Query /users/:id/following for a user's followers (return users only)
//TODO implement proper pagination
func QueryUsersFollowingUsers(user_id int64) ([]AngelUser, error) {

	endpoint := fmt.Sprintf("/users/%d/following", user_id)
	resp_ch := make(chan QueryResponse)

	queryQueue <- QueryChan{endpoint, map[string]string{"type": "user"}, resp_ch}

	r := <-resp_ch
	res := r.result
	if err := r.err; err != nil {
		return nil, err
	}

	var batch_response UsersBatchResponse
	resp_bts, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(resp_bts, &batch_response); err != nil {
		return nil, err
	}
	return batch_response.Users, nil
}

//Query /users/:id/following for a user's followers (return startups only)
//TODO implement proper pagination
func QueryUsersFollowingStartups(user_id int64) ([]Startup, error) {

	endpoint := fmt.Sprintf("/users/%d/following", user_id)
	resp_ch := make(chan QueryResponse)

	queryQueue <- QueryChan{endpoint, map[string]string{"type": "startup"}, resp_ch}

	r := <-resp_ch
	res := r.result
	if err := r.err; err != nil {
		return nil, err
	}

	var batch_response StartupsBatchResponse
	resp_bts, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(resp_bts, &batch_response); err != nil {
		return nil, err
	}
	return batch_response.Startups, nil
}
