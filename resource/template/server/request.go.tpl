package request

import (
	"RemindMe/model/autocode"
	"RemindMe/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}
