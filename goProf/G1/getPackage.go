package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/zhashkevych/scheduler"
)

func main() {
	ctx := context.Background()

	worker := scheduler.NewScheduler()
	worker.Add(ctx, parseSubscriptionData, time.Second*5)
	worker.Add(ctx, sendStatistics, time.Second*10)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	worker.Stop()
}

func parseSubscriptionData(ctx context.Context) {
	time.Sleep(time.Second * 1)
	fmt.Printf("subscription parsed successfuly at %s\n", time.Now().String())
}

func sendStatistics(ctx context.Context) {
	time.Sleep(time.Second * 5)
	fmt.Printf("statistics sent at %s\n", time.Now().String())
}
