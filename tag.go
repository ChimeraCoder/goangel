package angel

type Tag struct {
	Angellist_url *string
	Statistics    *struct {
		Direct *struct {
			Investor_followers float64
			Users              float64
			Followers          float64
			Startups           float64
		}
		All *struct {
			Followers          float64
			Startups           float64
			Investor_followers float64
			Users              float64
		}
	}
	Name         *string
	Display_name *string
	Id           float64
	Tag_type     *string
}
