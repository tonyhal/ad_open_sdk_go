/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
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

type ApiOpenApiV30ReportBrandCampaignGetGetRequestExample struct {
	AdvertiserId int64                                `json:"advertiser_id"`
	Page         int64                                `json:"page"`
	Size         int64                                `json:"size"`
	LandingType  ReportBrandCampaignGetV30LandingType `json:"landing_type,omitempty"`
	PricingType  ReportBrandCampaignGetV30PricingType `json:"pricing_type,omitempty"`
	CampaignIds  []string                             `json:"campaign_ids,omitempty"`
	StartTime    string                               `json:"start_time,omitempty"`
	EndTime      string                               `json:"end_time,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/report/brand/campaign/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ReportBrandCampaignGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportBrandCampaignGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Page(request.Page).Size(request.Size).LandingType(request.LandingType).PricingType(request.PricingType).CampaignIds(request.CampaignIds).StartTime(request.StartTime).EndTime(request.EndTime).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
