/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeListV2ResponseData
type StarChallengeListV2ResponseData struct {
	//
	ChallengeTasks []*StarChallengeListV2ResponseDataChallengeTasksInner `json:"challenge_tasks,omitempty"`
	PageInfo       *StarChallengeListV2ResponseDataPageInfo              `json:"page_info,omitempty"`
}
