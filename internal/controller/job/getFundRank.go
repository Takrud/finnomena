package job

import (
	"context"
	"finno/internal/controller/job/model"
	"finno/internal/core/coreEntity"
	"fmt"
	"time"
)

func GetFundRank(ctx context.Context, date string, timeRange string) []model.Fund {
	fmt.Println(date)
	endDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	firstDate := FindStartDateFromRange(endDate, timeRange)
	listOfFunds := coreEntity.Entity().ThirdParty.HTTPFundService.GetFundRanking(ctx, firstDate, endDate)
	for _, i := range listOfFunds {
		fmt.Println("===============================================================")
		fmt.Printf("Name: %s\nRank Of Fund: %f\nUpdated Date: %s\nPerformance: %f\nPrice: %f\n", i.Name, i.RankOfFund, i.UpdatedDate.String(), i.Performance, i.Price)
	}
	fmt.Println("===============================================================")
	return listOfFunds
}

func FindStartDateFromRange(endDate time.Time, timeRange string) time.Time {
	var startDate = time.Time{}
	switch timeRange {
	case "1Y":
		startDate = time.Date(endDate.Year()-1, endDate.Month(), endDate.Day()-1, 0, 0, 0, 0, time.UTC)
	case "1M":
		startDate = time.Date(endDate.Year(), endDate.Month()-1, endDate.Day()-1, 0, 0, 0, 0, time.UTC)
	case "1W":
		startDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day()-7, 0, 0, 0, 0, time.UTC)
	case "1D":
		startDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day()-1, 0, 0, 0, 0, time.UTC)
	}
	return startDate
}
