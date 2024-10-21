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

type ApiOpenApiV30DouplusOptionalTargetsListGetRequestExample struct {
	AwemeSecUid string `json:"aweme_sec_uid,omitempty"`
	ItemId      int64  `json:"item_id,omitempty"`
	LiveAwemeId string `json:"live_aweme_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/douplus/optional_targets/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30DouplusOptionalTargetsListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.DouplusOptionalTargetsListV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AwemeSecUid(request.AwemeSecUid).ItemId(request.ItemId).LiveAwemeId(request.LiveAwemeId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}