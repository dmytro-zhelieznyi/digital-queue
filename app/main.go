package main

import (
	"github.com/fatih/color"
	"math/rand"
	"time"
)

var closeMain = make(chan bool)
var numOfEmployees = 3
var shopWorkTime = 15

func main() {
	color.Blue("***** DIGITAL QUEUE START *****")

	shop := initShop(numOfEmployees, time.Second*time.Duration(shopWorkTime))
	shop.work()

	<-closeMain
	close(closeMain)
	color.Blue("***** DIGITAL QUEUE END *****")
}

/*
get random num from 0 to 999
*/
func getRandomMs() int {
	rand.Seed(time.Now().UnixNano())
	delayMs := rand.Int() % 1000
	return delayMs
}
