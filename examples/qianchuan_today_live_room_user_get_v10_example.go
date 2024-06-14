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

type ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequestExample struct {
	AdvertiserId int64                                        `json:"advertiser_id"`
	RoomId       int64                                        `json:"room_id"`
	ActionEvent  QianchuanTodayLiveRoomUserGetV10ActionEvent  `json:"action_event"`
	Dimension    []*QianchuanTodayLiveRoomUserGetV10Dimension `json:"dimension"`
	FlowSource   QianchuanTodayLiveRoomUserGetV10FlowSource   `json:"flow_source,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/today_live/room/user/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanTodayLiveRoomUserGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).RoomId(request.RoomId).ActionEvent(request.ActionEvent).Dimension(request.Dimension).FlowSource(request.FlowSource).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
