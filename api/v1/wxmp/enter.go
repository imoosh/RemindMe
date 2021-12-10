package wxmp

import "RemindMe/service"

type ApiGroup struct {
    UserApi
    CommonApi
}

var jwtService = service.ServiceGroupApp.WXMPServiceGroup.JwtService