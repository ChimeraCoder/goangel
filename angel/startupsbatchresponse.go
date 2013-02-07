package angel

type startupsBatchResponse struct {
	Startups  []Startup
	Page      float64
	Per_page  float64
	Total     float64
	Last_page float64
}
