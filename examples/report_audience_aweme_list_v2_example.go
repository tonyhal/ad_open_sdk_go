/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
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

type ApiOpenApi2ReportAudienceAwemeListGetRequestExample struct {
	AdvertiserId int64                              `json:"advertiser_id,omitempty"`
	EndDate      *string                            `json:"end_date,omitempty"`
	Filtering    ReportAudienceAwemeListV2Filtering `json:"filtering,omitempty"`
	Metrics      []string                           `json:"metrics,omitempty"`
	Page         int64                              `json:"page,omitempty"`
	PageSize     int64                              `json:"page_size,omitempty"`
	StartDate    *string                            `json:"start_date,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/report/audience/aweme/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ReportAudienceAwemeListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportAudienceAwemeListV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).EndDate(request.EndDate).Filtering(request.Filtering).Metrics(request.Metrics).Page(request.Page).PageSize(request.PageSize).StartDate(request.StartDate).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
