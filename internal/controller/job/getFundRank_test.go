package job_test

import (
	"context"
	"finno/internal/controller/job"
	"finno/internal/controller/job/model"
	"finno/internal/core/coreEntity"
	"finno/internal/core/coreEntity/coreStub"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestGetFundRank(t *testing.T) {
	type args struct {
		inputDate string
		timeRange string
	}
	tests := []struct {
		name         string
		args         args
		expectedFund []model.Fund
	}{
		{
			name: "Success Case 1Y",
			args: args{
				inputDate: "2022-02-04",
				timeRange: "1Y",
			},
			expectedFund: expectedSuccess1Y(),
		},
		{
			name: "Success Case 1M",
			args: args{
				inputDate: "2021-10-05",
				timeRange: "1M",
			},
			expectedFund: expectedSuccess1M(),
		},
		{
			name: "Success Case 1W",
			args: args{
				inputDate: "2021-09-15",
				timeRange: "1W",
			},
			expectedFund: expectedSuccess1W(),
		},
		{
			name: "Success Case 1D",
			args: args{
				inputDate: "2021-09-15",
				timeRange: "1D",
			},
			expectedFund: expectedSuccess1D(),
		},
		{
			name: "Success Case 1D with No Fund",
			args: args{
				inputDate: "2021-10-05",
				timeRange: "1D",
			},
			expectedFund: expectedSuccess1DNoFund(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("case: ", tt.name)

			os.Setenv("GATEWAY_API", "https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-1Y.json")
			coreEntity.InitCoreEntity(coreStub.NewStubEntity())
			ctx := context.Background()
			res := job.GetFundRank(ctx, tt.args.inputDate, tt.args.timeRange)
			assert.Equal(t, len(tt.expectedFund), len(res))
			for o, i := range res {
				assert.Equal(t, tt.expectedFund[o].Name, i.Name)
				assert.Equal(t, tt.expectedFund[o].RankOfFund, i.RankOfFund)
				assert.Equal(t, tt.expectedFund[o].UpdatedDate, i.UpdatedDate)
				assert.Equal(t, tt.expectedFund[o].Performance, i.Performance)
				assert.Equal(t, tt.expectedFund[o].Price, i.Price)
			}
		})
	}
}

func expectedSuccess1Y() []model.Fund {
	fundList := []model.Fund{}
	fundList = append(fundList, model.Fund{
		Name:        "SCBKEQTGE",
		RankOfFund:  22.149986,
		UpdatedDate: time.Date(2021, 9, 6, 0, 0, 0, 0, time.UTC),
		Performance: 39.434631,
		Price:       13.607800,
	})
	fundList = append(fundList, model.Fund{
		Name:        "MBT-G",
		RankOfFund:  19.821329,
		UpdatedDate: time.Date(2021, 9, 7, 0, 0, 0, 0, time.UTC),
		Performance: 39.210892,
		Price:       25.266600,
	})
	return fundList
}

func expectedSuccess1M() []model.Fund {
	fundList := []model.Fund{}
	fundList = append(fundList, model.Fund{
		Name:        "SCBKEQTGE",
		RankOfFund:  22.149986,
		UpdatedDate: time.Date(2021, 9, 6, 0, 0, 0, 0, time.UTC),
		Performance: 39.434631,
		Price:       13.607800,
	})
	return fundList
}

func expectedSuccess1W() []model.Fund {
	fundList := []model.Fund{}
	fundList = append(fundList, model.Fund{
		Name:        "KT-WTAI-A",
		RankOfFund:  28.829454,
		UpdatedDate: time.Date(2021, 9, 14, 0, 0, 0, 0, time.UTC),
		Performance: 39.315048,
		Price:       21.254601,
	})
	fundList = append(fundList, model.Fund{
		Name:        "ASP-SME-SSF",
		RankOfFund:  38.081154,
		UpdatedDate: time.Date(2021, 9, 8, 0, 0, 0, 0, time.UTC),
		Performance: 38.703129,
		Price:       16.048300,
	})
	return fundList
}

func expectedSuccess1D() []model.Fund {
	fundList := []model.Fund{}
	fundList = append(fundList, model.Fund{
		Name:        "KT-WTAI-A",
		RankOfFund:  28.829454,
		UpdatedDate: time.Date(2021, 9, 14, 0, 0, 0, 0, time.UTC),
		Performance: 39.315048,
		Price:       21.254601,
	})
	return fundList
}

func expectedSuccess1DNoFund() []model.Fund {
	fundList := []model.Fund{}
	return fundList
}

func TestFindStartDateFromRange(t *testing.T) {
	type args struct {
		inputDate time.Time
		timeRange string
	}
	tests := []struct {
		name         string
		args         args
		expectedFund time.Time
	}{
		{
			name: "Success Case 1Y",
			args: args{
				inputDate: time.Date(2022, 2, 4, 0, 0, 0, 0, time.UTC),
				timeRange: "1Y",
			},
			expectedFund: expected1Y(),
		},
		{
			name: "Success Case 1M",
			args: args{
				inputDate: time.Date(2021, 10, 5, 0, 0, 0, 0, time.UTC),
				timeRange: "1M",
			},
			expectedFund: expected1M(),
		},
		{
			name: "Success Case 1W",
			args: args{
				inputDate: time.Date(2021, 9, 15, 0, 0, 0, 0, time.UTC),
				timeRange: "1W",
			},
			expectedFund: expected1W(),
		},
		{
			name: "Success Case 1D",
			args: args{
				inputDate: time.Date(2021, 9, 15, 0, 0, 0, 0, time.UTC),
				timeRange: "1D",
			},
			expectedFund: expected1D(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("case: ", tt.name)
			res := job.FindStartDateFromRange(tt.args.inputDate, tt.args.timeRange)
			assert.Equal(t, tt.expectedFund.Format("2006-01-02"), res.Format("2006-01-02"))
		})
	}
}

func expected1Y() time.Time {
	return time.Date(2021, 2, 3, 0, 0, 0, 0, time.UTC)
}

func expected1M() time.Time {
	return time.Date(2021, 9, 4, 0, 0, 0, 0, time.UTC)
}

func expected1W() time.Time {
	return time.Date(2021, 9, 8, 0, 0, 0, 0, time.UTC)
}

func expected1D() time.Time {
	return time.Date(2021, 9, 14, 0, 0, 0, 0, time.UTC)
}
