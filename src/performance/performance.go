package main

import (
	"time"
	"github.com/tsenart/vegeta/lib"
	"fmt"
	"math"
)

func main()  {
	rate := uint64(100) // per second
	duration := 3 * time.Second
	login := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "https://cool-tasks.herokuapp.com/v1/login?login=admin&password=admin",
	})
	//users := vegeta.NewStaticTargeter(vegeta.Target{
	//	Method: "GET",
	//	URL:    "https://cool-tasks.herokuapp.com/v1/users",
	//})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(login, rate, duration, "") {
		metrics.Add(res)
	}
	//for res := range attacker.Attack(users, rate, duration, "") {
	//	metrics.Add(res)
	//}
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
