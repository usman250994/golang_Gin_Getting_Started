package controllers

import (
	"net/http"

	models "example.com/usman-apis/model"
	services "example.com/usman-apis/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context)  {
	var user models.User

	if err:= ctx.ShouldBindJSON(&user); err != nil {
	ctx.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
	return
}
	err:= uc.UserService.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func  (uc *UserController) GetUser(ctx *gin.Context)  {
	username:= ctx.Param("name")
    user, err :=   uc.UserService.GetUser(&username)
	if err !=nil {
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}
		ctx.JSON(http.StatusOK,gin.H{"message": user})
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	 users, err :=   uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}
		ctx.JSON(http.StatusOK,gin.H{"message": users})
}

func(uc *UserController) UpdateUser(ctx *gin.Context)  {
	var user models.User

	if err:= ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)

	if err !=nil {
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) Delete(ctx *gin.Context)  {
	 username := ctx.Param("name")
	err := uc.UserService.Delete(&username)

	if err !=nil {
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func ( uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute:=rg.Group("/user")
	userroute.POST("/create",uc.CreateUser)
	userroute.GET("/getUser:name",uc.GetUser)
	userroute.GET("/getAll",uc.GetAll)
	userroute.PATCH("/update",uc.UpdateUser)
	userroute.DELETE("/delete:name",uc.Delete)
}