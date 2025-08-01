package ito_user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UsersRouter struct{}

func (s *UsersRouter) InitUsersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	usersRouter := Router.Group("users").Use(middleware.OperationRecord())
	usersRouterWithoutRecord := Router.Group("users")
	usersRouterWithoutAuth := PublicRouter.Group("users")
	{
		usersRouter.POST("createUsers", usersApi.CreateUsers)
		usersRouter.DELETE("deleteUsers", usersApi.DeleteUsers)
		usersRouter.DELETE("deleteUsersByIds", usersApi.DeleteUsersByIds)
		usersRouter.PUT("updateUsers", usersApi.UpdateUsers)
	}
	{
		usersRouterWithoutRecord.GET("findUsers", usersApi.FindUsers)
		usersRouterWithoutRecord.GET("getUsersList", usersApi.GetUsersList)
	}
	{
		usersRouterWithoutAuth.GET("getUsersPublic", usersApi.GetUsersPublic)
	}
}
