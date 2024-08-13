/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
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

type ApiOpenApiV10QianchuanReportCustomGetGetRequestExample struct {
	DataTopic    QianchuanReportCustomGetV10DataTopic       `json:"data_topic"`
	Dimensions   []string                                   `json:"dimensions"`
	Metrics      []string                                   `json:"metrics"`
	Filters      []*QianchuanReportCustomGetV10FiltersInner `json:"filters"`
	StartTime    string                                     `json:"start_time"`
	EndTime      string                                     `json:"end_time"`
	OrderBy      []*QianchuanReportCustomGetV10OrderByInner `json:"order_by"`
	AdvertiserId int64                                      `json:"advertiser_id,omitempty"`
	Page         int32                                      `json:"page,omitempty"`
	PageSize     int32                                      `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/report/custom/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanReportCustomGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanReportCustomGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		DataTopic(request.DataTopic).Dimensions(request.Dimensions).Metrics(request.Metrics).Filters(request.Filters).StartTime(request.StartTime).EndTime(request.EndTime).OrderBy(request.OrderBy).AdvertiserId(request.AdvertiserId).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
