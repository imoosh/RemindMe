package wxmp

import "RemindMe/service"

type ApiGroup struct {
    UserApi
    BaseApi
    ActivityApi
}

var jwtService = service.ServiceGroupApp.WXMPServiceGroup.JwtService
var userService = service.ServiceGroupApp.WXMPServiceGroup.UserService
var activityService = service.ServiceGroupApp.WXMPServiceGroup.ActivityService
