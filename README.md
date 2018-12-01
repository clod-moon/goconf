iniconf
========

**[The official website](https://github.com/clod-moon)**
## 描述

使用iniconf更简单的读取go的ini配置文件以及根据特定格式的各种配置文件。

## 安装方法

	go get github.com/clod-moon/goconf

## 使用方法

>ini配置文件格式样列

	[database]
	username = root
	password = password
	hostname = localhost
	
	[admin]
	username = root
	password = password
	
	[nihao]
	username = root
	password = password

>初始化

	conf := goini.InitConfig("./conf/conf.ini") //iniconf.SetConfig(filepath) 其中filepath是你ini 配置文件的所在位置

>获取单个配置信息

	username := conf.GetValue("database", "username") //database是你的[section]，username是你要获取值的key名称
	fmt.Println(username) //root

>删除一个配置信息

	conf.DeleteValue("database", "username")	//username 是你删除的key
	username = conf.GetValue("database", "username")
	if len(username) == 0 {
		fmt.Println("username is not exists") //this stdout username is not exists
	}

>添加一个配置信息

	conf.SetValue("database", "username", "chun")
	username = conf.GetValue("database", "username")
	fmt.Println(username) //chun 添加配置信息如果存在[section]则添加或者修改对应的值，如果不存在则添加section

>获取所有配置信息

	conf.GetAllSetion() //返回map[string]map[string]string的格式 即setion=>key->value

---

iniconf
========


## About

使用iniconf更简单的读取go的ini配置文件以及根据特定格式的各种配置文件。

<<<<<<< HEAD
## install
=======
## install 
>>>>>>> 0de79bcc56fe3d3b472cbc1ec716489d0774b025

	go get github.com/clod-moon/goconf

## use example

>conf.ini

	[database]
	username = root
	password = password
	hostname = localhost
	
	[admin]
	username = root
	password = password
	
	[nihao]
	username = root
	password = password

>initialize

	conf := goini.InitConfig("./conf/conf.ini") //goini.InitConfig(filepath) filepath = directory+file

>To obtain a single configuration information

	username := conf.GetValue("database", "username") //username is your key you want get the value
	fmt.Println(username) //root

>To delete a configuration information

	conf.DeleteValue("database", "username")	//username is your delete the key
	username = conf.GetValue("database", "username")
	if len(username) == 0 {
		fmt.Println("username is not exists") //this stdout username is not exists
	}

>Add a configuration information

	conf.SetValue("database", "username", "chun")
	username = conf.GetValue("database", "username")
	fmt.Println(username) //chun Adding/section configuration information if there is to add or modify the value of the corresponding, if there is no add section

>Get all the configuration information

	conf.GetAllSetion() //return map[string]map[string]string  example:setion=>key->value



