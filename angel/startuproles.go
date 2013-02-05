package angel

import (
	"fmt"
	"strconv"
)

//Query /startup_roles for all startup roles associated with the user with the given id
func QueryStartupRoles(id int64, id_type int) (startuproles []StartupRole, err error) {
	var tmp struct {
		Startup_roles []StartupRole
	}
	switch id_type {
	case UserId:
		{
			err = execQueryThrottled("/startup_roles", GET, map[string]string{"user_id": strconv.FormatInt(id, 10)}, &tmp)
		}
	case StartupId:
		{
			err = execQueryThrottled("/startup_roles", GET, map[string]string{"startup_id": strconv.FormatInt(id, 10)}, &tmp)
		}
	default:
		return nil, fmt.Errorf("invalid id_type provided")
	}
	startuproles = tmp.Startup_roles
	return
}
