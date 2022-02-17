package cmd

import (
	"context"
	"fmt"
	"flag"
	"os"
	"time"

	"finnomena/internal/controller/job"
	"finnomena/internal/core/coreEntity"
	"finnomena/internal/core/coreEntity/coreThirdParty"
)

func Execute() {
	inputRange := flag.String("t", "1Y", "time range (e.g. 1D, 1W, 1M, 1Y) you want to view a list of funds")
	flag.Parse()

	fmt.Println("=========================================================================")
	fmt.Println("=\t\t\t\t\t\t\t\t\t=")
	fmt.Printf("=\tToday is %s\t\t\t\t\t\t=\n", time.Now().Format("2006-01-02"))
	fmt.Printf("=\tTime range that you want to view a list of funds is %s\t\t=\n", *inputRange)
	fmt.Println("=\t\t\t\t\t\t\t\t\t=")
	fmt.Println("=========================================================================")

	os.Setenv("GATEWAY_API", "https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-1Y.json")

	//Init Core Entity
	coreEntity.InitCoreEntity(coreThirdParty.NewEntity())
	//Init Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Execute
	_, err := job.GetFundRank(ctx, time.Now().Format("2006-01-02"), *inputRange)
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(0)
}
