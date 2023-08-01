/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"fmt"
	"io"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

type ApiOpenApiOauth2RefreshTokenPostRequestExample struct {
	Oauth2RefreshTokenRequest Oauth2RefreshTokenRequest `json:"Oauth2RefreshTokenRequest,omitempty"`
}

// url: https://api.oceanengine.com/open_api/oauth2/refresh_token/ Post
func main() {
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiOauth2RefreshTokenPostRequestExample
	request.Oauth2RefreshTokenRequest.RefreshToken = ""
	request.Oauth2RefreshTokenRequest.Secret = ""
	request.Oauth2RefreshTokenRequest.AppId = PtrInt64(1)
	resp, httpRes, err := apiClient.Oauth2RefreshTokenApi().
		Post(ctx).
		Oauth2RefreshTokenRequest(request.Oauth2RefreshTokenRequest).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}