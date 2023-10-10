/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
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

type ApiOpenApi2ToolsAppManagementAndroidBasicPackageUpdatePostRequestExample struct {
	ToolsAppManagementAndroidBasicPackageUpdateV2Request ToolsAppManagementAndroidBasicPackageUpdateV2Request `json:"ToolsAppManagementAndroidBasicPackageUpdateV2Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/app_management/android_basic_package/update/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAppManagementAndroidBasicPackageUpdatePostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAppManagementAndroidBasicPackageUpdateV2Api().
		Post(ctx).
		AccessToken(accessToken).
		ToolsAppManagementAndroidBasicPackageUpdateV2Request(request.ToolsAppManagementAndroidBasicPackageUpdateV2Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
