/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


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

type ApiOpenApiV30BrandQueryStockGetRequestExample struct {
	AdvertiserId        int64                          `json:"advertiser_id"`
	Classify            BrandQueryStockV30Classify     `json:"classify"`
	LandingType         BrandQueryStockV30LandingType  `json:"landing_type"`
	AppOrigin           BrandQueryStockV30AppOrigin    `json:"app_origin"`
	AdForm              BrandQueryStockV30AdForm       `json:"ad_form"`
	GdSendType          BrandQueryStockV30GdSendType   `json:"gd_send_type"`
	PolicyNo            string                         `json:"policy_no"`
	ScheduleInfo        BrandQueryStockV30ScheduleInfo `json:"schedule_info"`
	AudienceInfo        BrandQueryStockV30AudienceInfo `json:"audience_info,omitempty"`
	MerchantIntentionNo string                         `json:"merchant_intention_no,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/brand/query_stock/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30BrandQueryStockGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.BrandQueryStockV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Classify(request.Classify).LandingType(request.LandingType).AppOrigin(request.AppOrigin).AdForm(request.AdForm).GdSendType(request.GdSendType).PolicyNo(request.PolicyNo).ScheduleInfo(request.ScheduleInfo).AudienceInfo(request.AudienceInfo).MerchantIntentionNo(request.MerchantIntentionNo).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
