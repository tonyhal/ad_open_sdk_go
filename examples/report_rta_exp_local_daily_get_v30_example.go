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

type ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequestExample struct {
	RtaId        int64                                 `json:"rta_id"`
	AdvertiserId int64                                 `json:"advertiser_id"`
	StartDate    string                                `json:"start_date"`
	EndDate      string                                `json:"end_date"`
	Vid          int64                                 `json:"vid,omitempty"`
	CusVid       int64                                 `json:"cus_vid,omitempty"`
	Filtering    ReportRtaExpLocalDailyGetV30Filtering `json:"filtering,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/report/rta_exp_local_daily/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ReportRtaExpLocalDailyGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		RtaId(request.RtaId).AdvertiserId(request.AdvertiserId).StartDate(request.StartDate).EndDate(request.EndDate).Vid(request.Vid).CusVid(request.CusVid).Filtering(request.Filtering).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
