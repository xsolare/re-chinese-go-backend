package main

import (
	"github.com/xsolare/re-chinese-go-backend/config"
	"github.com/xsolare/re-chinese-go-backend/models"

	//"net/http"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

var ORM orm.Ormer

func init() {
	config.GetEnvVariables()
    models.ConnectToDb()
    ORM = models.GetOrmObject()
}

func main() {
    router := gin.Default()
    // router.POST("/createUser", createUser)
    // router.GET("/readUsers", readUsers)
    // router.PUT("/updateUser", updateUser)
    // router.DELETE("/deleteUser", deleteUser)
    router.Run(":3000")
}

// func createUser(c *gin.Context) {
//     var newUser models.Users
//     c.BindJSON(&newUser)
//     _, err := ORM.Insert(&newUser)
//     if err == nil {
//         c.JSON(http.StatusOK, gin.H{
//             "status": http.StatusOK, 
//             "email": newUser.Email,
//             "user_name": newUser.UserName, 
//             "user_id": newUser.UserId})
//     } else {
//         c.JSON(http.StatusInternalServerError, 
//             gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
//     } 
// }