/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportRubeexGetV2ResponseDataListInnerMetrics
type ReportRubeexGetV2ResponseDataListInnerMetrics struct {
	//
	ClickCnt *int64 `json:"click_cnt,omitempty"`
	//
	ConversionCost *float64 `json:"conversion_cost,omitempty"`
	//
	ConversionRate *float64 `json:"conversion_rate,omitempty"`
	//
	ConvertCnt *int64 `json:"convert_cnt,omitempty"`
	//
	Cost *float64 `json:"cost,omitempty"`
	//
	Cpc *float64 `json:"cpc,omitempty"`
	//
	Cpm *float64 `json:"cpm,omitempty"`
	//
	Ctr *float64 `json:"ctr,omitempty"`
	//
	CustomFinishPlayPlayableCount *int64 `json:"custom_finish_play_playable_count,omitempty"`
	//
	CustomFinishPlayPlayableLayerRate *float64 `json:"custom_finish_play_playable_layer_rate,omitempty"`
	//
	CustomFinishPlayPlayableRate *float64 `json:"custom_finish_play_playable_rate,omitempty"`
	//
	CustomLoadMainSceneCount *int64 `json:"custom_load_main_scene_count,omitempty"`
	//
	CustomLoadMainSceneRate *float64 `json:"custom_load_main_scene_rate,omitempty"`
	//
	CustomStartPlayPlayableCount *int64 `json:"custom_start_play_playable_count,omitempty"`
	//
	CustomStartPlayPlayableLayerRate *float64 `json:"custom_start_play_playable_layer_rate,omitempty"`
	//
	CustomStartPlayPlayableRate *float64 `json:"custom_start_play_playable_rate,omitempty"`
	//
	DownloadButtonClickCountAll *int64 `json:"download_button_click_count_all,omitempty"`
	//
	EnterSectionCount *int64 `json:"enter_section_count,omitempty"`
	//
	PlayDurationAvgSdk *float64 `json:"play_duration_avg_sdk,omitempty"`
	//
	PlayDurationSdk *float64 `json:"play_duration_sdk,omitempty"`
	//
	PlayableCtr *float64 `json:"playable_ctr,omitempty"`
	//
	PlayableDurationUserCount *int64 `json:"playable_duration_user_count,omitempty"`
	//
	PrePlayDurationAvgSdk *float64 `json:"pre_play_duration_avg_sdk,omitempty"`
	//
	PrePlayDurationSdk *float64 `json:"pre_play_duration_sdk,omitempty"`
	//
	PredefinePageviewCount *int64 `json:"predefine_pageview_count,omitempty"`
	//
	SdkPlayableShowCount *int64 `json:"sdk_playable_show_count,omitempty"`
	//
	ShowCnt *int64 `json:"show_cnt,omitempty"`
}
