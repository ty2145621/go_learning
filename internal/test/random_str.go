package main

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var dataVec1 []*DBValue
var dataVec2 []*DBValue

func randomStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func randomMap() map[string]string {
	rand.Seed(time.Now().UnixNano())
	dataMap := map[string]string{}
	for i := 0; i <= 128; i++ {
		dataMap[randomStr(5)] = randomStr(16)
	}
	return dataMap
}

func getDataVec() []*DBValue {
	dataVec := []*DBValue{}
	for i := 0; i <= 500; i++ {
		dataVec = append(dataVec, &DBValue{Data: randomMap()})
	}
	return dataVec
}
