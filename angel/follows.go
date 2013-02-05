/* Implement the endpoints documented at 
https://angel.co/api/spec/follows
*/
package angel

import (
	"fmt"
	"strings"
)

type IdsResponse struct {
	Total     int64
	Per_page  int64
	Page      int64
	Last_page int64
	Ids       []int64
}

//POST query to /follows to follow the specified user/startup
//Will throw an error (400) if the client's user is already following the specified user/startup
func (c AngelClient) QueryPostFollows(id int64, id_type int) (result interface{}, err error) {
	vals := map[string]string{"id": fmt.Sprintf("%d", id)}
	switch id_type {
	case UserId:
		vals["type"] = "user"
	case StartupId:
		vals["type"] = "startup"
	}
	err = c.execAuthenticatedQueryThrottled("/follows", POST, vals, &result)
	return
}

//DELETE query to /follows to stop following the specified user/startup
func (c AngelClient) QueryDeleteFollows(id int64, id_type int) (result interface{}, err error) {
	vals := map[string]string{"id": fmt.Sprintf("%d", id)}
	switch id_type {
	case UserId:
		vals["type"] = "user"
	case StartupId:
		vals["type"] = "startup"
	}
	err = c.execAuthenticatedQueryThrottled("/follows", DELETE, vals, &result)
	return
}

//Query /follows/batch for the followers of several users
//TODO implement proper pagination
//TODO implement the proper return type here
func (c AngelClient) QueryFollowsBatch(ids ...int64) (results interface{}, err error) {
	id_strings := make([]string, len(ids))
	for i, id := range ids {
		id_strings[i] = fmt.Sprintf("%d", id) //do this more elegantly
	}
	err = c.execQueryThrottled("/follows/batch", GET, map[string]string{"ids": strings.Join(id_strings, ",")}, &results)
	return
}

//Query /users/:id/followers for a user's followers
//TODO implement proper pagination
func (c AngelClient) QueryUsersFollowers(user_id int64) (users []AngelUser, err error) {
	err = c.execQueryThrottled(fmt.Sprintf("/startups/%d/followers", user_id), GET, map[string]string{}, users)
	return
}

//Query /users/:id/followers/ids for a user's followers
//TODO implement proper pagination
func (c AngelClient) QueryUsersFollowersIds(user_id int64) (ids []int64, err error) {

	var tmp IdsResponse

	err = c.execQueryThrottled(fmt.Sprintf("/users/%d/followers/ids", user_id), GET, map[string]string{}, &tmp)
	ids = tmp.Ids
	return
}

//Query /users/:id/following for a user's followers (return users only)
//TODO implement proper pagination
func (c AngelClient) QueryUsersFollowingUsers(user_id int64) (users []AngelUser, err error) {

	var batch_response UsersBatchResponse
	endpoint := fmt.Sprintf("/users/%d/following", user_id)
	err = c.execQueryThrottled(endpoint, GET, map[string]string{}, &batch_response)
	users = batch_response.Users
	return
}

//Query /users/:id/following for a user's followers (return startups only)
//TODO implement proper pagination
func (c AngelClient) QueryUsersFollowingStartups(user_id int64) (startups []Startup, err error) {

	var batch_response StartupsBatchResponse
	endpoint := fmt.Sprintf("/users/%d/following", user_id)
	err = c.execQueryThrottled(endpoint, GET, map[string]string{"type": "startup"}, &batch_response)
	startups = batch_response.Startups
	return
}

//Query /users/:id/following/ids for a user's followers (return users only)
//TODO implement proper pagination
func (c AngelClient) QueryUsersFollowingUsersIds(user_id int64) (ids []int64, err error) {
	var tmp IdsResponse
	endpoint := fmt.Sprintf("/users/%d/following/ids", user_id)
	err = c.execQueryThrottled(endpoint, GET, map[string]string{"type": "user"}, &tmp)
	ids = tmp.Ids
	return
}

//Query /users/:id/following/ids for a user's followers (return startups only)
//TODO implement proper pagination
func (c AngelClient) QueryUsersFollowingStartupsIds(user_id int64) (ids []int64, err error) {
	var tmp IdsResponse
	endpoint := fmt.Sprintf("/users/%d/following/ids", user_id)
	err = c.execQueryThrottled(endpoint, GET, map[string]string{"type": "startup"}, &tmp)
	ids = tmp.Ids
	return
}

//Query /startups/:id/followers for a startup's followers
//TODO implement proper pagination
func (c AngelClient) QueryStartupsFollowers(user_id int64) (users []AngelUser, err error) {

	err = c.execQueryThrottled(fmt.Sprintf("/users/%d/followers", user_id), GET, map[string]string{}, users)
	return
}

//Query /startups/:id/followers/ids for a startup's followers
//TODO implement proper pagination
func (c AngelClient) QueryStartupsFollowersIds(user_id int64) (ids []int64, err error) {
	var tmp IdsResponse
	err = c.execQueryThrottled(fmt.Sprintf("/users/%d/followers/ids", user_id), GET, map[string]string{}, &tmp)
	ids = tmp.Ids
	return
}
