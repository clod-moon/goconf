package iniconf

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	conf := InitConfig("./conf/conf.ini")
	username := conf.GetValue("database", "username")
	fmt.Println(username) //root
	conf.DeleteValue("database", "username")
	username = conf.GetValue("database", "username")
	if len(username) == 0 {
		fmt.Println("username is not exists") //this stdout username is not exists
	}
	conf.SetValue("database", "username", "widuu")
	username = conf.GetValue("database", "username")
	fmt.Println(username) //widuu

	data := conf.GetAllSetion()
	fmt.Println(data)
}
