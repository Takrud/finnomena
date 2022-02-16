package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"finno/internal/controller/job/model"
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

func (g *HTTPGateway) GetFundRanking(ctx context.Context, startDate time.Time, endDate time.Time) []model.Fund {

	url, ok := os.LookupEnv("GATEWAY_API")
	if !ok {
		panic("GATEWAY_API not set")
	} 
	
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("The HTTP request failed with error ", err)
		panic(err)
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
		panic(err)
	}

	fundList := []model.Fund{}
	for _, i := range rs.Data {
		if i.UpdatedDate.After(startDate) {
			fundList = append(fundList, model.Fund{
				Name: i.Name,
				RankOfFund: i.RankOfFund,
				UpdatedDate: i.UpdatedDate,
				Performance: i.Performance,
				Price: i.Price,
			})
		}
	}

	return fundList
}