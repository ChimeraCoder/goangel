package angel

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

	Linkedin_url   *string
	Bio            *string
	Twitter_url    *string
	Follower_count float64
	Image          *string
	Facebook_url   *string
	Locations      []interface{}
	Blog_url       *string
	Online_bio_url *string
}
