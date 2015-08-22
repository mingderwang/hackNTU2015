// generated by ginger from go generate -- DO NOT EDIT

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	SvcHost    string
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
	Token      string
	Url        string
}


type TaipeiCityService struct {
}

func (s *TaipeiCityService) getDb(cfg Config) (gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./"+cfg.DbName)
	//db.LogMode(true)
	return db, err
}

func (s *TaipeiCityService) Migrate(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	db.AutoMigrate(&TaipeiCity{})
	return nil
}
func (s *TaipeiCityService) Run(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	taipeiCityResource := &TaipeiCityResource{db: db}

	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	r.GET("/taipeiCity", taipeiCityResource.GetAllTaipeiCitys)
	r.GET("/taipeiCity/:id", taipeiCityResource.GetTaipeiCity)
	r.POST("/taipeiCity", taipeiCityResource.CreateTaipeiCity)
	r.PUT("/taipeiCity/:id", taipeiCityResource.UpdateTaipeiCity)
	r.PATCH("/taipeiCity/:id", taipeiCityResource.PatchTaipeiCity)
	r.DELETE("/taipeiCity/:id", taipeiCityResource.DeleteTaipeiCity)

	r.Run(cfg.SvcHost)

	return nil
}
