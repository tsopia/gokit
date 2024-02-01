package main

import (
	"fmt"
	"github.com/tsopia/gokit/config"
)

func main() {
	// 加载配置文件
	err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Printf("Error loading config file: %v\n", err)
		return
	}

	// 使用配置
	value := config.GetString("key")
	fmt.Printf("Value from config: %s\n", value)
}
