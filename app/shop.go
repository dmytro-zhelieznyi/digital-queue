package main

import (
	"github.com/fatih/color"
	"time"
)

type Shop struct {
	NumOfEmployees   int
	EmployeeDoneChan chan *Employee
	ClientsChan      chan *Client
	WorkTime         time.Duration
	Open             bool
}

func initShop(numOfEmployees int, workTime time.Duration) Shop {
	return Shop{
		NumOfEmployees:   numOfEmployees,
		EmployeeDoneChan: make(chan *Employee, numOfEmployees),
		ClientsChan:      make(chan *Client, 10),
		WorkTime:         workTime,
		Open:             true,
	}
}

func (shop *Shop) work() {
	shop.checkEndOfWorkingDay()
	runClientGenerator(shop)
	runEmployees(shop)
}

func (shop *Shop) stop() {
	color.Blue("Shop is closing...")
	shop.Open = false
	for i := 0; i < shop.NumOfEmployees; i++ {
		<-shop.EmployeeDoneChan
	}
	close(shop.EmployeeDoneChan)
	closeMain <- true
}

func (shop *Shop) checkEndOfWorkingDay() {
	go func() {
		<-time.After(shop.WorkTime)
		shop.stop()
	}()
}
