package main

import (
	"fmt"
	"go-viper/core"
	"go-viper/global"
)

func main() {
	global.YAN_VP = core.Viper() // 初始化Viper
	fmt.Println(global.YAN_VP.AllKeys())
	fmt.Println(global.YAN_VP.Get("mysql.path"))
}
