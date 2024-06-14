/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInnerSceneListInnerSuggestionsInner struct for ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInnerSceneListInnerSuggestionsInner
type ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInnerSceneListInnerSuggestionsInner struct {
	// 诊断结论，当scene=POTENTIAL时，该字段为空
	Conclusion *string `json:"conclusion,omitempty"`
	// 工具详细描述
	Msg *string `json:"msg,omitempty"`
	// 建议操作名称
	Name *string `json:"name,omitempty"`
	// 场景下的细分类型，允许值：  当 scene=CLEAN时，值为学习期失败率高，空耗，活跃度低  当 scene=POTENTIAL时，值为潜力计划 当 scene=PROBLEM时，值为起量困难，掉量，衰减
	SceneType *string                                                                      `json:"scene_type,omitempty"`
	ToolType  *ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType `json:"tool_type,omitempty"`
	// 工具对应的参数集合
	Tools []*ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInnerSceneListInnerSuggestionsInnerToolsInner `json:"tools,omitempty"`
}
