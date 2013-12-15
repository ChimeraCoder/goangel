package angel

import (
	"strconv"
)

type AngelUser struct {
	Angellist_url *string
	Id            int64
	Name          *string

	Roles []struct { //This corresponds to the "roles" field in AngelList, not the StartupRole
		Angellist_url *string
		Display_name  *string
		Id            int
		Name          *string
		Tag_type      *string
	}

	//StartupRoles []StartupRole

	Linkedin_url     *string
	Bio              *string
	Twitter_url      *string
	Follower_count   float64
	Image            *string
	Facebook_url     *string
	Locations        []Location
	Investor         bool
	Investor_details *struct {
		Startups_per_year string
		Average_amount    string
		Accreditation     string
		Markets           []Market
		Investments       []struct {
			Id      int64
			Name    string
			Quality int
		}
	}
}

type Location struct {
	Id            int64
	Tag_type      string
	Name          string
	Display_name  string
	Angellist_url string
}

type Market struct {
	Id            int64
	Tag_type      string
	Name          string
	Display_name  string
	Angellist_url string
}

// GetUsersId implements /users/:id
func (c AngelClient) GetUsersId(id int64, details string) (user AngelUser, err error) {
	v := map[string]string{}
	if details != "" {
		v["include_details"] = details
	}
	err = c.execQueryThrottled("/users/"+strconv.FormatInt(id, 10), GET, v, &user)
	return
}

func (u AngelUser) MarketNames() []string {
	names := []string{}
	for _, m := range u.Investor_details.Markets {
		names = append(names, m.Name)
	}
	return names
}
