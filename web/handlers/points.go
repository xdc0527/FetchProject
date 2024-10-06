package handlers

import (
	"FetchProject/internal"
	"FetchProject/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func AddPoints(c *gin.Context) {
	var req api.AddPointsReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrapf(err, "invalid req")},
		)
		return
	}

	internal.AddTransaction(&req)

	c.JSON(http.StatusOK, gin.H{})
}

func SpendPoints(c *gin.Context) {
	var req api.SpendPointsReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrapf(err, "invalid req")},
		)
		return
	}

	valid, infos := internal.SpendPoints(&req)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the user doesn’t have enough points"})
		return
	}

	payInfos := make([]*api.PayerInfo, 0, len(infos))
	for payer, points := range infos {
		payInfos = append(payInfos, &api.PayerInfo{
			Payer:  payer,
			Points: points,
		})
	}

	c.JSON(http.StatusOK, gin.H{"payer_infos": payInfos})
}

func GetPointsBalance(c *gin.Context) {
	payerBalance := internal.GetBalance()

	c.JSON(http.StatusOK, gin.H{"payer_balance": payerBalance})
}
