package stats

import (
	"github.com/MSHE97/hworks.git/pkg/bank/types"
)

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	sum := 0
	for _, pay := range payments {
		sum += int(pay.Amount)
	}

	return types.Money( sum / len(payments) )
}