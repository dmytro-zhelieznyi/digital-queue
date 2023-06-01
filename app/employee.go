package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

type Employee struct {
	Name string
	Work bool
}

func (employee *Employee) work(shop *Shop) {
	go func() {
		color.Yellow("%s has started working!", employee.Name)
		for {
			select {
			case client := <-shop.ClientsChan:
				if !employee.Work {
					timeBackToWork := 200
					color.Yellow("%s is going back to work... %dms", employee.Name, timeBackToWork)
					time.Sleep(time.Millisecond * time.Duration(timeBackToWork))
				}

				timeToTakeStuff := getRandomMs() * 2
				color.Yellow("%s is tacking stuff from stock for %s... %dms",
					employee.Name, client.Name, timeToTakeStuff)
				time.Sleep(time.Millisecond * time.Duration(timeToTakeStuff))
			default:
				timeForBreak := 500
				time.Sleep(time.Millisecond * time.Duration(timeForBreak))
				color.Yellow("%s has coffee break (%dms) since there is no clients in the shop.",
					employee.Name, timeForBreak)
				employee.Work = false
			}
			if !shop.Open && len(shop.ClientsChan) == 0 {
				employee.stop(shop)
				return
			}
		}
	}()
}

func (employee *Employee) stop(shop *Shop) {
	color.Blue("%s has finished all work and going to go home.", employee.Name)
	shop.EmployeeDoneChan <- employee
}

func runEmployees(shop *Shop) {
	for i := 1; i <= shop.NumOfEmployees; i++ {
		employee := Employee{
			Name: fmt.Sprintf("Employee#%d", i),
			Work: true,
		}
		employee.work(shop)
	}
}
