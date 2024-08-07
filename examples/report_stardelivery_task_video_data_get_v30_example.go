/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
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

type ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequestExample struct {
	AdvertiserId int64                                          `json:"advertiser_id"`
	StarTaskId   int64                                          `json:"star_task_id"`
	StartDate    string                                         `json:"start_date"`
	EndDate      string                                         `json:"end_date"`
	Filtering    ReportStardeliveryTaskVideoDataGetV30Filtering `json:"filtering,omitempty"`
	OrderField   string                                         `json:"order_field,omitempty"`
	OrderType    ReportStardeliveryTaskVideoDataGetV30OrderType `json:"order_type,omitempty"`
	Page         int32                                          `json:"page,omitempty"`
	PageSize     int32                                          `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/report/stardelivery/task_video_data/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportStardeliveryTaskVideoDataGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).StarTaskId(request.StarTaskId).StartDate(request.StartDate).EndDate(request.EndDate).Filtering(request.Filtering).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
