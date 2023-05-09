package main

import (
	"First_Go_Project/dtos"
	"First_Go_Project/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	guid, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("生成 UUID 時出錯：%v\n", err)
		return
	}
	ggininder := guid.String()
	fmt.Println(ggininder)

	parsed, err := uuid.Parse(ggininder)
	if err != nil {
		fmt.Printf("解析 UUID 時出錯：%v\n", err)
		return
	}
	fmt.Println(parsed)

	dto := dtos.CustomerDTO{
		UID:   "5566",
		Name:  "Nion",
		Age:   32,
		Email: "niosinmine@rrrr.com",
	}

	jsonBytes, _ := json.Marshal(dto)
	json1 := string(jsonBytes)
	fmt.Println(jsonBytes)
	fmt.Println(json1)
	entity, _ := services.DesrilizeToStruct(&json1)
	fmt.Println("我在這")
	fmt.Println(entity)
	fmt.Println("我結束")

	var dto1 dtos.CustomerDTO
	errrrr := json.Unmarshal(jsonBytes, &dto1)

	fmt.Println(errrrr)
	fmt.Println(dto1)

	a, _ := json.Marshal(dto1)
	fmt.Println(string(a))

	return

	// router := gin.Default()

	// router.GET("/", func(context *gin.Context) {
	// 	context.JSON(http.StatusOK, "RRRRRRRRRR")
	// })

	// router.Run(":8080")

	router := gin.Default()

	// 註冊一個 POST 請求的路由，接收 Customer JSON 數據
	router.POST("/customer", func(c *gin.Context) {
		var customer Customer

		// 解析 JSON 數據並將其綁定到 Customer 結構體
		if err := c.BindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 正常處理後返回 JSON 數據
		c.JSON(http.StatusOK, gin.H{
			"name":    customer.Name,
			"age":     customer.Age,
			"address": customer.Address,
		})
	})

	router.Run(":8080")
}

type Customer struct {
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
