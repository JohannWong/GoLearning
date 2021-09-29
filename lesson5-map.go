package main

import "fmt"

func lesson5() {
	num_map := map[string]string{"1": "draft", "2": "in progress", "3": "done"}
	for key, value := range num_map {
		fmt.Println(key, value)
	}

	var country_map map[string]string
	country_map = make(map[string]string)
	country_map["法国"] = "巴黎"
	country_map["意大利"] = "罗马"
	country_map["德国"] = "柏林"
	country_map["英国"] = "伦敦"
	for country_name, capital_name := range country_map {
		fmt.Println(country_name, capital_name)
	}
	us_capital, us_capital_flag := country_map["美帝"]
	if us_capital_flag {
		fmt.Println(us_capital)
	} else {
		fmt.Println(us_capital + "没有美帝的首都信息")
	}

	country_map_2 := make(map[string]string)
	for country_name, capital_name := range country_map {
		country_map_2[country_name] = capital_name
	}
	fmt.Println(country_map_2)
	delete(country_map_2, "德国")
	fmt.Println(country_map_2)
}
