package Articles

import (
	"AICommunity/Authorization"
	. "AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DefaultList(ctx *gin.Context) {
	db := database.GetDB()

	temp1, exists := ctx.GetQuery("limit")

	limit := 10
	err := error(nil)
	if exists {
		limit, err = strconv.Atoi(temp1)
		if err != nil {
			ResponseWithNoData(ctx, "limit 请传入数值类型")
			return
		}
	}

	offset := 0
	temp2, exists := ctx.GetQuery("offset")
	if exists {
		offset, err = strconv.Atoi(temp2)
		if err != nil {
			ResponseWithNoData(ctx, "offset 请传入数值类型")
			return
		}
	}

	// Start Process
	var articles []Articles

	temp3, exists := ctx.GetQuery("category")
	if exists {
		category, parseError := strconv.Atoi(temp3)
		if parseError != nil {
			ResponseWithNoData(ctx, "category 请传入数值类型")
			return
		}
		db.Limit(limit).Offset(offset).Where("category = ?", category).Order("created_at DESC").Find(&articles)
	} else {
		db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&articles)
	}

	var userid uint = 999999999
	user, status := ctx.Get("user")
	if status {
		userid = user.(Authorization.Users).ID
	}
	Response(ctx, 0, ToListArticleDto(articles, userid), "完成")
}
