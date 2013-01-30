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
func QueryUsersFollowers(user_id int64) (users []AngelUser, err error) {
    err = execQueryThrottled(fmt.Sprintf("/startups/%d/followers", user_id), map[string]string{}, users)
    return 
}

//Query /startups/:id/followers for a startup's followers
//TODO implement proper pagination
func QueryStartupsFollowers(user_id int64) (users []AngelUser, err error) {

    err = execQueryThrottled(fmt.Sprintf("/users/%d/followers", user_id), map[string]string{}, users)
    return 
}


//Query /users/:id/following for a user's followers (return users only)
//TODO implement proper pagination
func QueryUsersFollowingUsers(user_id int64) (users []AngelUser, err error) {

	var batch_response UsersBatchResponse
	endpoint := fmt.Sprintf("/users/%d/following", user_id)
    err = execQueryThrottled(endpoint, map[string]string{}, &batch_response)
    users = batch_response.Users
    return
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
