package v1

import (
    "RemindMe/api/v1/autocode"
    "RemindMe/api/v1/example"
    "RemindMe/api/v1/system"
    "RemindMe/api/v1/wxmp"
)

type ApiGroup struct {
    ExampleApiGroup  example.ApiGroup
    SystemApiGroup   system.ApiGroup
    AutoCodeApiGroup autocode.ApiGroup
    WxmpApiGroup     wxmp.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
