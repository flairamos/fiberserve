package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "D:\\Programs\\cwm-gov-platform-latest\\cwm-gov-platform\\internal\\autoMigrate",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(postgres.Open("postgres://postgres:langong@2023@113.59.48.145:6024/hainan"), &gorm.Config{})
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	g.ApplyBasic(
		// Generate struct `TIotVideoDevice` based on table `t_iot_video_device`
		g.GenerateModel("transport_direction"),
		g.GenerateModel("transport_order"),
		g.GenerateModel("transport_order_trash"),
		g.GenerateModel("transport_order_vehicle"),
		g.GenerateModel("transport_task_record"),

		// Generate struct `Employee` based on table `users`
		// g.GenerateModelAs("users", "Employee"),

		// Generate struct `User` based on table `users` and generating options
		// g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),

		// Generate struct `Customer` based on table `customer` and generating options
		// customer table may have a tags column, it can be JSON type, gorm/gen tool can generate for your JSON data type
		// g.GenerateModel("customer", gen.FieldType("tags", "datatypes.JSON")),
	)
	//g.ApplyBasic(
	//	// Generate structs from all tables of current database
	//	g.GenerateAllTable()...,
	//)
	// Generate the code
	g.Execute()
}
