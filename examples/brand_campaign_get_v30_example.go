/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
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

type ApiOpenApiV30BrandCampaignGetGetRequestExample struct {
	AdvertiserId    int64                             `json:"advertiser_id"`
	Page            int64                             `json:"page"`
	Size            int64                             `json:"size"`
	CampaignIds     []string                          `json:"campaign_ids,omitempty"`
	CampaignNames   []string                          `json:"campaign_names,omitempty"`
	CampaignStatus  BrandCampaignGetV30CampaignStatus `json:"campaign_status,omitempty"`
	CreateStartTime string                            `json:"create_start_time,omitempty"`
	CreateEndTime   string                            `json:"create_end_time,omitempty"`
	StartTime       string                            `json:"start_time,omitempty"`
	EndTime         string                            `json:"end_time,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/brand/campaign/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30BrandCampaignGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.BrandCampaignGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Page(request.Page).Size(request.Size).CampaignIds(request.CampaignIds).CampaignNames(request.CampaignNames).CampaignStatus(request.CampaignStatus).CreateStartTime(request.CreateStartTime).CreateEndTime(request.CreateEndTime).StartTime(request.StartTime).EndTime(request.EndTime).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}