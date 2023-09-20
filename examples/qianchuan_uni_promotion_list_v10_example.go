/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
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

type ApiOpenApiV10QianchuanUniPromotionListGetRequestExample struct {
	AdvertiserId  int64                                     `json:"advertiser_id"`
	StartTime     string                                    `json:"start_time"`
	EndTime       string                                    `json:"end_time"`
	MarketingGoal QianchuanUniPromotionListV10MarketingGoal `json:"marketing_goal"`
	Fields        []*QianchuanUniPromotionListV10Fields     `json:"fields"`
	OrderType     QianchuanUniPromotionListV10OrderType     `json:"order_type,omitempty"`
	OrderField    QianchuanUniPromotionListV10OrderField    `json:"order_field,omitempty"`
	Page          int32                                     `json:"page,omitempty"`
	PageSize      QianchuanUniPromotionListV10PageSize      `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/uni_promotion/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanUniPromotionListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanUniPromotionListV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).StartTime(request.StartTime).EndTime(request.EndTime).MarketingGoal(request.MarketingGoal).Fields(request.Fields).OrderType(request.OrderType).OrderField(request.OrderField).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}