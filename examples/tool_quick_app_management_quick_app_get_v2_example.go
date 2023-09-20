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

type ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequestExample struct {
	AdvertiserId int64                                         `json:"advertiser_id"`
	Status       []*ToolQuickAppManagementQuickAppGetV2Status  `json:"status,omitempty"`
	Page         int32                                         `json:"page,omitempty"`
	PageSize     int32                                         `json:"page_size,omitempty"`
	UpdateTime   ToolQuickAppManagementQuickAppGetV2UpdateTime `json:"update_time,omitempty"`
	SearchKey    string                                        `json:"search_key,omitempty"`
	QuickAppIds  []int64                                       `json:"quick_app_ids,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tool/quick_app_management/quick_app/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolQuickAppManagementQuickAppGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Status(request.Status).Page(request.Page).PageSize(request.PageSize).UpdateTime(request.UpdateTime).SearchKey(request.SearchKey).QuickAppIds(request.QuickAppIds).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
