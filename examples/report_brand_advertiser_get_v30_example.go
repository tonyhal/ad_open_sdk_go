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

type ApiOpenApiV30ReportBrandAdvertiserGetGetRequestExample struct {
	AdvertiserId int64                                  `json:"advertiser_id"`
	StartTime    string                                 `json:"start_time"`
	EndTime      string                                 `json:"end_time"`
	Page         int64                                  `json:"page"`
	Size         int64                                  `json:"size"`
	LandingType  ReportBrandAdvertiserGetV30LandingType `json:"landing_type,omitempty"`
	PricingType  ReportBrandAdvertiserGetV30PricingType `json:"pricing_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/report/brand/advertiser/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ReportBrandAdvertiserGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportBrandAdvertiserGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).StartTime(request.StartTime).EndTime(request.EndTime).Page(request.Page).Size(request.Size).LandingType(request.LandingType).PricingType(request.PricingType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
