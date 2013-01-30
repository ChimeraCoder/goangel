package angel

import (
    "fmt"
)


//Query /users/:id for a user's information given an id


//Query /users/:id/startups for a user's startup roles
func QueryUsersStartups(id int64) ([]StartupRole, error){
    var tmp struct{
        Startup_roles []StartupRole
    }
    endpoint := fmt.Sprintf("/users/%d/startups", id)
    if err := execQueryThrottled(endpoint, map[string]string{}, &tmp); err != nil{
        return nil, err
    }

	return tmp.Startup_roles, nil
}

//Query /users/batch for up to 50 users at a time, givne a comma-separated list of IDs



//Query /users/search for a user with the specified slug
//TODO fix this to search for emails as well
func QueryUsersSearch(slug string) (user AngelUser, err error) {
    err = execQueryThrottled("/users/search", map[string]string{"slug" : slug}, &user)
    return
}

