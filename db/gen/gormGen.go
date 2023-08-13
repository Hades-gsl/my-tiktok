package main

import (
	"tiktok/db/model"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
// type Querier interface {
// 	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
// 	FilterWithNameAndRole(name, role string) ([]gen.T, error)
// }

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "tiktok/db",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open("root:gsl.2326@(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API
	g.ApplyBasic(model.User{}, model.Video{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Video{})

	// Generate the code
	g.Execute()
}
