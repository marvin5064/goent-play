package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marvin5064/goent-play/ent"
	"github.com/marvin5064/goent-play/ent/migrate"
	"github.com/spf13/viper"
)

func main() {
	viper.BindEnv("MYSQL_HOST", "MYSQL_HOST")
	viper.BindEnv("MYSQL_PORT", "MYSQL_PORT")
	viper.BindEnv("MYSQL_DB", "MYSQL_DB")
	viper.BindEnv("MYSQL_USERNAME", "MYSQL_USERNAME")
	viper.BindEnv("MYSQL_PASSWORD", "MYSQL_PASSWORD")

	store, err := ent.Open("mysql",
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
			viper.GetString("MYSQL_USERNAME"),
			viper.GetString("MYSQL_PASSWORD"),
			viper.GetString("MYSQL_HOST"),
			viper.GetString("MYSQL_PORT"),
			viper.GetString("MYSQL_DB"),
		))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// run the auto migration tool.
	if err := store.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
