/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
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

type ApiOpenApi2StarDemandListGetRequestExample struct {
	Filtering StarDemandListV2Filtering `json:"filtering,omitempty"`
	Page      int64                     `json:"page,omitempty"`
	PageSize  int64                     `json:"page_size,omitempty"`
	StarId    int64                     `json:"star_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/demand/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarDemandListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarDemandListV2Api().
		Get(ctx).
		AccessToken(accessToken).
		Filtering(request.Filtering).Page(request.Page).PageSize(request.PageSize).StarId(request.StarId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
