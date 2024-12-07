package util

import "strconv"

func Stoi(number string) int {
	acutual_num, err := strconv.Atoi(number)

	if err != nil {
		panic(err)
	}

	return acutual_num
}

func Itos(number int) string {
	return strconv.Itoa(number)
}