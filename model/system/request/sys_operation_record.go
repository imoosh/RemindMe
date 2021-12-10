package request

import (
	"RemindMe/model/common/request"
	"RemindMe/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
