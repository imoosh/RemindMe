package source

import (
	"RemindMe/model"
	"time"

	"RemindMe/global"
	"RemindMe/model/system"
	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

var apis = []system.SysApi{
	{models.Model{ID: 1, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/base/login", "用户登录（必选）", "base", "POST"},
	{models.Model{ID: 2, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/user/register", "用户注册（必选）", "user", "POST"},
	{models.Model{ID: 3, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/api/createApi", "创建api", "api", "POST"},
	{models.Model{ID: 4, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/api/getApiList", "获取api列表", "api", "POST"},
	{models.Model{ID: 5, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/api/getApiById", "获取api详细信息", "api", "POST"},
	{models.Model{ID: 6, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/api/deleteApi", "删除Api", "api", "POST"},
	{models.Model{ID: 7, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/api/updateApi", "更新Api", "api", "POST"},
	{models.Model{ID: 8, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/api/getAllApis", "获取所有api", "api", "POST"},
	{models.Model{ID: 9, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/authority/createAuthority", "创建角色", "authority", "POST"},
	{models.Model{ID: 10, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/authority/deleteAuthority", "删除角色", "authority", "POST"},
	{models.Model{ID: 11, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/authority/getAuthorityList", "获取角色列表", "authority", "POST"},
	{models.Model{ID: 12, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/menu/getMenu", "获取菜单树（必选）", "menu", "POST"},
	{models.Model{ID: 13, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/menu/getMenuList", "分页获取基础menu列表", "menu", "POST"},
	{models.Model{ID: 14, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/menu/addBaseMenu", "新增菜单", "menu", "POST"},
	{models.Model{ID: 15, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/menu/getBaseMenuTree", "获取用户动态路由", "menu", "POST"},
	{models.Model{ID: 16, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/menu/addMenuAuthority", "增加menu和角色关联关系", "menu", "POST"},
	{models.Model{ID: 17, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/menu/getMenuAuthority", "获取指定角色menu", "menu", "POST"},
	{models.Model{ID: 18, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/menu/deleteBaseMenu", "删除菜单", "menu", "POST"},
	{models.Model{ID: 19, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/menu/updateBaseMenu", "更新菜单", "menu", "POST"},
	{models.Model{ID: 20, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/menu/getBaseMenuById", "根据id获取菜单", "menu", "POST"},
	{models.Model{ID: 21, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/user/changePassword", "修改密码（建议选择）", "user", "POST"},
	{models.Model{ID: 23, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/user/getUserList", "获取用户列表", "user", "POST"},
	{models.Model{ID: 24, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/user/setUserAuthority", "修改用户角色（必选）", "user", "POST"},
	{models.Model{ID: 25, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/fileUploadAndDownload/upload", "文件上传示例", "fileUploadAndDownload", "POST"},
	{models.Model{ID: 26, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/fileUploadAndDownload/getFileList", "获取上传文件列表", "fileUploadAndDownload", "POST"},
	{models.Model{ID: 27, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/casbin/updateCasbin", "更改角色api权限", "casbin", "POST"},
	{models.Model{ID: 28, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/casbin/getPolicyPathByAuthorityId", "获取权限列表", "casbin", "POST"},
	{models.Model{ID: 29, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/fileUploadAndDownload/deleteFile", "删除文件", "fileUploadAndDownload", "POST"},
	{models.Model{ID: 30, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/jwt/jsonInBlacklist", "jwt加入黑名单(退出，必选)", "jwt", "POST"},
	{models.Model{ID: 31, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/authority/setDataAuthority", "设置角色资源权限", "authority", "POST"},
	{models.Model{ID: 32, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/system/getSystemConfig", "获取配置文件内容", "system", "POST"},
	{models.Model{ID: 33, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/system/setSystemConfig", "设置配置文件内容", "system", "POST"},
	{models.Model{ID: 34, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/customer/customer", "创建客户", "customer", "POST"},
	{models.Model{ID: 35, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/customer/customer", "更新客户", "customer", "PUT"},
	{models.Model{ID: 36, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/customer/customer", "删除客户", "customer", "DELETE"},
	{models.Model{ID: 37, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/customer/customer", "获取单一客户", "customer", "GET"},
	{models.Model{ID: 38, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/customer/customerList", "获取客户列表", "customer", "GET"},
	{models.Model{ID: 39, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/casbin/casbinTest/:pathParam", "RESTFUL模式测试", "casbin", "GET"},
	{models.Model{ID: 40, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/autoCode/createTemp", "自动化代码", "autoCode", "POST"},
	{models.Model{ID: 41, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/authority/updateAuthority", "更新角色信息", "authority", "PUT"},
	{models.Model{ID: 42, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/authority/copyAuthority", "拷贝角色", "authority", "POST"},
	{models.Model{ID: 43, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/user/deleteUser", "删除用户", "user", "DELETE"},
	{models.Model{ID: 44, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysDictionaryDetail/createSysDictionaryDetail", "新增字典内容", "sysDictionaryDetail", "POST"},
	{models.Model{ID: 45, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysDictionaryDetail/deleteSysDictionaryDetail", "删除字典内容", "sysDictionaryDetail", "DELETE"},
	{models.Model{ID: 46, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysDictionaryDetail/updateSysDictionaryDetail", "更新字典内容", "sysDictionaryDetail", "PUT"},
	{models.Model{ID: 47, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysDictionaryDetail/findSysDictionaryDetail", "根据ID获取字典内容", "sysDictionaryDetail", "GET"},
	{models.Model{ID: 48, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysDictionaryDetail/getSysDictionaryDetailList", "获取字典内容列表", "sysDictionaryDetail", "GET"},
	{models.Model{ID: 49, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysDictionary/createSysDictionary", "新增字典", "sysDictionary", "POST"},
	{models.Model{ID: 50, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysDictionary/deleteSysDictionary", "删除字典", "sysDictionary", "DELETE"},
	{models.Model{ID: 51, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysDictionary/updateSysDictionary", "更新字典", "sysDictionary", "PUT"},
	{models.Model{ID: 52, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysDictionary/findSysDictionary", "根据ID获取字典", "sysDictionary", "GET"},
	{models.Model{ID: 53, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysDictionary/getSysDictionaryList", "获取字典列表", "sysDictionary", "GET"},
	{models.Model{ID: 54, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysOperationRecord/createSysOperationRecord", "新增操作记录", "sysOperationRecord", "POST"},
	{models.Model{ID: 55, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysOperationRecord/deleteSysOperationRecord", "删除操作记录", "sysOperationRecord", "DELETE"},
	{models.Model{ID: 56, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysOperationRecord/findSysOperationRecord", "根据ID获取操作记录", "sysOperationRecord", "GET"},
	{models.Model{ID: 57, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysOperationRecord/getSysOperationRecordList", "获取操作记录列表", "sysOperationRecord", "GET"},
	{models.Model{ID: 58, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/autoCode/getTables", "获取数据库表", "autoCode", "GET"},
	{models.Model{ID: 59, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/autoCode/getDB", "获取所有数据库", "autoCode", "GET"},
	{models.Model{ID: 60, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/autoCode/getColumn", "获取所选table的所有字段", "autoCode", "GET"},
	{models.Model{ID: 61, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/sysOperationRecord/deleteSysOperationRecordByIds", "批量删除操作历史", "sysOperationRecord", "DELETE"},
	{models.Model{ID: 62, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/simpleUploader/upload", "插件版分片上传", "simpleUploader", "POST"},
	{models.Model{ID: 63, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/simpleUploader/checkFileMd5", "文件完整度验证", "simpleUploader", "GET"},
	{models.Model{ID: 64, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/simpleUploader/mergeFileMd5", "上传完成合并文件", "simpleUploader", "GET"},
	{models.Model{ID: 65, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/user/setUserInfo", "设置用户信息（必选）", "user", "PUT"},
	{models.Model{ID: 66, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/system/getServerInfo", "获取服务器信息", "system", "POST"},
	{models.Model{ID: 67, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/email/emailTest", "发送测试邮件", "email", "POST"},
	{models.Model{ID: 80, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/autoCode/preview", "预览自动化代码", "autoCode", "POST"},
	{models.Model{ID: 81, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/excel/importExcel", "导入excel", "excel", "POST"},
	{models.Model{ID: 82, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/excel/loadExcel", "下载excel", "excel", "GET"},
	{models.Model{ID: 83, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/excel/exportExcel", "导出excel", "excel", "POST"},
	{models.Model{ID: 84, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/excel/downloadTemplate", "下载excel模板", "excel", "GET"},
	{models.Model{ID: 85, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/api/deleteApisByIds", "批量删除api", "api", "DELETE"},
	{models.Model{ID: 86, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/autoCode/getSysHistory", "查询回滚记录", "autoCode", "POST"},
	{models.Model{ID: 87, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/autoCode/rollback", "回滚自动生成代码", "autoCode", "POST"},
	{models.Model{ID: 88, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/autoCode/getMeta", "获取meta信息", "autoCode", "POST"},
	{models.Model{ID: 89, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/autoCode/delSysHistory", "删除回滚记录", "autoCode", "POST"},
	{models.Model{ID: 90, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/user/setUserAuthorities", "设置权限组", "user", "POST"},
	{models.Model{ID: 91, CreatedAt: models.LocalTime{time.Now()}, UpdatedAt: models.LocalTime{time.Now()}}, "/user/getUserInfo", "获取自身信息（必选）", "user", "GET"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_apis 表数据初始化
func (a *api) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 67}).Find(&[]system.SysApi{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_apis 表初始数据成功!")
		return nil
	})
}
