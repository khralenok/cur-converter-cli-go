package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/currecy-converter/utils"
)

func CalcExchangeRate(rates map[string]float64, baseCur *string) error {
	var curToConvert string
	var amountToConvert float64
	var err error
	
	for {
		curToConvert, err = getCurToConvert(rates, *baseCur)
	
		if err != nil {
			fmt.Println(err)
			continue
		}

		break
	} 

	for {
		amountToConvert, err = getAmountToConvert(baseCur)
	
		if err != nil {
			fmt.Println(err)
			continue
		}

		break
	} 

	result := amountToConvert * rates[curToConvert]

	fmt.Printf("%.2f %s is equal to %.2f %s\n", amountToConvert, *baseCur, result, curToConvert)

	return nil
}

func getCurToConvert(rates map[string]float64, baseCur string) (string, error){
	var curToConvert string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type name of currency you want exchange to: ")
	curToConvert, err := reader.ReadString('\n')
	curToConvert = strings.TrimSpace(curToConvert)

	if err != nil {
		return "", errors.New("can't decode your input, please, try again")
	}

	return utils.ValidateCur(curToConvert, baseCur)
}

func getAmountToConvert(baseCur *string)(float64, error){
	var amountToConvert float64
	var input string
	var err error

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Type amount of %s you want to change: ", *baseCur)


	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	amountToConvert, err = strconv.ParseFloat(input, 64)

	if err != nil || amountToConvert <= 0.0 {
		return 0.0, errors.New("invalid input. Amount should be positive number")
	}

	return amountToConvert, nil
}

