package request

import (
	"RemindMe/model/common/request"
	"RemindMe/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
