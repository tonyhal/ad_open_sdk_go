/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsVideoPredictV2ResponseDataListInner struct for ToolsVideoPredictV2ResponseDataListInner
type ToolsVideoPredictV2ResponseDataListInner struct {
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	Score *string `json:"score,omitempty"`
	//
	ScoreAppeal *string `json:"score_appeal,omitempty"`
	//
	ScoreBasicVideo *string `json:"score_basic_video,omitempty"`
	//
	ScoreEmotion *string `json:"score_emotion,omitempty"`
	//
	ScoreHighOrder *string `json:"score_high_order,omitempty"`
	//
	ScoreStory *string `json:"score_story,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	//
	WeightAppeal *string `json:"weight_appeal,omitempty"`
	//
	WeightBasicVideo *string `json:"weight_basic_video,omitempty"`
	//
	WeightEmotion *string `json:"weight_emotion,omitempty"`
	//
	WeightHighOrder *string `json:"weight_high_order,omitempty"`
	//
	WeightStory *string `json:"weight_story,omitempty"`
}
