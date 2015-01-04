package db

import (
    "github.com/revel/revel"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "strings"
    "fmt"
    // "time"
)

var Gdb gorm.DB
// var GCacheTime time.Duration
// var TimeFormat string

func getConnectionString() string {
    host := getParamString("db.host", "")
    port := getParamString("db.port", "3306")
    user := getParamString("db.user", "")
    pass := getParamString("db.password", "")
    dbname := getParamString("db.name", "")
    protocol := getParamString("db.protocol", "tcp")
    dbargs := getParamString("db.args", " ")

    if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }

    conStr := fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", user, pass, protocol, host, port, dbname, dbargs)
    // revel.INFO.Println(conStr)
    return conStr
}

func InitDB() {
    var err error

    connectionString := getConnectionString()
    Gdb, err = gorm.Open("mysql", connectionString)
    if err != nil {
        revel.ERROR.Println("FATAL", err)
        panic( err )
    }
    Gdb.LogMode(getParamBool("db.debug"))
    // GCacheTime = getCacheTime()
    // TimeFormat = "2006-01-02 15:04:06"
}

func getParamString(param string, defaultValue string) string {
    p, found := revel.Config.String(param)
    if !found {
        if defaultValue == "" {
            revel.ERROR.Fatal("Cound not find parameter: " + param)
        } else {
            return defaultValue
        }
    }
    return p
}

func getParamInt(param string) int {
    p, found := revel.Config.Int(param)
    if !found {
        revel.ERROR.Fatal("Cound not find parameter: " + param)
    }
    return p
}

func getParamBool(param string) bool {
    p, found := revel.Config.Bool(param)
    if !found {
        revel.ERROR.Fatal("Cound not find parameter: " + param)
    }
    return p
}

// func getCacheTime() time.Duration {
//     cacheTime := time.Duration(getParamInt("cache.time")) * time.Minute
//     if cacheTime < 0 {
//         cacheTime = 1 * time.Nanosecond
//     }
//     return cacheTime
// }
