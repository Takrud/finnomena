package job

import (
	"context"
	"finnomena/internal/controller/job/model"
	"finnomena/internal/core/coreEntity"
	"fmt"
	"time"
)

func GetFundRank(ctx context.Context, date string, timeRange string) ([]model.Fund, error) {
	fmt.Println("Execute GetFundRank ...")
	endDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	firstDate := FindStartDateFromRange(endDate, timeRange)
	listOfFunds, err := coreEntity.Entity().ThirdParty.HTTPFundService.GetFundRanking(ctx, firstDate, endDate)
	if err != nil {
		return nil, err
	}

	for _, i := range listOfFunds {
		fmt.Println("===============================================================")
		fmt.Printf("Name:\t\t%s\nRank Of Fund:\t%f\nUpdated Date:\t%s\nPerformance:\t%f\nPrice:\t\t%f\n", i.Name, i.RankOfFund, i.UpdatedDate.String(), i.Performance, i.Price)
	}
	fmt.Println("===============================================================")

	return listOfFunds, nil
}

func FindStartDateFromRange(endDate time.Time, timeRange string) time.Time {
	var startDate time.Time
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
