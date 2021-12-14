// 自动生成模板SysDictionaryDetail
package autocode

import (
	"RemindMe/model"
)

// 如果含有time.Time 请自行import time包
type AutoCodeExample struct {
	models.Model
	AutoCodeExampleField string `json:"autoCodeExampleField" form:"autoCodeExampleField" gorm:"column:auto_code_example_field;comment:仅作示例条目无实际作用"` // 展示值
}
