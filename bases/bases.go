package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/TUSF/base"
	"github.com/TUSF/base/dozenal"
	"honnef.co/go/js/dom"
)

var BAS base.Formatter

var sex base.Formatter = base.NewFormatter([]string{"0", "1", "2", "3", "4", "5"})

func main() {
	document := dom.GetWindow().Document()
	input := document.GetElementByID("input").(*dom.HTMLInputElement)
	output := document.GetElementByID("output")
	options := document.GetElementByID("base").(*dom.HTMLSelectElement)
	BAS = dozenal.ASCII

	input.AddEventListener("keydown", false, func(e dom.Event) {
		ke := e.(*dom.KeyboardEvent)
		if ke.KeyCode == '\r' {
			output.SetTextContent(convert(input.Value))
			ke.PreventDefault()
		}
	})
	options.AddEventListener("change", false, func(e dom.Event) {
		switch options.Value {
		case "sex":
			BAS = sex
		case "dozascii":
			BAS = dozenal.ASCII
		case "dozamer":
			BAS = dozenal.Amer
		case "dozbrit":
			BAS = dozenal.Brit
		}
		output.SetTextContent(convert(input.Value))
	})
}

func convert(s string) string {
	var INT big.Int
	var RAT big.Rat

	s = strings.TrimSpace(s)

	if INT, t := INT.SetString(s, 0); t {
		//First, assume it's a plain integer.
		return BAS.BigInt(INT)

	} else if RAT, t := RAT.SetString(s); t {
		//Second, assume it's a fraction. ("12/7")
		return BAS.BigRat(RAT)

	} else {
		//Third, assume it's a decimal number. ("10.123")
		if strings.Index(s, ".") > -1 {
			if nums := strings.Split(s, "."); len(nums) == 2 {
				//Convert each side of the decimal point into a big.Int?
				if nums[1] == "" {
					return "Not a valid number. Integers, Fractions or Decimals only."
				}
				if nums[0] == "" {
					RAT.SetInt64(0)
				} else {
					if _, err := strconv.Atoi(nums[0]); err != nil {
						return "Not a valid number. Integers, Fractions or Decimals only."
					}
					if _, err := strconv.Atoi(nums[1]); err != nil {
						return "Not a valid number. Integers, Fractions or Decimals only."
					}
					RAT, t := RAT.SetString(nums[0])
					if !t {
						return "Not a valid number. Integers, Fractions or Decimals only."
					}

					d, t := new(big.Rat).SetString(
						fmt.Sprintf("%s/1%s", nums[1], strings.Repeat("0", len(nums))))
					if !t {
						return "Not a valid number. Integers, Fractions or Decimals only."
					}
					RAT.Add(RAT, d)
					return BAS.BigRat(RAT)
				}
			} else {
				return "Not a valid number. Integers, Fractions or Decimals only."
			}
		} else {
			return "Not a valid number. Integers, Fractions or Decimals only."
		}
	}
	return ""
}
