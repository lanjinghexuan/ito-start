
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type UsersSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Username  *string `json:"username" form:"username"` 
      Mobile  *string `json:"mobile" form:"mobile"` 
      Image  *string `json:"image" form:"image"` 
      Password  *string `json:"password" form:"password"` 
    request.PageInfo
}
