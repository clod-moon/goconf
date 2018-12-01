/**
 * Read the configuration file
 *
 * @copyright           (C) 2014  clodmoon
 * @lastmodify          2018-11-30
 * @website		https://github.com/clod-moon
 *
 */

package iniconf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	filepath string                         //your ini file path directory+file
	Conflist map[string]map[string]string //configuration information slice
}

const (
	RENAME  = "rename"
	REMOVE  = "remove"
	CHANGE  = "change"
)
//Create an empty configuration file
func InitConfig(filepath string) *Config {
	c := new(Config)
	c.filepath = filepath
	c.readList()
	return c
}

//To obtain corresponding value of the key values
func (c *Config) GetValue(section, name string) string {
	_,ok := c.Conflist[section][name]
	if ok{
		return c.Conflist[section][name]
	}else{
		return ""
	}
}

//Set the corresponding value of the key value, if not add, if there is a key change
//设置键值的对应值，如果没有添加，如果有键更改
func (c *Config) SetValue(section, key, value string){

	_,ok := c.Conflist[section]
	if ok {
		c.Conflist[section][key] = value
	}else{
		c.Conflist[section] = make(map[string]string)
		c.Conflist[section][key] = value
	}
}

//Update the configuration file at the same time
//更新内存的时候,同时更新配置文件
//暂未实现，等待后续
func (c *Config) SetValueToFile(section, key, value string) {
	c.SetValue(section,key,value)
}

//Delete the corresponding key values
//删除相应的键值
func (c *Config) DeleteValue(section, name string) bool {
	//for i, v := range c.Conflist {
	//	for key, _ := range v {
	//		if key == name {
	//			delete(c.Conflist[i][key], name)
	//			return true
	//		}
	//	}
	//}
	_,ok := c.Conflist[section][name]
	if ok{
		delete(c.Conflist[section],name)
		return true
	}else{
		return true
	}

	return false
}

// delete the corresponding key value and update the configuration file at the same time
//删除相应的键值，同时更新配置文件
//暂未实现，等待后续
func (c *Config) DeleteValueToFile(section, name string) bool {
	value := c.GetValue(section,name)
	c.DeleteValue(section,name)
	return false
}

//获取所有配置项
//List all the configuration file
func (c *Config) readList() map[string]map[string]string {
	file, err := os.Open(c.filepath)
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()
	c.Conflist = make(map[string]map[string]string)
	var section string
	var sectionMap map[string]string
	isFirstSection := true
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				CheckErr(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case string(line[0]) == "#":	//增加配置文件备注
		case line[0] == '[' && line[len(line)-1] == ']':
			if !isFirstSection{
				c.Conflist[section] = sectionMap
			}else{
				isFirstSection = false
			}
			section = strings.TrimSpace(line[1 : len(line)-1])
			sectionMap = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			if i == -1 {
				continue
			}
			value := strings.TrimSpace(line[i+1 : len(line)])
			sectionMap[strings.TrimSpace(line[0:i])] = value
		}
	}
	c.Conflist[section] = sectionMap
	return c.Conflist
}

//获取所有配置项
//List all the configuration file
func (c *Config) GetAllSetion() map[string]map[string]string{
	return c.Conflist
}

func CheckErr(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Notfound this error"
}


//Ban repeated appended to the slice method
func (c *Config) uniquappend(conf string) bool {
	for _, v := range c.Conflist {
		for k, _ := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}

