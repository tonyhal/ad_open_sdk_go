/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
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

type ApiOpenApi2ToolsPlayableListGetGetRequestExample struct {
	AdvertiserId   int64                                `json:"advertiser_id,omitempty"`
	Page           int64                                `json:"page,omitempty"`
	PageSize       int64                                `json:"page_size,omitempty"`
	PlayableSource ToolsPlayableListGetV2PlayableSource `json:"playable_source,omitempty"`
	PlayableUrl    string                               `json:"playable_url,omitempty"`
	Status         ToolsPlayableListGetV2Status         `json:"status,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/playable_list/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsPlayableListGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsPlayableListGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Page(request.Page).PageSize(request.PageSize).PlayableSource(request.PlayableSource).PlayableUrl(request.PlayableUrl).Status(request.Status).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
