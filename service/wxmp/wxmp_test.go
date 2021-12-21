package wxmp

import (
    "RemindMe/global"
    "RemindMe/model/system"
    "RemindMe/model/wxmp"
    wxmpReq "RemindMe/model/wxmp/request"
    "RemindMe/utils"
    "flag"
    "fmt"
    "github.com/fsnotify/fsnotify"
    "github.com/songzhibin97/gkit/cache/local_cache"
    "github.com/spf13/viper"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
    "path/filepath"
    "testing"
    "time"
)

func Gorm() *gorm.DB {
    switch global.Config.System.DbType {
    case "mysql":
        return GormMysql()
    default:
        return GormMysql()
    }
}

// MysqlTables
//@author: SliverHorn
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func MysqlTables(db *gorm.DB) {
    err := db.AutoMigrate(
        system.SysUser{},
        system.SysAuthority{},
        system.SysApi{},
        system.SysBaseMenu{},
        system.SysBaseMenuParameter{},
        system.JwtBlacklist{},
        system.SysDictionary{},
        system.SysDictionaryDetail{},
        system.SysOperationRecord{},
        system.SysAutoCodeHistory{},

        wxmp.User{},
        wxmp.Dict{},
        wxmp.Activity{},
        wxmp.ActivitySubscriber{},
    )
    if err != nil {
        os.Exit(0)
    }
}

//@author: SliverHorn
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
    m := global.Config.Mysql
    if m.Dbname == "" {
        return nil
    }
    dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
    mysqlConfig := mysql.Config{
        DSN:                       dsn,   // DSN data source name
        DefaultStringSize:         191,   // string 类型字段的默认长度
        DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
        DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
        DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
        SkipInitializeWithVersion: false, // 根据版本自动配置
    }
    if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
        //global.Log.Error("MySQL启动异常", zap.Any("err", err))
        //os.Exit(0)
        //return nil
        return nil
    } else {
        sqlDB, _ := db.DB()
        sqlDB.SetMaxIdleConns(m.MaxIdleConns)
        sqlDB.SetMaxOpenConns(m.MaxOpenConns)
        return db
    }
}
func gormConfig() *gorm.Config {
    config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
    return config
}

var level zapcore.Level

func Zap() (logger *zap.Logger) {
    if ok, _ := utils.PathExists(global.Config.Zap.Director); !ok { // 判断是否有Director文件夹
        fmt.Printf("create %v directory\n", global.Config.Zap.Director)
        _ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
    }

    switch global.Config.Zap.Level { // 初始化配置文件的Level
    case "debug":
        level = zap.DebugLevel
    case "info":
        level = zap.InfoLevel
    case "warn":
        level = zap.WarnLevel
    case "error":
        level = zap.ErrorLevel
    case "dpanic":
        level = zap.DPanicLevel
    case "panic":
        level = zap.PanicLevel
    case "fatal":
        level = zap.FatalLevel
    default:
        level = zap.InfoLevel
    }

    if level == zap.DebugLevel || level == zap.ErrorLevel {
        logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
    } else {
        logger = zap.New(getEncoderCore())
    }
    if global.Config.Zap.ShowLine {
        logger = logger.WithOptions(zap.AddCaller())
    }
    return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
    config = zapcore.EncoderConfig{
        MessageKey:     "message",
        LevelKey:       "level",
        TimeKey:        "time",
        NameKey:        "logger",
        CallerKey:      "caller",
        StacktraceKey:  global.Config.Zap.StacktraceKey,
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     CustomTimeEncoder,
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.FullCallerEncoder,
    }
    switch {
    case global.Config.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
        config.EncodeLevel = zapcore.LowercaseLevelEncoder
    case global.Config.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
        config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
    case global.Config.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
        config.EncodeLevel = zapcore.CapitalLevelEncoder
    case global.Config.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
        config.EncodeLevel = zapcore.CapitalColorLevelEncoder
    default:
        config.EncodeLevel = zapcore.LowercaseLevelEncoder
    }
    return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
    if global.Config.Zap.Format == "json" {
        return zapcore.NewJSONEncoder(getEncoderConfig())
    }
    return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
    writer, err := utils.GetWriteSyncer() // 使用file-rotatelogs进行日志分割
    if err != nil {
        fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
        return
    }
    return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format(global.Config.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}

func Viper(path ...string) *viper.Viper {
    var config string
    if len(path) == 0 {
        flag.StringVar(&config, "c", "", "choose config file.")
        flag.Parse()
        if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
            if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
                config = utils.ConfigFile
                fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
            } else {
                config = configEnv
                fmt.Printf("您正在使用Config环境变量,config的路径为%v\n", config)
            }
        } else {
            fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
        }
    } else {
        config = path[0]
        fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
    }

    v := viper.New()
    v.SetConfigFile(config)
    v.SetConfigType("yaml")
    err := v.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
    v.WatchConfig()

    v.OnConfigChange(func(e fsnotify.Event) {
        fmt.Println("config file changed:", e.Name)
        if err := v.Unmarshal(&global.Config); err != nil {
            fmt.Println(err)
        }
    })
    if err := v.Unmarshal(&global.Config); err != nil {
        fmt.Println(err)
    }
    // root 适配性
    // 根据root位置去找到对应迁移位置,保证root路径有效
    global.Config.AutoCode.Root, _ = filepath.Abs("..")
    global.BlackCache = local_cache.NewCache(
        local_cache.SetDefaultExpire(time.Second * time.Duration(global.Config.JWT.ExpiresTime)),
    )
    return v
}

func Init() {
    global.VP = Viper("/Users/wayne/go/src/RemindMe/config.yaml")
    global.Log = Zap()
    global.DB = Gorm() // gorm连接数据库
    if global.DB == nil {
        fmt.Println("gorm failed")
        os.Exit(-1)
    }

    {
        MysqlTables(global.DB) // 初始化表
        // 程序结束前关闭数据库链接
        //db, _ := global.DB.DB()
        //defer db.Close()
    }
}

var userService = &UserService{}
var activityService = &ActivityService{}

// 创建活动
func TestActivityService_CreateActivity(t *testing.T) {
    Init()

    //u := &wxmp.User{
    //    Nickname: "布谷鸟",
    //    OpenID:  uuid.NewV4().String(),
    //}
    //userService.CreateUser(u)

    a := &wxmp.Activity{
        Title: "吃饭啊",
    }
    activityService.CreateActivity(2, a)
}

// 查询活动列表
func TestActivityService_QueryActivities(t *testing.T) {
    Init()
    list, _ := activityService.QueryActivities(1)
    t.Log(list)
}

// 查询活动详情
func TestActivityService_QueryActivityDetail(t *testing.T) {
    Init()
    a, _ := activityService.ActivityDetail(2)
    t.Log(a)
}

func TestActivityService_SubscribeActivity(t *testing.T) {
    Init()
    activityService.SubscribeActivity(2, &wxmpReq.ActivitySubscribeRequest{Id: 1})
    activityService.SubscribeActivity(2, &wxmpReq.ActivitySubscribeRequest{Id: 2})
}

func TestActivityService_UnsubscribeActivity(t *testing.T) {
    Init()
    activityService.UnsubscribeActivity(1, 1)
    activityService.UnsubscribeActivity(1, 2)
}

func TestActivityService_ActivitySubscribers(t *testing.T) {
    Init()
    ac, _ := activityService.ActivitySubscribers(1)
    t.Log(ac)
    t.Log(ac.Subscriptions)
}
