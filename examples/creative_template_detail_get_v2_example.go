/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
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

type ApiOpenApi2CreativeTemplateDetailGetGetRequestExample struct {
	AdvertiserId int64                                   `json:"advertiser_id"`
	TemplateId   int64                                   `json:"template_id"`
	TemplateType CreativeTemplateDetailGetV2TemplateType `json:"template_type"`
}

// url: https://api.oceanengine.com/open_api/2/creative/template/detail/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2CreativeTemplateDetailGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.CreativeTemplateDetailGetV2Api().
		Get(ctx).
		AdvertiserId(request.AdvertiserId).TemplateId(request.TemplateId).TemplateType(request.TemplateType).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
