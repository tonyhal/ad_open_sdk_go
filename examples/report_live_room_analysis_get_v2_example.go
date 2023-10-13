/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
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

type ApiOpenApi2ReportLiveRoomAnalysisGetGetRequestExample struct {
	AdvertiserId int64                                 `json:"advertiser_id,omitempty"`
	EndTime      *string                               `json:"end_time,omitempty"`
	Fields       []string                              `json:"fields,omitempty"`
	Filtering    ReportLiveRoomAnalysisGetV2Filtering  `json:"filtering,omitempty"`
	OrderField   ReportLiveRoomAnalysisGetV2OrderField `json:"order_field,omitempty"`
	OrderType    ReportLiveRoomAnalysisGetV2OrderType  `json:"order_type,omitempty"`
	Page         int64                                 `json:"page,omitempty"`
	PageSize     int64                                 `json:"page_size,omitempty"`
	StartTime    *string                               `json:"start_time,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/report/live_room/analysis/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ReportLiveRoomAnalysisGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportLiveRoomAnalysisGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).EndTime(request.EndTime).Fields(request.Fields).Filtering(request.Filtering).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).StartTime(request.StartTime).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
