package main

import (
	"First_Go_Project/dtos"
	"First_Go_Project/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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

	if guid == parsed {
		fmt.Println("射惹")
	}

	num1, _ := decimal.NewFromString("1.5")
	num2, _ := decimal.NewFromString("1.3")
	num3 := num1.Sub(num2)
	num4, _ := decimal.NewFromString("0.2")

	if num3.Equal(num4) {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
		fmt.Println(num3, "!=", num4)
	}

	dto := dtos.CustomerDTO{
		UID:   "",
		Name:  "",
		Age:   0,
		Email: "",
	}

	jsonBytes, _ := json.Marshal(dto)
	json := string(jsonBytes)
	fmt.Println(jsonBytes)
	entity := services.DesrilizeToStruct(&json)
	fmt.Println(entity)

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
