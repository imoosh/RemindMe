package wxmp

import "RemindMe/service"

type ApiGroup struct {
    UserApi
    CommonApi
    ActivityApi
}

var jwtService = service.ServiceGroupApp.WXMPServiceGroup.JwtService