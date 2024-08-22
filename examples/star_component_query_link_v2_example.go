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

type ApiOpenApi2StarComponentQueryLinkGetRequestExample struct {
	StarId          int64   `json:"star_id"`
	ComponentStatus int32   `json:"component_status,omitempty"`
	Page            int32   `json:"page,omitempty"`
	Limit           int32   `json:"limit,omitempty"`
	LinkIds         []int64 `json:"link_ids,omitempty"`
	LinkType        int32   `json:"link_type,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/component/query_link/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarComponentQueryLinkGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarComponentQueryLinkV2Api().
		Get(ctx).
		AccessToken(accessToken).
		StarId(request.StarId).ComponentStatus(request.ComponentStatus).Page(request.Page).Limit(request.Limit).LinkIds(request.LinkIds).LinkType(request.LinkType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
