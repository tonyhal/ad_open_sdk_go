/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
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

type ApiOpenApi2StarVasAppendOrderToBoostItemGroupPostRequestExample struct {
	StarVasAppendOrderToBoostItemGroupV2Request StarVasAppendOrderToBoostItemGroupV2Request `json:"StarVasAppendOrderToBoostItemGroupV2Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/vas/append_order_to_boost_item_group/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarVasAppendOrderToBoostItemGroupPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarVasAppendOrderToBoostItemGroupV2Api().
		Post(ctx).
		AccessToken(accessToken).
		StarVasAppendOrderToBoostItemGroupV2Request(request.StarVasAppendOrderToBoostItemGroupV2Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
