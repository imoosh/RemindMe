package wxmp

import "RemindMe/service"

type ApiGroup struct {
    UserApi
    BaseApi
    PlatformApi
    ActivityApi
}

var jwtService = service.ServiceGroupApp.WXMPServiceGroup.JwtService
var userService = service.ServiceGroupApp.WXMPServiceGroup.UserService
var activityService = service.ServiceGroupApp.WXMPServiceGroup.ActivityService
var platformService = service.ServiceGroupApp.WXMPServiceGroup.PlatformService
