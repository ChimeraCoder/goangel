/* Implement the endpoints documented at 
https://angel.co/api/spec/follows
*/
package angel

import (
	"fmt"
    "strings"
)



//Query /follows/batch for the followers of several users
//TODO implement proper pagination
//TODO implement the proper return type here
func QueryFollowsBatch(ids ...int64) (results interface{}, err error) {
    id_strings := make([]string, len(ids))
    for i, id := range ids{
        id_strings[i] = fmt.Sprintf("%d", id) //do this more elegantly
    }
    err = execQueryThrottled("/follows/batch", map[string]string{"ids" : strings.Join(id_strings, ",")}, &results)
	return
}


//Query /users/:id/followers for a user's followers
//TODO implement proper pagination
func QueryUsersFollowers(user_id int64) (users []AngelUser, err error) {
	err = execQueryThrottled(fmt.Sprintf("/startups/%d/followers", user_id), map[string]string{}, users)
	return
}

//Query /users/:id/followers/ids for a user's followers
//TODO implement proper pagination
func QueryUsersFollowersIds(user_id int64) (ids []int64, err error) {

    var tmp struct {
        Total int64
        Per_page int64
        Page int64
        Last_page int64
        Ids []int64
    }

	err = execQueryThrottled(fmt.Sprintf("/users/%d/followers/ids", user_id), map[string]string{}, &tmp)
    ids = tmp.Ids
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
	err = execQueryThrottled(endpoint, map[string]string{"type": "startup"}, &batch_response)
	startups = batch_response.Startups
	return
}


//Query /users/:id/following/ids for a user's followers (return users only)
//TODO implement proper pagination
//TODO




//Query /users/:id/following/ids for a user's followers (return startups only)
//TODO implement proper pagination
//TODO


//Query /startups/:id/followers for a startup's followers
//TODO implement proper pagination
func QueryStartupsFollowers(user_id int64) (users []AngelUser, err error) {

	err = execQueryThrottled(fmt.Sprintf("/users/%d/followers", user_id), map[string]string{}, users)
	return
}


//Query /startups/:id/followers/ids for a startup's followers
//TODO implement proper pagination
//TODO


