/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         basic.go
@ Create Time:  2020/4/9 10:53
@ Software:     GoLand
*/

package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var (
	province, city, userAgent []string
)

func GetProvince() (province []string) {
	if buffer, err := ioutil.ReadFile("dataset/province.json"); err != nil {
		log.Println(err.Error())
	} else {
		if err = json.Unmarshal(buffer, &province); err != nil {
			log.Println(err.Error())
		}
	}
	return
}

func GetCity() (city []string) {
	if buffer, err := ioutil.ReadFile("dataset/city.json"); err != nil {
		log.Println(err.Error())
	} else {
		if err = json.Unmarshal(buffer, &city); err != nil {
			log.Println(err.Error())
		}
	}
	return
}

func GetUserAgent() (ua []string) {
	var userAgent map[string][]string

	if buffer, err := ioutil.ReadFile("dataset/user_agent.json"); err != nil {
		log.Println(err.Error())
	} else {
		if err = json.Unmarshal(buffer, &userAgent); err != nil {
			log.Println(err.Error())
		} else {
			for _, val := range userAgent {
				ua = append(ua, val...)
			}
		}
	}
	return
}

func init() {
	province = GetProvince()
	city = GetCity()
	userAgent = GetUserAgent()
}
