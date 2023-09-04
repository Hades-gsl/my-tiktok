package main

import (
	"tiktok/config"
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

	gormdb, _ := gorm.Open(mysql.Open(config.DSN))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API
	g.ApplyBasic(model.User{}, model.Video{}, model.Favorite{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface
	// g.ApplyInterface(func(Querier) {}, model.UserToken{})

	// Generate the code
	g.Execute()
}
