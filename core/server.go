package core

import (
    "RemindMe/global"
    "RemindMe/initialize"
    "RemindMe/service/system"
    "fmt"
)

type server interface {
    ListenAndServe() error
}

func RunWindowsServer() {
    if global.Config.System.UseMultipoint {
        // 初始化redis服务
        initialize.Redis()
    }

    // 从db加载jwt数据
    if global.DB != nil {
        system.LoadAll()
    }

    //pcRouter := initialize.Routers()
    wxmpRouter := initialize.WxmpRouters()

    //pcRouter.Static("/form-generator", "./resource/page")

    //pcAddr := fmt.Sprintf(":%d", global.Config.System.Addr)
    //pcServe := initServer(pcAddr, pcRouter)

    wxmpAddr := fmt.Sprintf(":80")
    wxmpSeve := initServer(wxmpAddr, wxmpRouter)

    //go global.Log.Error(pcServe.ListenAndServe().Error())
    global.Log.Error(wxmpSeve.ListenAndServe().Error())
}
