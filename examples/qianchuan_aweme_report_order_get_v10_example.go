/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

type ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequestExample struct {
	AdvertiserId int64                                    `json:"advertiser_id"`
	StartDate    string                                   `json:"start_date"`
	EndDate      string                                   `json:"end_date"`
	Filtering    QianchuanAwemeReportOrderGetV10Filtering `json:"filtering"`
	Fields       []string                                 `json:"fields,omitempty"`
	OrderField   string                                   `json:"order_field,omitempty"`
	OrderType    QianchuanAwemeReportOrderGetV10OrderType `json:"order_type,omitempty"`
	Page         int64                                    `json:"page,omitempty"`
	PageSize     QianchuanAwemeReportOrderGetV10PageSize  `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/aweme/report/order/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanAwemeReportOrderGetV10Api().
		Get(ctx).
		AdvertiserId(request.AdvertiserId).StartDate(request.StartDate).EndDate(request.EndDate).Filtering(request.Filtering).Fields(request.Fields).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
