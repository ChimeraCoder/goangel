package angel

type UsersBatchResponse struct {
	Page      int64
	Users     []AngelUser
	Per_page  int64
	Total     int64
	Last_page int64
}
