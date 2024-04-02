/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
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

type ApiOpenApi2ToolsAdConvertDeepbidReadGetRequestExample struct {
	AdvertiserId       int64                                         `json:"advertiser_id,omitempty"`
	AssetsIds          []int64                                       `json:"assets_ids,omitempty"`
	CampaignId         int64                                         `json:"campaign_id,omitempty"`
	ConvertId          int64                                         `json:"convert_id,omitempty"`
	DeepExternalAction ToolsAdConvertDeepbidReadV2DeepExternalAction `json:"deep_external_action,omitempty"`
	DeliveryRange      ToolsAdConvertDeepbidReadV2DeliveryRange      `json:"delivery_range,omitempty"`
	ExternalAction     ToolsAdConvertDeepbidReadV2ExternalAction     `json:"external_action,omitempty"`
	FlowControlMode    ToolsAdConvertDeepbidReadV2FlowControlMode    `json:"flow_control_mode,omitempty"`
	SmartBidType       ToolsAdConvertDeepbidReadV2SmartBidType       `json:"smart_bid_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/ad_convert/deepbid/read/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAdConvertDeepbidReadGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAdConvertDeepbidReadV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AssetsIds(request.AssetsIds).CampaignId(request.CampaignId).ConvertId(request.ConvertId).DeepExternalAction(request.DeepExternalAction).DeliveryRange(request.DeliveryRange).ExternalAction(request.ExternalAction).FlowControlMode(request.FlowControlMode).SmartBidType(request.SmartBidType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
