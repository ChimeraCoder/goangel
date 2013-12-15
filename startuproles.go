package angel

import (
	"fmt"
	"strconv"
)

const (
	DIRECTION_INCOMING = iota
	DIRECTION_OUTGOING = iota
)

//QueryStartupRoles queries the now-deprecated /startup_roles for all startup roles associated with the user with the given id
func (c AngelClient) QueryStartupRolesDeprecated(id int64, id_type int) (startuproles []StartupRole, err error) {
	var tmp struct {
		Startup_roles []StartupRole
	}
	switch id_type {
	case UserId:
		{
			err = c.execQueryThrottled("/startup_roles", GET, map[string]string{"user_id": strconv.FormatInt(id, 10)}, &tmp)
		}
	case StartupId:
		{
			err = c.execQueryThrottled("/startup_roles", GET, map[string]string{"startup_id": strconv.FormatInt(id, 10)}, &tmp)
		}
	default:
		return nil, fmt.Errorf("invalid id_type provided")
	}
	startuproles = tmp.Startup_roles
	return
}

//QueryStartupRoles queries the /startup_roles&v=1 endpoint for all startup roles associated with the user with the given id
func (c AngelClient) QueryStartupRoles(id int64, id_type int, v map[string]string) (startuproles []StartupRole, err error) {
	var tmp struct {
		Startup_roles []StartupRole
	}
	if v == nil {
		v = map[string]string{}
	}
	v["v"] = "1"
	switch id_type {
	case UserId:
		{
			v["user_id"] = strconv.FormatInt(id, 10)
			err = c.execQueryThrottled("/startup_roles", GET, v, &tmp)
		}
	case StartupId:
		{
			v["startup_id"] = strconv.FormatInt(id, 10)
			err = c.execQueryThrottled("/startup_roles", GET, v, &tmp)
		}
	default:
		return nil, fmt.Errorf("invalid id_type provided")
	}
	startuproles = tmp.Startup_roles
	return
}

// QueryStartupIdRoles queries  /startups/:id/roles
func (c AngelClient) QueryStartupIdRoles(id int64, direction int) (startuproles []StartupRole, err error) {
	var tmp struct {
		Startup_roles []StartupRole
	}

	v := map[string]string{}
	switch direction {
	case DIRECTION_INCOMING:
		{
			v["direction"] = "incoming"
		}
	case DIRECTION_OUTGOING:
		{
			v["direction"] = "outgoing"
		}
	default:
		return nil, fmt.Errorf("invalid id_type provided")
	}
	err = c.execQueryThrottled("/startups/"+strconv.FormatInt(id, 10)+"/roles", GET, v, &tmp)
	startuproles = tmp.Startup_roles
	return
}
