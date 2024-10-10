package api

type SpendPointsRsp struct {
	PayerInfos []*PayerInfo `json:"payer_infos"`
}

type GetPointsBalanceRsp struct {
	PayerBalance map[string]int32 `json:"payer_balance"`
}
