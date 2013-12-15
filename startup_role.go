package angel

type StartupRole struct {
	Id        int64
	Confirmed bool
	User      *struct { //This isn't a full AngelUser - just a snippet returned in the query response
		Follower_count int64
		Image          string
		Name           string
		Id             int64
		Angellist_url  string
		Bio            string
	}
	Role       string
	Created_at string
	Startup    *struct {
		Id   int64
		Name string
	}

	Tagged struct {
		Type           string
		Name           string
		Id             int64
		Bio            string
		Follower_count int64
		Angellist_url  string
		Image          string
	}
}
