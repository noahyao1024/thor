package container

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	localDBClient *gorm.DB
	sqlLitePath   string
)

// InitDB ...
func InitDB() {
	sqlLitePath, _ = os.LookupEnv("SQLITE_PATH")
	if len(sqlLitePath) > 0 {
		var err error
		localDBClient, err = gorm.Open(sqlite.Open(strings.TrimRight(sqlLitePath, "/")+"/crucial.db"), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("%s - %s", strings.TrimRight(sqlLitePath, "/")+"/crucial.db", err.Error()))
		}
	}
}

// GetLocalDB ...
func GetLocalDB() *gorm.DB {
	return localDBClient
}

// GetSQLLitePath ...
func GetSQLLitePath() string {
	return sqlLitePath
}
