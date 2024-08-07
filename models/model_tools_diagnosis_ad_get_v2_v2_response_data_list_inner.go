/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisAdGetV2V2ResponseDataListInner struct for ToolsDiagnosisAdGetV2V2ResponseDataListInner
type ToolsDiagnosisAdGetV2V2ResponseDataListInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdName *string `json:"ad_name,omitempty"`
	//
	Bid       *float64                                               `json:"bid,omitempty"`
	BidResult *ToolsDiagnosisAdGetV2V2ResponseDataListInnerBidResult `json:"bid_result,omitempty"`
	//
	Budget       *float64                                                  `json:"budget,omitempty"`
	BudgetResult *ToolsDiagnosisAdGetV2V2ResponseDataListInnerBudgetResult `json:"budget_result,omitempty"`
	//
	CostBias  *float64                                               `json:"cost_bias,omitempty"`
	CtrResult *ToolsDiagnosisAdGetV2V2ResponseDataListInnerCtrResult `json:"ctr_result,omitempty"`
	CvrResult *ToolsDiagnosisAdGetV2V2ResponseDataListInnerCvrResult `json:"cvr_result,omitempty"`
	//
	DiagnosisTime *int64 `json:"diagnosis_time,omitempty"`
	//
	DifferenceRatioShow *float64 `json:"difference_ratio_show,omitempty"`
	//
	MatureShow      *float64                                                     `json:"mature_show,omitempty"`
	PotentialResult *ToolsDiagnosisAdGetV2V2ResponseDataListInnerPotentialResult `json:"potential_result,omitempty"`
	QualityResult   *ToolsDiagnosisAdGetV2V2ResponseDataListInnerQualityResult   `json:"quality_result,omitempty"`
	//
	RecSuggestionContent *string `json:"rec_suggestion_content,omitempty"`
	//
	RecSuggestionMode *string `json:"rec_suggestion_mode,omitempty"`
	//
	RecSuggestionReason *string `json:"rec_suggestion_reason,omitempty"`
	//
	SceneDesc *string `json:"scene_desc,omitempty"`
	//
	SceneText *string `json:"scene_text,omitempty"`
	//
	Show *int64 `json:"show,omitempty"`
	//
	SimilarRatioShow *float64 `json:"similar_ratio_show,omitempty"`
	//
	Summary      []string                                                  `json:"summary,omitempty"`
	TargetResult *ToolsDiagnosisAdGetV2V2ResponseDataListInnerTargetResult `json:"target_result,omitempty"`
}
