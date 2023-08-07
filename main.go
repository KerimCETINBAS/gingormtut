package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/gotuts/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// host=localhost user=postgres password=postgres dbname=elitas port=5432 sslmode=disable TimeZone=Asia/Istanbul
func main() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=blogs port=5432 sslmode=disable TimeZone=Asia/Istanbul"))

	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&types.Blog{})
	db.AutoMigrate(&types.Author{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		blogs := []types.Blog{}
		if err := db.Joins("Author").Find(&blogs).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		fmt.Println(blogs)
		c.JSON(http.StatusOK, blogs)
	})

	r.POST("/", func(c *gin.Context) {

		var blog types.Blog
		c.BindJSON(&blog)

		db.Model(&types.Blog{}).Create(&blog)

	})

	r.Run()
}
