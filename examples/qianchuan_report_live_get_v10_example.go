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

type ApiOpenApiV10QianchuanReportLiveGetGetRequestExample struct {
	AdvertiserId   int64                                   `json:"advertiser_id"`
	AwemeId        int64                                   `json:"aweme_id"`
	StartTime      string                                  `json:"start_time"`
	EndTime        string                                  `json:"end_time"`
	Fields         []string                                `json:"fields"`
	StatsAuthority QianchuanReportLiveGetV10StatsAuthority `json:"stats_authority,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/report/live/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanReportLiveGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanReportLiveGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AwemeId(request.AwemeId).StartTime(request.StartTime).EndTime(request.EndTime).Fields(request.Fields).StatsAuthority(request.StatsAuthority).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
