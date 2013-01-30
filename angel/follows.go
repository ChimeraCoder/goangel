/* Implement the endpoints documented at 
https://angel.co/api/spec/follows
*/
package angel

import (
    "fmt"
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
func QueryUsersFollowingStartups(user_id int64) (startups []Startup, err error) {

	var batch_response StartupsBatchResponse
	endpoint := fmt.Sprintf("/users/%d/following", user_id)
    err = execQueryThrottled(endpoint, map[string]string{"type" : "startup"}, &batch_response)
    startups = batch_response.Startups
    return 
}
