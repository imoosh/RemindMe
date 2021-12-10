package service

import (
	"RemindMe/service/autocode"
	"RemindMe/service/example"
	"RemindMe/service/system"
	"RemindMe/service/wxmp"
)

type ServiceGroup struct {
	ExampleServiceGroup  example.ServiceGroup
	SystemServiceGroup   system.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
	WXMPServiceGroup wxmp.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
