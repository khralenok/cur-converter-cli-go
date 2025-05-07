package utils

import (
	"errors"

	"example.com/currecy-converter/internal/api"
)

func ValidateCur(curToValidate string, baseCur string)(string,error){
	if curToValidate == baseCur {
		return "", errors.New("it's your current base currency. please choose another one")
	}

	availableCur, err := api.FetchAvailableCurrencies()

	if err != nil {
		panic("can't fetch the data")
	}

	if availableCur[curToValidate] == "" {
		return "", errors.New("there is no such currency. Check the list above for available options")
	}

	return curToValidate, nil
}