package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Let us do some channel witchcraft

const MAX_CHICKEN_PRICE float32 = 500

// const MAX_BUTTER_PRICE float32 = 800

func main() {
	var t0 = time.Now()
	fmt.Println("Chicken Finder Pro Max")

	var offerShop = make(chan string)
	var shops = []string{"carrefour.ke", "quickmart.co.ke", "navias.online"}

	for _, shop := range shops {
		go checkChickenPrices(shop, offerShop)
	}

	sendSMS(offerShop)

	fmt.Println()
	fmt.Println("Time Elapsed: ", time.Since(t0))

}

func checkChickenPrices(shop string, offerShop chan string) {
	// Check chicken price every second
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 1000

		// If we find the cheapest chicken around
		if chickenPrice <= MAX_CHICKEN_PRICE {
			// Store it in our channel for the best offer
			offerShop <- shop
			break
		}
	}

}

func sendSMS(offerShop chan string) {
	fmt.Println("We Found a deal at: ", <-offerShop)
}
