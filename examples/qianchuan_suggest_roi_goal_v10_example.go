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

type ApiOpenApiV10QianchuanSuggestRoiGoalGetRequestExample struct {
	AdvertiserId       int64                                        `json:"advertiser_id"`
	AwemeId            int64                                        `json:"aweme_id,omitempty"`
	MarketingScene     QianchuanSuggestRoiGoalV10MarketingScene     `json:"marketing_scene,omitempty"`
	MarketingGoal      QianchuanSuggestRoiGoalV10MarketingGoal      `json:"marketing_goal,omitempty"`
	ProductId          int64                                        `json:"product_id,omitempty"`
	ProductNewOpen     bool                                         `json:"product_new_open,omitempty"`
	ExternalAction     QianchuanSuggestRoiGoalV10ExternalAction     `json:"external_action,omitempty"`
	CampaignScene      QianchuanSuggestRoiGoalV10CampaignScene      `json:"campaign_scene,omitempty"`
	DeepExternalAction QianchuanSuggestRoiGoalV10DeepExternalAction `json:"deep_external_action,omitempty"`
	DeepBidType        QianchuanSuggestRoiGoalV10DeepBidType        `json:"deep_bid_type,omitempty"`
	EcomGuestType      QianchuanSuggestRoiGoalV10EcomGuestType      `json:"ecom_guest_type,omitempty"`
	ShopId             int64                                        `json:"shop_id,omitempty"`
	BrandId            int64                                        `json:"brand_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/suggest/roi/goal/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanSuggestRoiGoalGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanSuggestRoiGoalV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AwemeId(request.AwemeId).MarketingScene(request.MarketingScene).MarketingGoal(request.MarketingGoal).ProductId(request.ProductId).ProductNewOpen(request.ProductNewOpen).ExternalAction(request.ExternalAction).CampaignScene(request.CampaignScene).DeepExternalAction(request.DeepExternalAction).DeepBidType(request.DeepBidType).EcomGuestType(request.EcomGuestType).ShopId(request.ShopId).BrandId(request.BrandId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
