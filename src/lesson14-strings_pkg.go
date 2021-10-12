package main

import (
	"fmt"
	"strings"
)

func lesson14() {
	var str_1 string = "This is a test. is it?"

	fmt.Println(strings.HasPrefix(str_1, "TH"))
	fmt.Println(strings.HasPrefix(str_1, "Th"))
	fmt.Println(strings.HasSuffix(str_1, "test"))
	fmt.Println(strings.HasSuffix(str_1, "t?"))

	fmt.Println(strings.Contains(str_1, "IS"))
	fmt.Println(strings.Contains(str_1, "is"))

	fmt.Println(strings.Index(str_1, "is"))
	fmt.Println(strings.LastIndex(str_1, "is"))
	fmt.Println(strings.Index(str_1, "good"))

	fmt.Println(strings.Replace(str_1, "test", "testing", -1))

	fmt.Println(strings.Count(str_1, "is"))

	fmt.Println(strings.Repeat(str_1, 3))

	fmt.Println(strings.ToLower(str_1))
	fmt.Println(strings.ToUpper(str_1))

	var str_2 string = " another string "
	fmt.Println(strings.TrimSpace(str_2))
	fmt.Println(strings.Trim(str_2, " "))
	fmt.Println(strings.Trim("td another string td", "td"))
	fmt.Println(strings.TrimLeft(str_2, " "))

	var str_1_list []string = strings.Fields(str_1)
	fmt.Println(str_1_list)
	str_2_list := strings.Split(str_2, " ")
	fmt.Println(str_2_list)
	fmt.Println(strings.Join(str_1_list, "_"))

	str_1_reader := strings.NewReader(str_1)
	var read_byte byte
	for i := 0; i < len(str_1); i++ {
		read_byte, _ = str_1_reader.ReadByte()
		fmt.Println(string(read_byte), read_byte)
	}
}
