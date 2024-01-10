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

type ApiOpenApi2StarChallengeGetPushAdResultsGetRequestExample struct {
	StarId          int64   `json:"star_id"`
	ChallengeTaskId int64   `json:"challenge_task_id"`
	ItemIds         []int64 `json:"item_ids"`
}

// url: https://api.oceanengine.com/open_api/2/star/challenge/get_push_ad_results/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarChallengeGetPushAdResultsGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarChallengeGetPushAdResultsV2Api().
		Get(ctx).
		AccessToken(accessToken).
		StarId(request.StarId).ChallengeTaskId(request.ChallengeTaskId).ItemIds(request.ItemIds).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
