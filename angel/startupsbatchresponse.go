package angel

type StartupsBatchResponse struct {
	Startups  []Startup
	Page      float64
	Per_page  float64
	Total     float64
	Last_page float64
}
