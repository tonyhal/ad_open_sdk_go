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

type ApiOpenApiV30DouplusOptionalItemsListGetRequestExample struct {
	AwemeSecUid    string                                    `json:"aweme_sec_uid"`
	ExternalAction DouplusOptionalItemsListV30ExternalAction `json:"external_action"`
	AwemeId        string                                    `json:"aweme_id"`
	Count          int64                                     `json:"count"`
	Cursor         int64                                     `json:"cursor"`
}

// url: https://api.oceanengine.com/open_api/v3.0/douplus/optional_items/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30DouplusOptionalItemsListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.DouplusOptionalItemsListV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AwemeSecUid(request.AwemeSecUid).ExternalAction(request.ExternalAction).AwemeId(request.AwemeId).Count(request.Count).Cursor(request.Cursor).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}