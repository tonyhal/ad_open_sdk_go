/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueLifeCallbackV2Request struct for ToolsClueLifeCallbackV2Request
type ToolsClueLifeCallbackV2Request struct {
	ClueConvertState ToolsClueLifeCallbackV2ClueConvertState `json:"clue_convert_state"`
	//
	ClueId    string                                   `json:"clue_id"`
	EventData *ToolsClueLifeCallbackV2RequestEventData `json:"event_data,omitempty"`
	//
	LocalAccountIds []string `json:"local_account_ids"`
}
