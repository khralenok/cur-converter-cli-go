package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/currecy-converter/utils"
)

func SwitchBaseCur(rates map[string]float64, baseCur *string)(error){
	var curToSwitch string
	var err error
	
	for {
		curToSwitch, err = getCurToSwitch(rates, *baseCur)
	
		if err != nil {
			fmt.Println(err)
			continue
		}

		break
	} 

	*baseCur = curToSwitch

	return nil
}

func getCurToSwitch(rates map[string]float64, baseCur string) (string, error){
	var curToSwitch string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type name of currency you want switch to: ")
	curToSwitch, err := reader.ReadString('\n')
	curToSwitch = strings.TrimSpace(curToSwitch)

	if err != nil {
		return "", errors.New("can't decode your input, please, try again")
	}

	return utils.ValidateCur(curToSwitch, baseCur)
}

