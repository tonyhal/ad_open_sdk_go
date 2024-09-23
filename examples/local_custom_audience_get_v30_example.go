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

type ApiOpenApiV30LocalCustomAudienceGetGetRequestExample struct {
	LocalAccountId int64                             `json:"local_account_id"`
	TagsType       LocalCustomAudienceGetV30TagsType `json:"tags_type,omitempty"`
	Page           int64                             `json:"page,omitempty"`
	PageSize       int64                             `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/local/custom_audience/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30LocalCustomAudienceGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.LocalCustomAudienceGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		LocalAccountId(request.LocalAccountId).TagsType(request.TagsType).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
