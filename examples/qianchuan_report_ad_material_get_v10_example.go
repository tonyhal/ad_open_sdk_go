/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


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

type ApiOpenApiV10QianchuanReportAdMaterialGetGetRequestExample struct {
	AdvertiserId int64                                    `json:"advertiser_id"`
	AdId         int64                                    `json:"ad_id"`
	StartDate    string                                   `json:"start_date"`
	EndDate      string                                   `json:"end_date"`
	Fields       []string                                 `json:"fields"`
	Filtering    QianchuanReportAdMaterialGetV10Filtering `json:"filtering"`
	OrderType    QianchuanReportAdMaterialGetV10OrderType `json:"order_type,omitempty"`
	OrderField   string                                   `json:"order_field,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/report/ad/material/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanReportAdMaterialGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanReportAdMaterialGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AdId(request.AdId).StartDate(request.StartDate).EndDate(request.EndDate).Fields(request.Fields).Filtering(request.Filtering).OrderType(request.OrderType).OrderField(request.OrderField).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
