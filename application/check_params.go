package application

import (
	"github.com/gin-gonic/gin"
)

func BindFromJson(c *gin.Context, ptr any) error {
	if ptr == nil {
		return nil
	}
	check(ptr)
	if err := c.ShouldBindJSON(&ptr); err != nil {
		return err
	}
	return nil
}

func BindFromQueryString(c *gin.Context, ptr any) error {
	if ptr == nil {
		return nil
	}
	check(ptr)
	if err := c.ShouldBindQuery(&ptr); err != nil {
		return err
	}
	return nil
}

func check(ptr any) {
	switch t := ptr.(type) {
	case string:
		if t != "" {
			panic(t)
		}
	case error:
		panic(t.Error())
	}
}
