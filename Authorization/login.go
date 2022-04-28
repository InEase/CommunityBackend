package Authorization

import (
	. "AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	// TODO:多次重复登录销毁上一次token，刷新存活时间
	db := database.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	// Start Process
	var user Users
	db.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		ResponseWithNoData(ctx, 1008)
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ResponseWithNoData(ctx, 1009)
		return
	}
	// 发放token
	token, err := ReleaseToken(user)
	if err != nil {
		ResponseWithNoData(ctx, 1010)
		return
	}
	// 返回结果
	Response(ctx, 0, gin.H{"token": token}, StatusMsg(0))
}