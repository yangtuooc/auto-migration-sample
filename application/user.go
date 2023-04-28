package application

import (
	"auto-migration-sample/application/response"
	"auto-migration-sample/common"
	"auto-migration-sample/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {

	queryParam := c.Query("name")
	user := domain.EmptyUser()
	err := common.UseDB().Where("name=?", queryParam).First(&user).Error
	if err != nil {
		response.Fail(http.StatusInternalServerError, err.Error(), c)
	}
	response.Success(c, user)
}

func RegisterUser(c *gin.Context) {

	user := domain.EmptyUser()
	err := BindFromJson(c, &user)
	if err != nil {
		response.Fail(http.StatusBadRequest, err.Error(), c)
	}

	err = common.UseDB().Save(&user).Error
	if err != nil {
		response.Fail(http.StatusInternalServerError, err.Error(), c)
	}
	response.Success(c, map[string]any{"id": user.ID})
}
