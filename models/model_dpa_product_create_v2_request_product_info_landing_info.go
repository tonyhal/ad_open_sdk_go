/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductCreateV2RequestProductInfoLandingInfo
type DpaProductCreateV2RequestProductInfoLandingInfo struct {
	//
	TargetUrl *string `json:"target_url,omitempty"`
	//
	TargetUrlAndroidApp *string `json:"target_url_android_app,omitempty"`
	//
	TargetUrlIosApp *string `json:"target_url_ios_app,omitempty"`
	//
	TargetUrlMobile *string `json:"target_url_mobile,omitempty"`
	//
	TargetUrlUniversalLink *string `json:"target_url_universal_link,omitempty"`
}
