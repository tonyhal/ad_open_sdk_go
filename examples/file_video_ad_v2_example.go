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

type ApiOpenApi2FileVideoAdPostRequestExample struct {
	AdvertiserId   int64                   `json:"advertiser_id"`
	Filename       string                  `json:"filename,omitempty"`
	IsAigc         bool                    `json:"is_aigc,omitempty"`
	IsGuideVideo   bool                    `json:"is_guide_video,omitempty"`
	Labels         []string                `json:"labels,omitempty"`
	UploadType     FileVideoAdV2UploadType `json:"upload_type,omitempty"`
	VideoFile      *FormFileInfo           `json:"video_file,omitempty"`
	VideoSignature string                  `json:"video_signature,omitempty"`
	VideoUrl       string                  `json:"video_url,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/file/video/ad/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2FileVideoAdPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.FileVideoAdV2Api().
		Post(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Filename(request.Filename).IsAigc(request.IsAigc).IsGuideVideo(request.IsGuideVideo).Labels(request.Labels).UploadType(request.UploadType).VideoFile(request.VideoFile).VideoSignature(request.VideoSignature).VideoUrl(request.VideoUrl).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
