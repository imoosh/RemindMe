package request

import (
	"RemindMe/model/common/request"
	"RemindMe/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
