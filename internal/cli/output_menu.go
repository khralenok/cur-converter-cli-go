package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/currecy-converter/internal/commands"
)

func OutputMenu(rates *map[string]float64, baseCur *string) (bool, error){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("----------------")
	fmt.Printf("THE MENU:\n")
	fmt.Println("----------------")
	fmt.Println("1. Calculate exchange rate")
	fmt.Println("2. Switch your base currency")
	fmt.Println("3. Exit the app")
	fmt.Print("Type the number of an action: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)

	if err != nil {
		return false, errors.New("invalid input: please enter number from 1 to 3")
	}
	
	switch choice {
		case 1:
			err := commands.CalcExchangeRate(*rates, baseCur)

			if err != nil {
				return false, err
			}

			return false, nil
		case 2:
			err := commands.SwitchBaseCur(*rates, baseCur)

			if err != nil {
				return false, err
			}

			OutputRates(baseCur)

			return false, nil
		case 3:
			return true, nil
		default:
			return false, errors.New("invalid input: please enter number from 1 to 3")
	}
}
