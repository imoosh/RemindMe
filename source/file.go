package source

import (
	"RemindMe/model"
	"time"

	"RemindMe/global"
	"RemindMe/model/example"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var File = new(file)

type file struct{}

var files = []example.ExaFileUploadAndDownload{
	{models.Model{ID: 1, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "10.png", "https://qmplusimg.henrongyi.top/gvalogo.png", "png", "158787308910.png"},
	{models.Model{ID: 2, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "logo.png", "https://qmplusimg.henrongyi.top/1576554439myAvatar.png", "png", "1587973709logo.png"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: exa_file_upload_and_downloads 表初始化数据
func (f *file) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]example.ExaFileUploadAndDownload{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> exa_file_upload_and_downloads 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&files).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> exa_file_upload_and_downloads 表初始数据成功!")
		return nil
	})
}
