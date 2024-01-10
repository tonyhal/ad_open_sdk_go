/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
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

type ApiOpenApi2ClueFormListGetRequestExample struct {
	AdvertiserId int64               `json:"advertiser_id,omitempty"`
	EndTime      *string             `json:"end_time,omitempty"`
	InstanceIds  []int64             `json:"instance_ids,omitempty"`
	IsDel        ClueFormListV2IsDel `json:"is_del,omitempty"`
	Page         int64               `json:"page,omitempty"`
	PageSize     int64               `json:"page_size,omitempty"`
	StartTime    *string             `json:"start_time,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/clue/form/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ClueFormListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ClueFormListV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).EndTime(request.EndTime).InstanceIds(request.InstanceIds).IsDel(request.IsDel).Page(request.Page).PageSize(request.PageSize).StartTime(request.StartTime).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
