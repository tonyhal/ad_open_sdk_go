/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerTrackUrlUpdateV2Request struct for EventManagerTrackUrlUpdateV2Request
type EventManagerTrackUrlUpdateV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 资产ID
	AssetId int64 `json:"asset_id"`
	// 应用下载链接，应用必填
	DownloadUrl   *string                                          `json:"download_url,omitempty"`
	TrackUrlGroup EventManagerTrackUrlUpdateV2RequestTrackUrlGroup `json:"track_url_group"`
}
