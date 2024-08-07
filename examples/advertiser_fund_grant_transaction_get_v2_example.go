/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
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

type ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequestExample struct {
	AdvertiserId    int64  `json:"advertiser_id"`
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	Page            string `json:"page,omitempty"`
	PageSize        string `json:"page_size,omitempty"`
	TransactionType string `json:"transaction_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/advertiser/fund/grant_transaction/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.AdvertiserFundGrantTransactionGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).StartTime(request.StartTime).EndTime(request.EndTime).Page(request.Page).PageSize(request.PageSize).TransactionType(request.TransactionType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
