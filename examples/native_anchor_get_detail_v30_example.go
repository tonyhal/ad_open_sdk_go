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

type ApiOpenApiV30NativeAnchorGetDetailGetRequestExample struct {
	AnchorIds    []string                           `json:"anchor_ids"`
	AdvertiserId int64                              `json:"advertiser_id"`
	AnchorType   NativeAnchorGetDetailV30AnchorType `json:"anchor_type"`
}

// url: https://api.oceanengine.com/open_api/v3.0/native_anchor/get/detail/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30NativeAnchorGetDetailGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.NativeAnchorGetDetailV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AnchorIds(request.AnchorIds).AdvertiserId(request.AdvertiserId).AnchorType(request.AnchorType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
