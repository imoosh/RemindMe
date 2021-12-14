package initialize

import (
	"fmt"

	"RemindMe/config"
	"RemindMe/global"
	"RemindMe/utils"
)

func Timer() {
	if global.Config.Timer.Start {
		for i := range global.Config.Timer.Detail {
			go func(detail config.Detail) {
				global.Timer.AddTaskByFunc("ClearDB", global.Config.Timer.Spec, func() {
					err := utils.ClearTable(global.DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.Config.Timer.Detail[i])
		}
	}
}
