package angel

import (
	"fmt"
)

//QueryTags queries /tags/:id for information on a given tag
//This endpoint is currently giving trouble on AngelList
func QueryTags(id int64) (tag Tag, err error) {
	err = execQueryThrottled(fmt.Sprintf("/tags/%d/", id), GET, map[string]string{}, &tag)
	return
}
