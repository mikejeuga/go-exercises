package utils

import "strconv"

func StringsToInts(numbers []string) []int {
	var values []int
	for _, val := range numbers {
		num, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		values = append(values, num)
	}
	return values
}