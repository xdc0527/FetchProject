package api

type PayerInfo struct {
	Payer  string `json:"payer" binding:"required"`
	Points int32  `json:"points" binding:"required"`
}

type SpendPointsRsp struct {
	PayerInfos []*PayerInfo `json:"payer_infos"`
}

type GetPointsBalanceRsp struct {
	PayerBalance map[string]int32 `json:"payer_balance"`
}
