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

type ApiOpenApiV30SubscribeAccountsAddPostRequestExample struct {
	APPAccessToken                 string                         `json:"APP-Access-Token"`
	SubscribeAccountsAddV30Request SubscribeAccountsAddV30Request `json:"SubscribeAccountsAddV30Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/subscribe/accounts/add/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30SubscribeAccountsAddPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.SubscribeAccountsAddV30Api().
		Post(ctx).
		AccessToken(accessToken).
		APPAccessToken(request.APPAccessToken).SubscribeAccountsAddV30Request(request.SubscribeAccountsAddV30Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
