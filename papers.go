package papers

import (
	"errors"
	"strconv"
)

var (
	ErrSnils                        = errors.New("Incorrect SNILS")
	ErrSnilsCotainInvalidCharacters = errors.New("SNILS must consist only of numbers")
	ErrINN                          = errors.New("Incorrect INN")
)

type Papers interface {
	ValidateSnils(snils string) error
	ValidateInn(inn string) error
}

func ValidateSnils(snils string) error {
	switch len([]rune(snils)) {
	case 9:
		err := checkMinValueSNILS(snils)
		return err
	case 11:
		err := checkMinValueSNILS(snils)
		if err != nil {
			return err
		}
		var sum = 0
		for i := 0; i < 9; i++ {
			number, err := strconv.Atoi(string(snils[i]))
			if err != nil {
				return ErrSnilsCotainInvalidCharacters
			}
			sum += number * (9 - i)
		}
		var checkDigit = 0
		if sum < 1 {
			return ErrSnils
		}
		if sum < 100 {
			conterSum, err := strconv.Atoi(snils[9:])
			if err != nil {
				return ErrSnilsCotainInvalidCharacters
			}
			if sum == conterSum {
				return nil
			}
			return ErrSnils
		} else {
			checkDigit = sum % 101
			if checkDigit == 100 {
				checkDigit = 0
			}

			conterSum, err := strconv.Atoi(snils[9:])
			if err != nil {
				return ErrSnilsCotainInvalidCharacters
			}
			if checkDigit == conterSum {
				return nil
			}
		}

	}
	return ErrSnils
}

func ValidateInn(inn string) error {
	switch len([]rune(inn)) {
	case 10:
		n10 := checkDigit(inn, []int{2, 4, 10, 3, 5, 9, 4, 6, 8})
		inn9, err := strconv.Atoi((string(inn[9])))
		if err != nil {
			return ErrINN
		}
		if n10 == inn9 {
			return nil
		}
	case 12:
		n11 := checkDigit(inn, []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8})
		n12 := checkDigit(inn, []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8})
		inn10, err := strconv.Atoi(string(inn[10]))
		if err != nil {
			return ErrINN
		}
		inn12, err := strconv.Atoi(string(inn[11]))
		if err != nil {
			return ErrINN
		}
		if (n11 == inn10) && (n12 == inn12) {
			return nil
		}
	}

	return ErrINN
}

func checkDigit(inn string, arr []int) int {
	var n = 0
	for i := range arr {
		number, err := strconv.Atoi(string(inn[i]))
		if err != nil {
			return -1
		}
		n += arr[i] * number
	}
	return n % 11 % 10
}

func checkMinValueSNILS(snils string) error {
	number, err := strconv.Atoi(string(snils[:2]))
	if err != nil {
		return ErrSnilsCotainInvalidCharacters
	}
	if number > 1 {
		return nil
	}
	number, err = strconv.Atoi(string(snils[2:5]))
	if err != nil {
		return ErrSnilsCotainInvalidCharacters
	}
	if number > 1 {
		return nil
	}
	number, err = strconv.Atoi(string(snils[2:5]))
	if err != nil {
		return ErrSnilsCotainInvalidCharacters
	}
	if number < 999 {
		return ErrSnilsCotainInvalidCharacters
	}
	return ErrSnils
}
