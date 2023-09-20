/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
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

type ApiOpenApi2ToolsRubeexGetGetRequestExample struct {
	AdvertiserId  int64                            `json:"advertiser_id,omitempty"`
	Filtering     ToolsRubeexGetV2Filtering        `json:"filtering,omitempty"`
	Page          int64                            `json:"page,omitempty"`
	PageSize      int64                            `json:"page_size,omitempty"`
	PlatformName  []*ToolsRubeexGetV2PlatformName  `json:"platform_name,omitempty"`
	ProjectStatus []*ToolsRubeexGetV2ProjectStatus `json:"project_status,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/rubeex/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsRubeexGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsRubeexGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Filtering(request.Filtering).Page(request.Page).PageSize(request.PageSize).PlatformName(request.PlatformName).ProjectStatus(request.ProjectStatus).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
