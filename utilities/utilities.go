package utilities


import (
	"errors"
)



func BmiCalculator(weight, height float64) (result float64, err error) {
    
	if height > weight {

		return 0, errors.New("height cannot be greater than weight")
	}

	result = weight/(height*height)
	return


}


func CurrencyCalculator(currencyToConvertFrom string, currencyToConvertTo string, money float64) (result float64) {



	   if currencyToConvertFrom == "NGN" && currencyToConvertTo == "USD"{

             result = money / 460.28
		   
	   } else if currencyToConvertFrom == "USD" && currencyToConvertTo == "NGN"{

		   result = money * 460.28
	   } else if currencyToConvertFrom == "AUD" && currencyToConvertTo == "NGN"{

		  result = money * 327.43
		  
	   } else if currencyToConvertFrom == "NGN" && currencyToConvertTo == "AUD"{

		  result = money / 327.43
		  
	   } else if currencyToConvertFrom == "AUD" && currencyToConvertTo == "USD"{

		  result = money / 1.41
	   } else if currencyToConvertFrom == "USD" && currencyToConvertTo == "AUD"{

		result = money * 1.41
	   } 

	   return
}


type Calculator struct{

    FirstOperand, SecondOperand float64


}


func (c *Calculator) Add() float64 {


	return c.FirstOperand + c.SecondOperand
}


func (c *Calculator) Subtract() float64 {


	return c.FirstOperand - c.SecondOperand
}


func (c *Calculator) Divide() float64 {


	return c.FirstOperand / c.SecondOperand
}


func (c *Calculator) Multiply() float64 {


	return c.FirstOperand * c.SecondOperand
}













