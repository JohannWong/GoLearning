package main

import (
	"fmt"
)

func lesson18() {
	str_2 := "Chinese: 日本語"
	fmt.Println("The length of str_2 is:", len(str_2))
	for index, rune := range str_2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}

	for i, j, s := 0, 2, "a"; i < 20 && j < 20 && len(s) < 20; i, j, s = j+1, i+2, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

	int_1 := 0
	for int_1 < 20 {
		switch {
		case int_1 <= 10:
			int_1++
			fmt.Println(int_1)
			fallthrough
		case int_1 <= 20:
			int_1 += 2
			fmt.Println(int_1)
			fallthrough
		case int_1 <= 30:
			int_1 += 3
			fmt.Println(int_1)
		}
	}

	int_2 := 0
MARK:
	if int_2 < 10 {
		int_2++
		goto MARK
	}
	fmt.Println(int_2)

}
