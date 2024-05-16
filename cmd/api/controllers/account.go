package controllers

import (
	db "github.com/Touch/datasource"
	"github.com/Touch/services"
	"github.com/gin-gonic/gin"
)

func (r *RestHandler) Login(c *gin.Context) {
	ctx := db.NewContext(c)

	var req services.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, err)
		return
	}

	var response map[string]string
	response, err := services.LoginService(ctx, &req)
	if err != nil {
		ErrorResponse(c, err)
	} else {
		SuccessResponse(c, response, "登录成功")
	}
}

func (r *RestHandler) ShowNft(c *gin.Context) {
	ctx := db.NewContext(c)

	nftInfo, _ := services.GetNftInfo(ctx)

	SuccessResponse(c, nftInfo)

}
