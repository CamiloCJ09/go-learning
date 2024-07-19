package utils

import (
	"strconv"
	"strings"
)

func TransformStringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func TransformStringListToIntList(s []string) ([]int, error) {
	var result []int
	for _, str := range s {
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

func TransformIntListToStringList(i []int) []string {
	var result []string
	for _, num := range i {
		result = append(result, strconv.Itoa(num))
	}
	return result
}

func ParseStringToInt(s string) ([]int, error) {
	var result []int
	s = strings.Replace(s, "]", "", -1)
	s = strings.Replace(s, "[", "", -1)
	resultStr := strings.Split(s, ",")
	result, err := TransformStringListToIntList(resultStr)
	if err != nil {
		return nil, err
	}

	return result, nil
}
