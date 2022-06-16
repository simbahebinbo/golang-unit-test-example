package gostub_demo

import "io/ioutil"

var (
	configFile = "config.json"
	maxNum     = 10
)

func GetConfig() ([]byte, error) {
	return ioutil.ReadFile(configFile)
}

func ShowNumBer() int {
	return maxNum
}
