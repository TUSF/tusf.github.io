package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/TUSF/base/dozenal"
	"honnef.co/go/js/dom"
)

func main() {
	document := dom.GetWindow().Document()
	input := document.GetElementByID("input").(*dom.HTMLInputElement)
	output := document.GetElementByID("output")

	input.AddEventListener("keydown", false, func(e dom.Event) {
		ke := e.(*dom.KeyboardEvent)
		if ke.KeyCode == '\r' {
			output.SetTextContent(convert(input.Value))
			ke.PreventDefault()
		}
	})
}

func convert(s string) string {
	var INT big.Int
	var RAT big.Rat

	s = strings.TrimSpace(s)

	if INT, t := INT.SetString(s, 0); t {
		//First, assume it's a plain integer.
		return dozenal.Amer.BigInt(INT)

	} else if RAT, t := RAT.SetString(s); t {
		//Second, assume it's a fraction. ("12/7")
		return dozenal.Amer.BigRat(RAT)

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

					// 1.123 = 1 + 123/1000
					// So, treat everything before the point as an integer
					// then feed the number after the point, divided by the next power of 10.
					//
					// Of course, once you convert it into dozenal, even a simple decimal like that becomes huge.
					// "1.123" becomes 1;15[…], followed by an infinitely repeating sequence of 50+ digits
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
					return dozenal.Amer.BigRat(RAT)
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
