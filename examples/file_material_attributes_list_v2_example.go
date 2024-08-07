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

type ApiOpenApi2FileMaterialAttributesListGetRequestExample struct {
	AccountId                   int64                                   `json:"account_id"`
	AccountType                 FileMaterialAttributesListV2AccountType `json:"account_type"`
	PageSize                    int32                                   `json:"page_size"`
	Page                        int32                                   `json:"page"`
	Filtering                   FileMaterialAttributesListV2Filtering   `json:"filtering,omitempty"`
	ReturnLowqualitySuggestions bool                                    `json:"return_lowquality_suggestions,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/file/material_attributes/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2FileMaterialAttributesListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.FileMaterialAttributesListV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AccountId(request.AccountId).AccountType(request.AccountType).PageSize(request.PageSize).Page(request.Page).Filtering(request.Filtering).ReturnLowqualitySuggestions(request.ReturnLowqualitySuggestions).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
