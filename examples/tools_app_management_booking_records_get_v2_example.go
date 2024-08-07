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

type ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequestExample struct {
	AdvertiserId int64                                           `json:"advertiser_id"`
	PackageId    string                                          `json:"package_id"`
	HostType     ToolsAppManagementBookingRecordsGetV2HostType   `json:"host_type"`
	Page         int64                                           `json:"page,omitempty"`
	PageSize     int64                                           `json:"page_size,omitempty"`
	CreateTime   ToolsAppManagementBookingRecordsGetV2CreateTime `json:"create_time,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/app_management/booking_records/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAppManagementBookingRecordsGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).PackageId(request.PackageId).HostType(request.HostType).Page(request.Page).PageSize(request.PageSize).CreateTime(request.CreateTime).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
