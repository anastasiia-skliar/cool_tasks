package main

import (
	"fmt"
	"github.com/tsenart/vegeta/lib"
	"math"
	"time"
)

func main() {
	rate := uint64(300) // per second
	duration := 5 * time.Second
	login := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "https://cool-tasks.herokuapp.com/v1/login?login=admin&password=admin",
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(login, rate, duration, "") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Println("Max latency:", metrics.Latencies.Max)
	fmt.Println("Requests:", metrics.Requests)
	fmt.Println("Rate:", math.Round(metrics.Rate))
	fmt.Println("BytesIn:", metrics.BytesIn)
	fmt.Println("BytesOut:", metrics.BytesOut)
	fmt.Println("Errors:", metrics.Errors)
	fmt.Println("StatusCodes: ", metrics.StatusCodes)
	fmt.Println("Success: ", metrics.Success)
}
