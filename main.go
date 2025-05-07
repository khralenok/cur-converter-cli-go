package main

import (
	"fmt"

	"example.com/currecy-converter/internal/cli"
)

func main(){
	baseCurrency := "USD"

	for {
			var isDone bool

			rates, err := cli.OutputRates(&baseCurrency)

			if err != nil{
				fmt.Println(err)
				continue
			}

			for {
				isDone, err = cli.OutputMenu(&rates, &baseCurrency)

				if err != nil{
					fmt.Println(err)
					continue
				}

				if isDone {
					break
				}
			}

			if isDone {
				break
			}
		}

	fmt.Println("Thank you you for using our app!")	
}