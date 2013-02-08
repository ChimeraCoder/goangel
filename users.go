package angel

import (
	"fmt"
	"strings"
)

//QueryUsers queries /users/:id for a user's information given an id
func QueryUsers(id int64) (user AngelUser, err error) {
	err = execQueryThrottled(fmt.Sprintf("/users/%d", id), GET, map[string]string{}, &user)
	return
}

//QueryUsersStartups queries /users/:id/startups for a user's startup roles
func QueryUsersStartups(id int64) ([]StartupRole, error) {
	var tmp struct {
		Startup_roles []StartupRole
	}
	endpoint := fmt.Sprintf("/users/%d/startups", id)
	if err := execQueryThrottled(endpoint, GET, map[string]string{}, &tmp); err != nil {
		return nil, err
	}

	return tmp.Startup_roles, nil
}

//QueryUsersBatch queries /users/batch for up to 50 users at a time, given a comma-separated list of IDs
func QueryUsersBatch(ids ...int64) (results interface{}, err error) {
	//TODO implement proper pagination
	//TODO implement the proper return type here
	id_strings := make([]string, len(ids))
	for i, id := range ids {
		id_strings[i] = fmt.Sprintf("%d", id) //do this more elegantly
	}
	err = execQueryThrottled("/users/batch", GET, map[string]string{"ids": strings.Join(id_strings, ",")}, &results)
	return
}

//QueryUsersSearch queries /users/search for a user with the specified slug
func QueryUsersSearch(slug string) (user AngelUser, err error) {
	//TODO fix this to search for emails as well
	err = execQueryThrottled("/users/search", GET, map[string]string{"slug": slug}, &user)
	return
}
