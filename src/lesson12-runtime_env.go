package main

import (
	"fmt"
	"os"
	"runtime"
)

func lesson12() {
	var goos string = runtime.GOOS
	fmt.Println("op_sys:", goos)
	path := os.Getenv("PATH")
	fmt.Println("sys_path:", path)
	env_list := os.Environ()
	for env_index, env_ele := range env_list {
		fmt.Println(env_index, ":", env_ele)
	}
}
