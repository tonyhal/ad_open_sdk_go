/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
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

type ApiOpenApiV10QianchuanReportUniPromotionGetGetRequestExample struct {
	AdvertiserId  int64                                          `json:"advertiser_id"`
	StartDate     string                                         `json:"start_date"`
	EndDate       string                                         `json:"end_date"`
	MarketingGoal QianchuanReportUniPromotionGetV10MarketingGoal `json:"marketing_goal"`
	LabAdType     QianchuanReportUniPromotionGetV10LabAdType     `json:"lab_ad_type"`
	Fields        []string                                       `json:"fields"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/report/uni_promotion/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanReportUniPromotionGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanReportUniPromotionGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).StartDate(request.StartDate).EndDate(request.EndDate).MarketingGoal(request.MarketingGoal).LabAdType(request.LabAdType).Fields(request.Fields).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
