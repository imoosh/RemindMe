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
    if global.GVA_CONFIG.System.UseMultipoint {
        // 初始化redis服务
        initialize.Redis()
    }

    // 从db加载jwt数据
    if global.GVA_DB != nil {
        system.LoadAll()
    }

    //pcRouter := initialize.Routers()
    wxmpRouter := initialize.WxmpRouters()

    //pcRouter.Static("/form-generator", "./resource/page")

    //pcAddr := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
    //pcServe := initServer(pcAddr, pcRouter)

    wxmpAddr := fmt.Sprintf(":8000")
    wxmpSeve := initServer(wxmpAddr, wxmpRouter)

    //go global.GVA_LOG.Error(pcServe.ListenAndServe().Error())
    global.GVA_LOG.Error(wxmpSeve.ListenAndServe().Error())
}
