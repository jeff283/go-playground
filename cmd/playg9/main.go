package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Let us do some channel witchcraft

const MAX_CHICKEN_PRICE float32 = 500

const MAX_BUTTER_PRICE float32 = 800

func main() {
	var t0 = time.Now()
	fmt.Println("Deal Finder Pro Max")

	var chickenOffer = make(chan string)
	var butterOffer = make(chan string)

	var shops = []string{"carrefour.ke", "quickmart.co.ke", "navias.online"}

	for _, shop := range shops {
		go checkChickenPrices(shop, chickenOffer)
		go checkButterPrices(shop, butterOffer)
	}

	sendMessage(chickenOffer, butterOffer)

	fmt.Println()
	fmt.Println("Time Elapsed: ", time.Since(t0))

}

func checkButterPrices(shop string, butterOffer chan string) {
	// get butter offers
	for {
		time.Sleep(time.Second * 1)
		var butterPrice = (rand.Float32() * 1000) + 200

		if butterPrice < MAX_BUTTER_PRICE {
			butterOffer <- shop
			break
		}
	}

}

func checkChickenPrices(shop string, chickenOffer chan string) {
	// Check chicken price every second
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 1000

		// If we find the cheapest chicken around
		if chickenPrice <= MAX_CHICKEN_PRICE {
			// Store it in our channel for the best offer
			chickenOffer <- shop
			break
		}
	}

}

func sendMessage(chickenOffer chan string, butterOffer chan string) {
	select {
	case shop := <-chickenOffer:
		fmt.Println("SMS Sent: Chicken Deal Found at: ", shop)
	case shop := <-butterOffer:
		fmt.Println("Email Sent: Butter Deal found at:: ", shop)

	}
}
