package papers

import (
	"fmt"
	"regexp"
	"strconv"
)

type SNILS string
type INN string

func ValidateSnils(snilsInPut string) (SNILS, error) {
	snils := ParseSNILS(snilsInPut)
	switch len([]rune(snils)) {
	case 9:
		err := snils.checkMinValueSNILS()
		return snils, err
	case 11:
		err := snils.checkMinValueSNILS()
		if err != nil {
			return snils, err
		}
		var sum = 0
		for i := 0; i < 9; i++ {
			number, err := strconv.Atoi(string(string(snils[i])))
			if err != nil {
				return snils, ErrSnilsCotainInvalidCharacters
			}
			sum += number * (9 - i)
		}
		var checkDigit = 0
		if sum < 1 {
			return snils, ErrSnils
		}
		if sum < 100 {
			conterSum, err := strconv.Atoi(string(snils[9:]))
			if err != nil {
				return snils, ErrSnilsCotainInvalidCharacters
			}
			if sum == conterSum {
				return snils, nil
			}
			return snils, ErrSnils
		} else {
			checkDigit = sum % 101
			if checkDigit == 100 {
				checkDigit = 0
			}

			conterSum, err := strconv.Atoi(string(snils[9:]))
			if err != nil {
				return snils, ErrSnilsCotainInvalidCharacters
			}
			if checkDigit == conterSum {
				return snils, nil
			}
		}

	}
	return snils, ErrSnils
}

func ParseSNILS(s string) SNILS {
	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	sResalt := re.ReplaceAllString(s, "")
	return SNILS(sResalt)
}

func (s SNILS) checkMinValueSNILS() error {
	number, err := strconv.Atoi(string(s[:2]))
	if err != nil {
		return ErrSnilsCotainInvalidCharacters
	}
	if number > 1 {
		return nil
	}
	number, err = strconv.Atoi(string(s[2:5]))
	if err != nil {
		return ErrSnilsCotainInvalidCharacters
	}
	if number > 1 {
		return nil
	}
	number, err = strconv.Atoi(string(s[2:5]))
	if err != nil {
		return ErrSnilsCotainInvalidCharacters
	}
	if number < 999 {
		return ErrSnilsCotainInvalidCharacters
	}
	return ErrSnils
}

func (s SNILS) PrettyPrint() string {
	switch len(s) {
	case 9:
		return fmt.Sprintf("%s-%s-%s", s[:3], s[3:6], s[6:])
	case 11:
		return fmt.Sprintf("%s-%s-%s %s", s[:3], s[3:6], s[6:9], s[9:])
	}
	return fmt.Sprintf("%s", s)
}

func (s SNILS) Minimized() string {
	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	sResalt := re.ReplaceAllString(string(s), "")
	return fmt.Sprintf(sResalt)
}

func (s SNILS) String() string {
	return fmt.Sprintf(string(s))
}

//INN
func ValidateInn(innInput string) error {
	inn := ParseINN(innInput)
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

func checkDigit(inn INN, arr []int) int {
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

func ParseINN(inn string) INN {
	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	innResalt := re.ReplaceAllString(inn, "")
	return INN(innResalt)
}

func (inn INN) String() string {
	return fmt.Sprintf(string(inn))
}
