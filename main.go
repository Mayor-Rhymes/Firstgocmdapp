package main

import (
	"fmt"
	 utilities "utilities/utilities"
	 "log"
	
)









func handleStopper(end *bool){


	fmt.Println("Do you want to stop using the utility program? y/n")
	var answer string

	fmt.Scanln(&answer)
	if answer == "y" || answer == "Y"{
                        
		*end = true
					
	}
}




func main(){


	if bmi, err := utilities.BmiCalculator(82, 1.75); err != nil{
         
		log.Fatal(err)

	} else{

		fmt.Printf("My bmi is %f:\n", bmi)
	}

    


	calcUtil := utilities.Calculator{

		FirstOperand: 10,
		SecondOperand: 20,
	}


	fmt.Println(calcUtil.Add())


	utilitiesArray := []string {"BMI Calculator", "Currency Converter", "Basic Calculator"}





	//Creating the command line interface


	fmt.Println("Hi, My name is MayorRhymes")
	fmt.Println("I am trying to be a more hardworking programming")
	


	var end bool = false

	for {


        fmt.Println("Please choose the utility you want to use: ")
		for index, element := range utilitiesArray{

			fmt.Println(index, element)
		}

	    //Get input from user


		var userEntry int


		fmt.Scanln(&userEntry)
		


		switch userEntry{


		    case 0:
				fmt.Println("You chose BMI Calculator.")
				fmt.Println("This is your friendly bmi calculator")
				fmt.Print("Please first enter your weight: ")
				var weight float64
				fmt.Scanln(&weight)
				fmt.Println()
				fmt.Print("Please enter your height: ")
				var height float64
				fmt.Scanln(&height)
				if answer, err := utilities.BmiCalculator(weight, height); err != nil{

					log.Fatal(err)
				} else{

					fmt.Printf("Your bmi is %.2f\n", answer);
				}

				handleStopper(&end)


				


			case 1:
				fmt.Println("You chose currency converter")
			    currencyArray := []string {"USD", "NGN", "AUD"}

                
				for index, currency := range currencyArray{

					fmt.Println(index, currency)

				}
				var currencyToConvertFrom, currencyToConvertTo int
				var money float64

				fmt.Println("Please enter the currency to convert from: ")


				fmt.Scanln(&currencyToConvertFrom)

				fmt.Println("Please enter the currency to convert to: ")
				fmt.Scanln(&currencyToConvertTo)


				fmt.Println("Please enter amount")
				fmt.Scanln(&money)


				var result = utilities.CurrencyCalculator(currencyArray[currencyToConvertFrom], currencyArray[currencyToConvertTo], money)


				fmt.Printf("%.3f %s is equal to %.3f %s\n", money, currencyArray[currencyToConvertFrom], result, currencyArray[currencyToConvertTo])
				handleStopper(&end)


			case 2:
				fmt.Println("You chose basic calculator")
                var firstOperand, secondOperand float64
				fmt.Print("Please enter the first operand: ")
				fmt.Scanln(&firstOperand)
			
			    fmt.Print("Please enter the second operand: ")	
				fmt.Scanln(&secondOperand)


				

				operationsArray := []string {"add", "sub", "mul", "div"}

				fmt.Println("Please choose the operation you want: ")
				for index, value := range operationsArray {


					fmt.Println(index, value)


				}


				

				var operationNumber int
				fmt.Scanln(&operationNumber)




				calculator := utilities.Calculator{

					FirstOperand: firstOperand,
					SecondOperand: secondOperand,
				}

			
                var result float64
				switch operationNumber {


				   case 0:
					  result = calculator.Add()
					  fmt.Printf("The answer is %.2f\n", result)
					
				   case 1:

					   result = calculator.Subtract()
					   fmt.Printf("The answer is %.2f\n", result)


				   case 2:
					 
					  result = calculator.Multiply()
					  fmt.Printf("The answer is %.2f\n", result)


				   case 3:
					   
					  result = calculator.Divide()
					  fmt.Printf("The answer is %.2f\n", result)



					  

				}

				handleStopper(&end)				



				
				
		}


		if end{

			break
		}
		
	}



}