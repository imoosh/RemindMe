package source

import (
	"RemindMe/model"
	"time"

	"RemindMe/global"
	"RemindMe/model/system"
	"github.com/gookit/color"

	"gorm.io/gorm"
)

var DictionaryDetail = new(dictionaryDetail)

type dictionaryDetail struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: dictionary_details 表数据初始化
func (d *dictionaryDetail) Init() error {
	var details = []system.SysDictionaryDetail{
		{models.Model{ID: 1, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "smallint", 1, status, 1, 2},
		{models.Model{ID: 2, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "mediumint", 2, status, 2, 2},
		{models.Model{ID: 3, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "int", 3, status, 3, 2},
		{models.Model{ID: 4, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "bigint", 4, status, 4, 2},
		{models.Model{ID: 5, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "date", 0, status, 0, 3},
		{models.Model{ID: 6, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "time", 1, status, 1, 3},
		{models.Model{ID: 7, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "year", 2, status, 2, 3},
		{models.Model{ID: 8, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "datetime", 3, status, 3, 3},
		{models.Model{ID: 9, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "timestamp", 5, status, 5, 3},
		{models.Model{ID: 10, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "float", 0, status, 0, 4},
		{models.Model{ID: 11, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "double", 1, status, 1, 4},
		{models.Model{ID: 12, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "decimal", 2, status, 2, 4},
		{models.Model{ID: 13, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "char", 0, status, 0, 5},
		{models.Model{ID: 14, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "varchar", 1, status, 1, 5},
		{models.Model{ID: 15, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "tinyblob", 2, status, 2, 5},
		{models.Model{ID: 16, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "tinytext", 3, status, 3, 5},
		{models.Model{ID: 17, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "text", 4, status, 4, 5},
		{models.Model{ID: 18, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "blob", 5, status, 5, 5},
		{models.Model{ID: 19, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "mediumblob", 6, status, 6, 5},
		{models.Model{ID: 20, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "mediumtext", 7, status, 7, 5},
		{models.Model{ID: 21, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "longblob", 8, status, 8, 5},
		{models.Model{ID: 22, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "longtext", 9, status, 9, 5},
		{models.Model{ID: 23, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "tinyint", 0, status, 0, 6},
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 23}).Find(&[]system.SysDictionaryDetail{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_dictionary_details 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&details).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_dictionary_details 表初始数据成功!")
		return nil
	})
}
