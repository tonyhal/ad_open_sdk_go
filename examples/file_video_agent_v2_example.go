/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
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

type ApiOpenApi2FileVideoAgentPostRequestExample struct {
	AgentId        int64                      `json:"agent_id"`
	FileName       string                     `json:"file_name"`
	IsNeedAuth     bool                       `json:"is_need_auth"`
	IsAigc         bool                       `json:"is_aigc,omitempty"`
	UploadType     FileVideoAgentV2UploadType `json:"upload_type,omitempty"`
	VideoFile      *FormFileInfo              `json:"video_file,omitempty"`
	VideoSignature string                     `json:"video_signature,omitempty"`
	VideoUrl       string                     `json:"video_url,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/file/video/agent/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2FileVideoAgentPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.FileVideoAgentV2Api().
		Post(ctx).
		AccessToken(accessToken).
		AgentId(request.AgentId).FileName(request.FileName).IsNeedAuth(request.IsNeedAuth).IsAigc(request.IsAigc).UploadType(request.UploadType).VideoFile(request.VideoFile).VideoSignature(request.VideoSignature).VideoUrl(request.VideoUrl).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
