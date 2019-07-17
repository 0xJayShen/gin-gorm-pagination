package gin_gorm_pagination

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func DepositAddressList(ctx *gin.Context) {
	type depositAddress struct {
		Address   string    `json:"address"`
		Path      string    `json:"path"`
		Series    string    `json:"series"`
		CreatedAt time.Time `json:"created_at"`
		UserID    string    `json:"user_id"`
	}
	var list []depositAddress
	var walletDB, _ = orm.GetDB("wallet", ctx)
	walletDB.SingularTable(true)
	userId := ctx.DefaultQuery("user_id", "")
	address := ctx.DefaultQuery("address", "")
	if userId != "" {
		walletDB = walletDB.Where("user_id = ?", userId)

	}
	if address != "" {
		walletDB = walletDB.Where("address = ?", address)

	}
	repo := repositories.Repo{
		Ctx:          ctx,
		Result:       &list,
		DB:           walletDB,
		AutoResponse: false,
	}
	repo.Fetch()
	ctx.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"data":       list,
		"pagination": repo.Pagination,
	})

}
