package router

import (
    "RemindMe/router/autocode"
    "RemindMe/router/example"
    "RemindMe/router/system"
    "RemindMe/router/wxmp"
)

type RouterGroup struct {
    System   system.RouterGroup
    Example  example.RouterGroup
    Autocode autocode.RouterGroup
    Wxmp     wxmp.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
