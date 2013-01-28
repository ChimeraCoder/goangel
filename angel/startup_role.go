package angel

type StartupRole struct {
	Id        float64
	Confirmed bool
	User      struct { //This isn't a full AngelUser - just a snippet returned in the query response
		Follower_count float64
		Image          string
		Name           string
		Id             float64
		Angellist_url  string
		Bio            string
	}
	Role       string
	Created_at string
}
