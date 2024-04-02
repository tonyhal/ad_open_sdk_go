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

type ApiOpenApi2KeywordFeedadsSuggestGetRequestExample struct {
	AdId           int64 `json:"ad_id,omitempty"`
	AdvertiserId   int64 `json:"advertiser_id,omitempty"`
	ReqKeywordType int64 `json:"req_keyword_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/keyword_feedads/suggest/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2KeywordFeedadsSuggestGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.KeywordFeedadsSuggestV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdId(request.AdId).AdvertiserId(request.AdvertiserId).ReqKeywordType(request.ReqKeywordType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
