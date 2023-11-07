/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
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

type ApiOpenApiV10QianchuanToolsAllowCouponGetRequestExample struct {
	AdvertiserId   int64                                      `json:"advertiser_id"`
	MarketingGoal  QianchuanToolsAllowCouponV10MarketingGoal  `json:"marketing_goal"`
	CampaignScene  QianchuanToolsAllowCouponV10CampaignScene  `json:"campaign_scene"`
	MarketingScene QianchuanToolsAllowCouponV10MarketingScene `json:"marketing_scene"`
	AwemeIds       []int64                                    `json:"aweme_ids,omitempty"`
	ProductIds     []int64                                    `json:"product_ids,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/tools/allow_coupon/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanToolsAllowCouponGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanToolsAllowCouponV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).MarketingGoal(request.MarketingGoal).CampaignScene(request.CampaignScene).MarketingScene(request.MarketingScene).AwemeIds(request.AwemeIds).ProductIds(request.ProductIds).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
