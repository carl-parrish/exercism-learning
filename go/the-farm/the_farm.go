package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	numOfCows int
	message string

}

func (e *InvalidCowsError) Error()string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numOfCows, e.message)
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cows int) (float64, error){
	amount, error := fc.FodderAmount(cows)
	if error != nil {
		return 0, error
	} else {
		multiplier, err := fc.FatteningFactor()
		if err != nil {
			return 0, err
		}else {
			ans := (amount / float64(cows)) * multiplier	
			return ans, nil		
		}	
	}
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cows int)(float64, error){
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	} else {
		return DivideFood(fc, cows)
	}
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error{
	switch {
	case cows < 0 :
			return &InvalidCowsError{ numOfCows: cows, message: "there are no negative cows"} 
	case cows == 0 :
		return &InvalidCowsError{ numOfCows: cows, message: "no cows don't need food"}
	default:
		return nil
		}
}

