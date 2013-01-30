package angel

import (
    "fmt"
    "strconv"
    "encoding/json"
)

//Query /startup_roles for all startup roles associated with the user with the given id
func QueryStartupRoles(id int64, id_type int) ([]StartupRole, error) {
    var err error
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

    var tmp struct{
        Startup_roles []StartupRole
    }

	if err := json.Unmarshal(res, &tmp); err != nil {
		return nil, err
	}
	return tmp.Startup_roles, err
}

