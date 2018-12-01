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
	c.ReadList()
	return c
}

//To obtain corresponding value of the key values
func (c *Config) GetValue(section, name string) string {
	c.ReadList()
	conf := c.ReadList()
	for _, v := range conf {
		for key, value := range v {
			if key == section {
				return value//[name]
			}
		}
	}
	return ""
}

//Set the corresponding value of the key value, if not add, if there is a key change
//设置键值的对应值，如果没有添加，如果有键更改
func (c *Config) SetValue(section, key, value string) bool {
	c.ReadList()
	data := c.conflist
	var ok bool
	var index = make(map[int]bool)
	var conf = make(map[string]map[string]string)
	for i, v := range data {
		_, ok = v[section]
		index[i] = ok
	}

	i, ok := func(m map[int]bool) (i int, v bool) {
		for i, v := range m {
			if v == true {
				return i, true
			}
		}
		return 0, false
	}(index)

	if ok {
		c.conflist[i][section][key] = value
		return true
	} else {
		conf[section] = make(map[string]string)
		conf[section][key] = value
		c.conflist = append(c.conflist, conf)
		return true
	}

	return false
}

//Update the configuration file at the same time
//更新内存的时候,同时更新配置文件
//暂未实现，等待后续
func (c *Config) SetValueToFile(section, key, value string) bool {
	if !c.SetValue(section,key,value){
		return false
	}
}

//Delete the corresponding key values
//删除相应的键值
func (c *Config) DeleteValue(section, name string) bool {
	for i, v := range c.Conflist {
		for key, _ := range v {
			if key == section {
				delete(c.Conflist[i][key], name)
				return true
			}
		}
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
	c.conflist = make(map[string]map[string]string)
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
				c.conflist[section] = sectionMap
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
	c.conflist[section] = sectionMap
	return c.conflist
}

//获取所有配置项
//List all the configuration file
func (c *Config) GetAllSection() map[string]map[string]string{
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
	for _, v := range c.conflist {
		for k, _ := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}
