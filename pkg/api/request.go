package api

type PayerInfo struct {
	Payer  string `json:"payer" binding:"required"`
	Points int32  `json:"points" binding:"required"`
}

type AddPointsReq struct {
	PayerInfo
	Timestamp string `json:"timestamp" binding:"required"`
}

type SpendPointsReq struct {
	Points int32 `json:"points" binding:"required"`
}
