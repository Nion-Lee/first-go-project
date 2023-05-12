// package main

// import (
// 	"first-go-project/routers"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	routers.RegisterRoutes(router)

//		router.Run(":8080")
//	}
package main

import (
	"first-go-project/repositories"
	"fmt"
)

func main() {
	repositories.InitializeDB()

	fmt.Println("Done")
}
