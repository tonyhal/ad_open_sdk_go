/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV2V2ResponseDataSuccessListInner struct for KeywordCreateV2V2ResponseDataSuccessListInner
type KeywordCreateV2V2ResponseDataSuccessListInner struct {
	//
	Bid     *float64                                 `json:"bid,omitempty"`
	BidType *KeywordCreateV2V2DataSuccessListBidType `json:"bid_type,omitempty"`
	//
	IsPause *int64 `json:"is_pause,omitempty"`
	//
	KeywordId *int64                                     `json:"keyword_id,omitempty"`
	MatchType *KeywordCreateV2V2DataSuccessListMatchType `json:"match_type,omitempty"`
	//
	Word string `json:"word"`
}
