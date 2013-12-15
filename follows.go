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

//Follow issues a POST query to /follows to follow the user/startup with the specified id
//Will throw an error (400) if the client's user is already following the specified user/startup
func (c AngelClient) Follow(id int64, id_type int) (result interface{}, err error) {
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

//Unfollow issues a DELETE query to /follows to stop following the user/startup with the specified id
func (c AngelClient) Unfollow(id int64, id_type int) (result interface{}, err error) {
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

//QueryFollowsBatch queries follows/batch for the followers of several users
//Returns the follower and followed information
func (c AngelClient) QueryFollowsBatch(ids ...int64) (results interface{}, err error) {
	//TODO implement proper pagination
	//TODO implement the proper return type here
	id_strings := make([]string, len(ids))
	for i, id := range ids {
		id_strings[i] = fmt.Sprintf("%d", id) //do this more elegantly
	}
	err = c.execQueryThrottled("/follows/batch", GET, map[string]string{"ids": strings.Join(id_strings, ",")}, &results)
	return
}

//QueryUsersFollowers queries /users/:id/followers for a user's followers
func (c AngelClient) QueryUsersFollowers(user_id int64) (users []AngelUser, err error) {
	//TODO implement proper pagination
	err = c.execQueryThrottled(fmt.Sprintf("/startups/%d/followers", user_id), GET, map[string]string{}, users)
	return
}

//QueryUsersFollowersIds queries /users/:id/followers/ids for a user's followers
//Returns ids only
func (c AngelClient) QueryUsersFollowersIds(user_id int64) (ids []int64, err error) {
	//TODO implement proper pagination

	var tmp IdsResponse

	err = c.execQueryThrottled(fmt.Sprintf("/users/%d/followers/ids", user_id), GET, map[string]string{}, &tmp)
	ids = tmp.Ids
	return
}

//QueryUsersFollowingUsers queries /users/:id/following for a user's followers (return users only)
func (c AngelClient) QueryUsersFollowingUsers(user_id int64) (users []AngelUser, err error) {
	//TODO implement proper pagination

	var batch_response usersBatchResponse
	endpoint := fmt.Sprintf("/users/%d/following", user_id)
	err = c.execQueryThrottled(endpoint, GET, map[string]string{}, &batch_response)
	users = batch_response.Users
	return
}

//QueryUsersFollowingStartups queries  /users/:id/following for a user's followers (return startups only)
func (c AngelClient) QueryUsersFollowingStartups(user_id int64) (startups []Startup, err error) {
	//TODO implement proper pagination
	var batch_response startupsBatchResponse
	endpoint := fmt.Sprintf("/users/%d/following", user_id)
	err = c.execQueryThrottled(endpoint, GET, map[string]string{"type": "startup"}, &batch_response)
	startups = batch_response.Startups
	return
}

//QueryUsersFollowingUsersIds queries  /users/:id/following/ids for all users that the given user is following (return users only)
func (c AngelClient) QueryUsersFollowingUsersIds(user_id int64) (ids []int64, err error) {
	//TODO implement proper pagination
	var tmp IdsResponse
	endpoint := fmt.Sprintf("/users/%d/following/ids", user_id)
	err = c.execQueryThrottled(endpoint, GET, map[string]string{"type": "user"}, &tmp)
	ids = tmp.Ids
	return
}

//QueryUsersFollowingStartupsIds queries /users/:id/following/ids for all startups that the given user is following (return startups only)
func (c AngelClient) QueryUsersFollowingStartupsIds(user_id int64) (ids []int64, err error) {
	//TODO implement proper pagination
	var tmp IdsResponse
	endpoint := fmt.Sprintf("/users/%d/following/ids", user_id)
	err = c.execQueryThrottled(endpoint, GET, map[string]string{"type": "startup"}, &tmp)
	ids = tmp.Ids
	return
}

//QueryStartupsFollowers queries /startups/:id/followers for a startup's followers
func (c AngelClient) QueryStartupsFollowers(user_id int64) (users []AngelUser, err error) {
	//TODO implement proper pagination
	err = c.execQueryThrottled(fmt.Sprintf("/users/%d/followers", user_id), GET, map[string]string{}, users)
	return
}

//QueryStartupsFollowersIds queries /startups/:id/followers/ids for a startup's followers
func (c AngelClient) QueryStartupsFollowersIds(user_id int64) (ids []int64, err error) {
	//TODO implement proper pagination
	var tmp IdsResponse
	err = c.execQueryThrottled(fmt.Sprintf("/users/%d/followers/ids", user_id), GET, map[string]string{}, &tmp)
	ids = tmp.Ids
	return
}
