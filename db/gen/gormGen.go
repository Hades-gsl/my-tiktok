package main

import (
	"tiktok/db/model"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE token = @token{{end}}
	FilterWithToken(token string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "db",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API
	g.ApplyBasic(model.User{}, model.Video{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface
	g.ApplyInterface(func(Querier) {}, model.UserToken{})

	// Generate the code
	g.Execute()
}
