package angel

type Comment struct {
	Created_at *string
	Id         float64
	User       *struct {
		Angellist_url  *string
		Bio            *string
		Follower_count *float64
		Image          *string
		Name           *string
		Id             *float64
	}
	Comment *string
}
