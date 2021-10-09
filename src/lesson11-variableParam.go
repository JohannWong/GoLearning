package main

import (
	"fmt"
	"strconv"
)

func sum_all(num ...int) (res int) {
	for _, num_v := range num {
		res += num_v
	}
	return res
}

func sum_all_with_convert(num ...interface{}) (res float64) {
	for _, num_v := range num {
		switch num_type := num_v.(type) {
		case int:
			res += float64(num_type)
		case int8:
			res += float64(num_type)
		case int16:
			res += float64(num_type)
		case int32:
			res += float64(num_type)
		case int64:
			res += float64(num_type)
		case float32:
			res += float64(num_type)
		case float64:
			res += num_type
		case string:
			if res_num, err := strconv.ParseFloat(num_type, 64); err == nil {
				res += res_num
			}
		}
	}
	return res
}

func lesson11() {
	fmt.Println(sum_all(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(sum_all_with_convert(1, 1.1, 2.1, 3.1, 4, true, "5.1"))
}
