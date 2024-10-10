package internal

import "FetchProject/pkg/api"

var (
	payerInfos   = make(map[string]int32, 0)
	transactions = make([]*api.PayerInfo, 0)
)

func AddTransaction(info *api.AddPointsReq) {
	transactions = append(transactions, &info.PayerInfo)
	if points, ok := payerInfos[info.Payer]; ok {
		payerInfos[info.Payer] = points + info.Points
	} else {
		payerInfos[info.Payer] = info.Points
	}
}

func SpendPoints(info *api.SpendPointsReq) (bool, map[string]int32) {
	index := 0
	spending := info.Points
	// deduct points from least recently added
	for i, transaction := range transactions {
		if transaction.Points <= spending {
			payerInfos[transaction.Payer] -= transaction.Points
			spending -= transaction.Points
			index++
		} else {
			payerInfos[transaction.Payer] -= spending
			transactions[i].Points -= spending
			spending = 0
			break
		}
	}
	if spending > 0 {
		return false, nil
	}
	// remove points block that has been used up
	transactions = transactions[index:]
	return true, payerInfos
}

func GetBalance() map[string]int32 {
	return payerInfos
}
