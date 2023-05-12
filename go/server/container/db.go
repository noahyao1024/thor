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
	if len(sqlLitePath) == 0 {
		// FIXME
		sqlLitePath = "/sdcard/Android/data/com.noahyao.thor"
		sqlLitePath = "./data/data/com.noahyao.thor"
	}

	if len(sqlLitePath) > 0 {
		var err error
		localDBClient, err = gorm.Open(sqlite.Open(strings.TrimRight(sqlLitePath, "/")+"/crucial.db"), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("%s - %s", strings.TrimRight(sqlLitePath, "/")+"/crucial.db", err.Error()))
		}

		localDBClient.Exec(`
		CREATE TABLE IF NOT EXISTS reports (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL,
				data TEXT NULL,
				status TEXT NULL,
				create_time DATETIME DEFAULT (datetime('now')),
				modify_time DATETIME DEFAULT (datetime('now', 'localtime'))
			  );
			  CREATE TABLE IF NOT EXISTS cases (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				template_id INTEGER,
				template_name TEXT NULL,
				report_id INTEGER NULL,
				status TEXT NULL,
				data TEXT NULL,
				apis TEXT NULL,
				create_time DATETIME DEFAULT (datetime('now')),
				modify_time DATETIME DEFAULT (datetime('now', 'localtime'))
			  );

			  `)
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
