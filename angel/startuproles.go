package angel

import (
    "fmt"
    "strconv"
    "log"
    "encoding/json"
)

//Query /startup_roles for all startup roles associated with the user with the given id
func QueryStartupRoles(id int64, id_type int) ([]StartupRole, error) {
	resp_ch := make(chan QueryResponse)

	switch id_type {
	case UserId:
		{
			queryQueue <- QueryChan{"/startup_roles", map[string]string{"user_id": strconv.FormatInt(id, 10)}, resp_ch}
		}
	case StartupId:
		{
			queryQueue <- QueryChan{"/startup_roles", map[string]string{"startup_id": strconv.FormatInt(id, 10)}, resp_ch}
		}
	default:
		return nil, fmt.Errorf("invalid id_type provided")
	}
	r := <-resp_ch
	res := r.result
	if err := r.err; err != nil {
		return nil, err
	}

	roles_array := res["startup_roles"].([]interface{})

	var roles []StartupRole
	roles_bts, err := json.Marshal(roles_array)
	if err != nil {
		log.Print("Woah, error occured while marshalling")
		panic(err)
	}
	if err := json.Unmarshal(roles_bts, &roles); err != nil {
		return nil, err
	}
	return roles, err
}

