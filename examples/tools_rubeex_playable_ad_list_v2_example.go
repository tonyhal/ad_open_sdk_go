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

type ApiOpenApi2ToolsRubeexPlayableAdListGetRequestExample struct {
	PlayableId   int32 `json:"playable_id"`
	AdvertiserId int64 `json:"advertiser_id"`
	Page         int32 `json:"page,omitempty"`
	PageSize     int32 `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/rubeex_playable/ad_list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsRubeexPlayableAdListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsRubeexPlayableAdListV2Api().
		Get(ctx).
		AccessToken(accessToken).
		PlayableId(request.PlayableId).AdvertiserId(request.AdvertiserId).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
