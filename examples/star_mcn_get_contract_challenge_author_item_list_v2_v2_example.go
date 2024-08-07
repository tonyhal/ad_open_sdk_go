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

type ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequestExample struct {
	StarId      int64 `json:"star_id"`
	DemandId    int64 `json:"demand_id"`
	Page        int32 `json:"page"`
	PageSize    int32 `json:"page_size"`
	DeveloperId int64 `json:"developer_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/mcn/get_contract_challenge_author_item_list_v2/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarMcnGetContractChallengeAuthorItemListV2GetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarMcnGetContractChallengeAuthorItemListV2V2Api().
		Get(ctx).
		AccessToken(accessToken).
		StarId(request.StarId).DemandId(request.DemandId).Page(request.Page).PageSize(request.PageSize).DeveloperId(request.DeveloperId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
