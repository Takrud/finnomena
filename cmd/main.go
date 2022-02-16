package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"finno/internal/controller/job"
	"finno/internal/core/coreEntity"
	"finno/internal/core/coreEntity/coreThirdParty"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter time range (e.g. 1D, 1W, 1M, 1Y) you want to view a list of funds")
	}
	inputRange := os.Args[1]
	fmt.Println("============================================================")
	fmt.Printf("=  Time range that you want to view a list of funds is %s  =\n", inputRange)
	fmt.Println("============================================================")

	os.Setenv("GATEWAY_API", "https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-1Y.json")
	
	//Init Core Entity
	coreEntity.InitCoreEntity(coreThirdParty.NewEntity())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	job.GetFundRank(ctx, time.Now().Format("2006-01-02"), inputRange)
}

