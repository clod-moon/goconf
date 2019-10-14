package iniconf

import (
	"testing"
)

var (
	conf *Config
	path = "./conf/conf.ini"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestInitConfig(t *testing.T) {

	conf = InitConfig("./conf/conf.ini")
	if conf == nil {
		t.Error("init failed")
	}else {
		t.Log("init ok")
	}
}

func TestGetValue(t *testing.T) {
	var tests = []struct{
		baseKey string
		key string
		value string
	}{
		{"database","username","root"},
		{"admin","username","root"},
		{"nihao","username","root"},
	}

	for _,test := range tests{
		value := conf.GetValue(test.baseKey,test.key)
		if value!= test.value{
			t.Errorf("expect value is %s,but get value is %s\n",test.value,value)
		}
	}
}
