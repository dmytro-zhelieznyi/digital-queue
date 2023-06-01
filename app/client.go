package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

type Client struct {
	Name string
}

func runClientGenerator(shop *Shop) {
	go func(clientNumber int) {
		for {
			if !shop.Open {
				color.Blue("***** wait while last clients take orders (%d) *****", len(shop.ClientsChan))
				close(shop.ClientsChan)
				return
			}

			client := Client{Name: fmt.Sprintf("Client#%d", clientNumber)}
			clientNumber++
			goToShopTime := getRandomMs() / 2

			select {
			case shop.ClientsChan <- &client:
				color.Cyan("%s went to shop with delay %dms", client.Name, goToShopTime)
			default:
				color.Red("%s decided come next time since queue is too big", client.Name)
			}

			time.Sleep(time.Millisecond * time.Duration(goToShopTime))
		}
	}(1)
}
