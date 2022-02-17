package http

import (
	"context"
	"encoding/json"
	"errors"
	"finnomena/internal/controller/job/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"time"
)

type Response struct {
	Status bool `json:"status"`
	Data   []model.Fund
}

type HTTPGateway struct {
}

//NewHTTPGateway - New Http SavDomain Repository
func NewHTTPGateway() *HTTPGateway {
	g := new(HTTPGateway)
	return g
}

func (g *HTTPGateway) GetFundRanking(ctx context.Context, startDate time.Time, endDate time.Time) ([]model.Fund, error) {

	url, ok := os.LookupEnv("GATEWAY_API")
	if !ok {
		panic("GATEWAY_API not set")
	}

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("The HTTP request failed with error ", err)
		return nil, err
	}

	fmt.Println(response)

	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))

	rs := Response{}
	err = json.Unmarshal(data, &rs)
	if err != nil {
		fmt.Println("The Json Unmarshall failed with error ", err)
		panic(err)
	}

	if rs.Status != true {
		fmt.Println("The HTTP response status is ", rs.Status)
		return nil, errors.New("The HTTP response status is false.")
	}

	var fundList []model.Fund
	for _, i := range rs.Data {
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
			return (fundList[i].Performance > fundList[j].Performance)
		})
	}
	return fundList, nil
}
