/*
API Title

巨量引擎开放平台

API version: 1.0.0
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

type ApiOpenApiV30ToolsKeywordsBidRatioDeletePostRequestExample struct {
	ToolsKeywordsBidRatioDeleteV30Request ToolsKeywordsBidRatioDeleteV30Request `json:"ToolsKeywordsBidRatioDeleteV30Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/tools/keywords_bid_ratio/delete/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ToolsKeywordsBidRatioDeletePostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsKeywordsBidRatioDeleteV30Api().
		Post(ctx).
		AccessToken(accessToken).
		ToolsKeywordsBidRatioDeleteV30Request(request.ToolsKeywordsBidRatioDeleteV30Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
