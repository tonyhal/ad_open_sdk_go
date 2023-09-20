/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInnerSceneListInnerSuggestionsInner struct for ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInnerSceneListInnerSuggestionsInner
type ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInnerSceneListInnerSuggestionsInner struct {
	//
	Conclusion *string `json:"conclusion,omitempty"`
	//
	Msg *string `json:"msg,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	SceneType *string                                                                                `json:"scene_type,omitempty"`
	ToolType  *ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType `json:"tool_type,omitempty"`
	//
	Tools []*ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInnerSceneListInnerSuggestionsInnerToolsInner `json:"tools,omitempty"`
}
