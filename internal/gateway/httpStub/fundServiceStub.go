package httpStub

import (
	"context"
	"errors"
	"finnomena/internal/controller/job/model"
	"fmt"
	"sort"
	"time"
)

type Response struct {
	Status bool `json:"status"`
	Data   []model.Fund
}

type StubFundService struct {
	fund map[string]Response
}

func NewHTTPGateway() *StubFundService {
	stub := new(StubFundService)
	stub.fund = fund()
	return stub
}

func (repo *StubFundService) GetFundRanking(ctx context.Context, startDate time.Time, endDate time.Time) ([]model.Fund, error) {
	res := repo.fund[endDate.Format("2006-01-02")]
	if res.Status != true {
		fmt.Println("The HTTP response status is", res.Status)
		return nil, errors.New("The HTTP response status is false")
	}
	var fundList []model.Fund
	for _, i := range res.Data {
		if !i.UpdatedDate.Before(startDate) {
			fundList = append(fundList, model.Fund{
				Name:        i.Name,
				RankOfFund:  i.RankOfFund,
				UpdatedDate: i.UpdatedDate,
				Performance: i.Performance,
				Price:       i.Price,
			})
		}
	}

	if len(fundList) > 1 {
		sort.Slice(fundList[:], func(i, j int) bool {
			return fundList[i].Performance > fundList[j].Performance
		})
	}
	return fundList, nil
}

func fund() (stub map[string]Response) {
	stub = map[string]Response{}
	{
		rs := Response{}
		rs.Status = true
		fund := []model.Fund{}
		fund = append(fund, model.Fund{
			Name:        "SCBKEQTGE",
			RankOfFund:  22.149986,
			UpdatedDate: time.Date(2021, 9, 6, 0, 0, 0, 0, time.UTC),
			Performance: 39.434631,
			Price:       13.607800,
		})
		fund = append(fund, model.Fund{
			Name:        "MBT-G",
			RankOfFund:  19.821329,
			UpdatedDate: time.Date(2021, 9, 7, 0, 0, 0, 0, time.UTC),
			Performance: 39.210892,
			Price:       25.266600,
		})
		rs.Data = fund
		stub["2022-02-04"] = rs
	}

	{
		rs := Response{}
		rs.Status = true
		fund := []model.Fund{}
		fund = append(fund, model.Fund{
			Name:        "SCBKEQTGE",
			RankOfFund:  22.149986,
			UpdatedDate: time.Date(2021, 9, 6, 0, 0, 0, 0, time.UTC),
			Performance: 39.434631,
			Price:       13.607800,
		})
		fund = append(fund, model.Fund{
			Name:        "MBT-G",
			RankOfFund:  19.821329,
			UpdatedDate: time.Date(2021, 8, 7, 0, 0, 0, 0, time.UTC),
			Performance: 39.210892,
			Price:       25.266600,
		})
		fund = append(fund, model.Fund{
			Name:        "ASP-SME-SSF",
			RankOfFund:  38.081154,
			UpdatedDate: time.Date(2021, 8, 7, 0, 0, 0, 0, time.UTC),
			Performance: 38.703129,
			Price:       16.048300,
		})
		rs.Data = fund
		stub["2021-10-05"] = rs
	}

	{
		rs := Response{}
		rs.Status = true
		fund := []model.Fund{}
		fund = append(fund, model.Fund{
			Name:        "SCBKEQTGE",
			RankOfFund:  22.149986,
			UpdatedDate: time.Date(2021, 9, 6, 0, 0, 0, 0, time.UTC),
			Performance: 39.434631,
			Price:       13.607800,
		})
		fund = append(fund, model.Fund{
			Name:        "MBT-G",
			RankOfFund:  19.821329,
			UpdatedDate: time.Date(2021, 9, 7, 0, 0, 0, 0, time.UTC),
			Performance: 39.210892,
			Price:       25.266600,
		})
		fund = append(fund, model.Fund{
			Name:        "ASP-SME-SSF",
			RankOfFund:  38.081154,
			UpdatedDate: time.Date(2021, 9, 8, 0, 0, 0, 0, time.UTC),
			Performance: 38.703129,
			Price:       16.048300,
		})
		fund = append(fund, model.Fund{
			Name:        "KT-WTAI-A",
			RankOfFund:  28.829454,
			UpdatedDate: time.Date(2021, 9, 14, 0, 0, 0, 0, time.UTC),
			Performance: 39.315048,
			Price:       21.254601,
		})
		rs.Data = fund
		stub["2021-09-15"] = rs
	}

	{
		rs := Response{}
		rs.Status = false
		var fund []model.Fund
		rs.Data = fund
		stub["2022-03-15"] = rs
	}
	return
}
