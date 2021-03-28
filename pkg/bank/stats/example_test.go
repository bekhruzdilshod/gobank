package stats

import (
	"fmt"
	"github.com/bekhruzdilshod/gobank/pkg/bank/types"
)


func ExampleAvg() {

	payments := []types.Payment{
		{
			Amount: 100,
			Category: "Common",
		},
		{
			Amount: 100,
		},
		{
			Amount: 100,
		},
	}

	result := Avg(payments)
	fmt.Println(result)

	// Output: 100

}