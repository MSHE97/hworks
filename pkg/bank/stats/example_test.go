package stats

import (
	"fmt"

	"github.com/MSHE97/hworks.git/pkg/bank/types")


func ExampleAvg() {

	payments := []types.Payment{
		{
			ID: 1,
			Amount: 100_00,
			Category: "cafe",
		},
		{
			ID: 2,
			Amount: 50_00,
			Category: "market",
		},{
			ID: 3,
			Amount: 125_00,
			Category: "transport",
		},
		{
			ID: 4,
			Amount: 25_00,
			Category: "cafe",
		},
	}

	fmt.Println( Avg(payments) )

	// Output: 7500
}