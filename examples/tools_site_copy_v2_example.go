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

type ApiOpenApi2ToolsSiteCopyPostRequestExample struct {
	XOrangeCaller          string                 `json:"X-Orange-Caller,omitempty"`
	ToolsSiteCopyV2Request ToolsSiteCopyV2Request `json:"ToolsSiteCopyV2Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/site/copy/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsSiteCopyPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsSiteCopyV2Api().
		Post(ctx).
		AccessToken(accessToken).
		XOrangeCaller(request.XOrangeCaller).ToolsSiteCopyV2Request(request.ToolsSiteCopyV2Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
