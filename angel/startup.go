package angel

type Startup struct {
	Id            int64          //Id refers to the AngelList ID
	Name          *string        ",omitempty"
	Company_url   *string        ",omitempty"
	Angellist_url *string        ",omitempty"
	StartupRoles  *[]StartupRole //People associated with the startup

	Product_desc      *string
	Community_profile bool
	Twitter_url       *string
	Follower_count    float64
	//Markets           []*Market
	//Locations         []Location
	//Screenshots       []*Screenshot
	Logo_url  *string
	Blog_url  *string
	Video_url *string
	Hidden    bool
	Status    struct {
		Created_at *string
		Message    *string
	}
}
