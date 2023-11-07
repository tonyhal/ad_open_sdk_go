/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteCreateV2RequestBricksInner struct for ToolsSiteCreateV2RequestBricksInner
type ToolsSiteCreateV2RequestBricksInner struct {
	//
	AutoPlay   *int64                                         `json:"auto_play,omitempty"`
	Background *ToolsSiteCreateV2RequestBricksInnerBackground `json:"background,omitempty"`
	//
	BgColor *string `json:"bg_color,omitempty"`
	//
	BgImage *string `json:"bg_image,omitempty"`
	//
	BgType *string `json:"bg_type,omitempty"`
	//
	BorderColor *string `json:"border_color,omitempty"`
	//
	BorderRadius *int64 `json:"border_radius,omitempty"`
	//
	BorderWidth *int64 `json:"border_width,omitempty"`
	//
	BulbEffect *bool `json:"bulb_effect,omitempty"`
	//
	Color *string `json:"color,omitempty"`
	//
	Comments []*ToolsSiteCreateV2RequestBricksInnerCommentsInner `json:"comments,omitempty"`
	//
	Content *string `json:"content,omitempty"`
	//
	Elements []*ToolsSiteCreateV2RequestBricksInnerElementsInner `json:"elements,omitempty"`
	//
	Events      []*ToolsSiteCreateV2RequestBricksInnerEventsInner `json:"events,omitempty"`
	FailureLink *ToolsSiteCreateV2RequestBricksInnerFailureLink   `json:"failure_link,omitempty"`
	//
	Float *string `json:"float,omitempty"`
	//
	FontFamily *string `json:"font_family,omitempty"`
	//
	FontSize *int64 `json:"font_size,omitempty"`
	//
	FontStyle []string                                     `json:"font_style,omitempty"`
	FormData  *ToolsSiteCreateV2RequestBricksInnerFormData `json:"form_data,omitempty"`
	//
	GroupContent []*ToolsSiteCreateV2RequestBricksInnerGroupContentInner `json:"group_content,omitempty"`
	//
	GroupType *string `json:"group_type,omitempty"`
	//
	Height *int64 `json:"height,omitempty"`
	//
	IcId *string `json:"ic_id,omitempty"`
	//
	Icon *string `json:"icon,omitempty"`
	//
	ImageUrl *string `json:"image_url,omitempty"`
	//
	IsCover *int64 `json:"is_cover,omitempty"`
	//
	LetterSpacing *int64 `json:"letter_spacing,omitempty"`
	//
	LineHeight *int64                                   `json:"line_height,omitempty"`
	Link       *ToolsSiteCreateV2RequestBricksInnerLink `json:"link,omitempty"`
	//
	Linkable    *int64                                          `json:"linkable,omitempty"`
	LocalSource *ToolsSiteCreateV2RequestBricksInnerLocalSource `json:"local_source,omitempty"`
	Marquee     *ToolsSiteCreateV2RequestBricksInnerMarquee     `json:"marquee,omitempty"`
	//
	Name string `json:"name"`
	//
	OffsetX *int64 `json:"offset_x,omitempty"`
	//
	OffsetY      *int64                                           `json:"offset_y,omitempty"`
	OnlineSource *ToolsSiteCreateV2RequestBricksInnerOnlineSource `json:"online_source,omitempty"`
	PackageInfo  *ToolsSiteCreateV2RequestBricksInnerPackageInfo  `json:"package_info,omitempty"`
	//
	Rewards []*ToolsSiteCreateV2RequestBricksInnerRewardsInner `json:"rewards,omitempty"`
	//
	RewardsButtonText *string `json:"rewards_button_text,omitempty"`
	//
	RewardsButtonTextColor *string `json:"rewards_button_text_color,omitempty"`
	//
	RewardsButtonTextFontSize *int64                                          `json:"rewards_button_text_font_size,omitempty"`
	RuleText                  *ToolsSiteCreateV2RequestBricksInnerRuleText    `json:"rule_text,omitempty"`
	SuccessLink               *ToolsSiteCreateV2RequestBricksInnerSuccessLink `json:"success_link,omitempty"`
	//
	SuccessTip *string `json:"success_tip,omitempty"`
	//
	Text *string `json:"text,omitempty"`
	//
	TextAlign *string `json:"text_align,omitempty"`
	//
	Width *int64 `json:"width,omitempty"`
	//
	X *int64 `json:"x,omitempty"`
	//
	Y *int64 `json:"y,omitempty"`
}
