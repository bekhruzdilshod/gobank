package stats

import (
	"github.com/bekhruzdilshod/gobank/pkg/bank/types"
)


// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {

	totalSum := types.Money(0)

	for _, payment := range payments {
		totalSum += payment.Amount
	}

	return totalSum / types.Money(len(payments))

}