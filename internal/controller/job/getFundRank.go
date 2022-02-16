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
	firstDate := FindStartDateFromRange(timeRange, endDate)
	listOfFunds := coreEntity.Entity().ThirdParty.HTTPFundService.GetFundRanking(ctx, firstDate, endDate)
	for _, i := range listOfFunds {
		fmt.Printf("Name: %s\nRank Of Fund: %f\nUpdated Date: %s\nPerformance: %f\nPrice: %f\n", i.Name, i.RankOfFund, i.UpdatedDate.String(), i.Performance, i.Price)
	}

	return listOfFunds
}

func FindStartDateFromRange(timeRange string, startDate time.Time) time.Time {
	var endDate = time.Time{}
	switch timeRange {
	case "1Y":
		endDate = time.Date(startDate.Year()-1, startDate.Month(), startDate.Day()-1, 0, 0, 0, 0, time.UTC)
	case "1M":
		endDate = time.Date(startDate.Year(), startDate.Month()-1, startDate.Day()-1, 0, 0, 0, 0, time.UTC)
	case "1W":
		endDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day()-7, 0, 0, 0, 0, time.UTC)
	case "1D":
		endDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day()-2, 0, 0, 0, 0, time.UTC)
	}
	return endDate
}
