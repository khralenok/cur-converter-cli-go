package cli

import (
	"errors"
	"fmt"
	"math"

	"example.com/currecy-converter/internal/api"
)

func OutputRates(baseCur *string) (map[string]float64, error){
	rates, err := api.FetchRates(*baseCur)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("can't fetch the data")
	}
	
	fmt.Println("----------------")
	fmt.Printf("TODAY %s RATE:\n", *baseCur)
	fmt.Println("----------------")
	
	var index int = 0;
	for cur, price := range rates{
		index++
		if math.Remainder(float64(index), 5) != 0{
			fmt.Printf("%s: %.2f   |   ", cur, price)
		} else {
			fmt.Printf("%s: %.2f\n", cur, price)
		}
	}
	fmt.Println("----------------")

	return rates, nil
}