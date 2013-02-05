package angel

import (
	"fmt"
	"strings"
)

const (
	DESC       = iota
	ASC        = iota
	POPULARITY = iota
)

//Query /startups/:id for a startup, given its id
func QueryStartupsId(startup_id int64) (startup Startup, err error) {
	err = execQueryThrottled(fmt.Sprintf("/startups/%d/", startup_id), GET, map[string]string{}, &startup)
	return
}

//Query /startups/:id/comments for a startup's comments, given its id
func QueryStartupsIdComments(startup_id int64) (comments []Comment, err error) {
	err = execQueryThrottled(fmt.Sprintf("/startups/%d/comments", startup_id), GET, map[string]string{}, &comments)
	return
}

//Query /startups/:id/users for a startup's tagged users, given its id
func QueryStartupsIdUsers(startup_id int64) (users []AngelUser, err error) {
	var result struct {
		Startup_roles []AngelUser
	}
	err = execQueryThrottled(fmt.Sprintf("/startups/%d/users", startup_id), GET, map[string]string{}, &result)
	users = result.Startup_roles
	return
}

//Query /startups/batch for the followers of several users
func QueryStartupsBatch(ids ...int64) (startups []Startup, err error) {
	id_strings := make([]string, len(ids))
	for i, id := range ids {
		id_strings[i] = fmt.Sprintf("%d", id) //do this more elegantly
	}
	err = execQueryThrottled("/startups/batch", GET, map[string]string{"ids": strings.Join(id_strings, ",")}, &startups)
	return
}

//Query /startups/search for the followers of several users
//TODO implement domain
func QueryStartupsSearch(slug string) (startup Startup, err error) {
	err = execQueryThrottled("/startups/search", GET, map[string]string{"slug": slug}, &startup)
	return
}

//Query /tags/:id/startups for startups tagged with the given tag or a child of the given tag
//TODO implement pagination
func QueryTagsStartups(startup_id int64, order int) (startups []Startup, err error) {
	var order_s string
	switch order {
	case DESC:
		order_s = "desc"
	case POPULARITY:
		order_s = "popularity"
	case ASC:
		order_s = "asc"
	default:
		order_s = "desc"
	}
	var result struct {
		Startups []Startup
	}
	err = execQueryThrottled(fmt.Sprintf("/tags/%d/startups", startup_id), GET, map[string]string{"order": order_s}, &result)
	startups = result.Startups
	return
}
