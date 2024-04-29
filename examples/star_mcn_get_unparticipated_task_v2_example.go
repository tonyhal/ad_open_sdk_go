/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
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

type ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequestExample struct {
	StarId             int64 `json:"star_id"`
	Page               int32 `json:"page,omitempty"`
	PageSize           int32 `json:"page_size,omitempty"`
	PayType            int32 `json:"pay_type,omitempty"`
	MinCreateTimeStamp int64 `json:"min_create_time_stamp,omitempty"`
	MaxCreateTimeStamp int64 `json:"max_create_time_stamp,omitempty"`
	DeveloperId        int64 `json:"developer_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/mcn/get_unparticipated_task/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarMcnGetUnparticipatedTaskV2Api().
		Get(ctx).
		AccessToken(accessToken).
		StarId(request.StarId).Page(request.Page).PageSize(request.PageSize).PayType(request.PayType).MinCreateTimeStamp(request.MinCreateTimeStamp).MaxCreateTimeStamp(request.MaxCreateTimeStamp).DeveloperId(request.DeveloperId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}