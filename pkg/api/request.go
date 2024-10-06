package api

type AddPointsReq struct {
	PayerInfo
	Timestamp string `json:"timestamp" binding:"required"`
}

type SpendPointsReq struct {
	Points int32 `json:"points" binding:"required"`
}
