/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2ArticleCategory
type ToolsBidSuggestV2ArticleCategory string

// List of tools_bid_suggest_v2_article_category
const (
	STORIES_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "STORIES"
	TECHNOLOGY_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "TECHNOLOGY"
	TIPS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "TIPS"
	DESIGN_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "DESIGN"
	ENTERTAINMENT_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "ENTERTAINMENT"
	LOCAL_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "LOCAL"
	GAMES_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "GAMES"
	AMUSEMENT_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "AMUSEMENT"
	TRAVEL_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "TRAVEL"
	MILITARY_ToolsBidSuggestV2ArticleCategory      ToolsBidSuggestV2ArticleCategory = "MILITARY"
	SOCIETY_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "SOCIETY"
	CARS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "CARS"
	PARENTING_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "PARENTING"
	HOME_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "HOME"
	GOURMET_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "GOURMET"
	PETS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "PETS"
	RUMOR_CRACKER_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "RUMOR_CRACKER"
	WEIGHT_LOSING_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "WEIGHT_LOSING"
	MOVIE_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "MOVIE"
	ESTATE_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "ESTATE"
	WORKPLACE_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "WORKPLACE"
	ESSAY_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "ESSAY"
	REGIMEN_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "REGIMEN"
	EXPLORE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "EXPLORE"
	CULTURE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "CULTURE"
	SCIENCE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "SCIENCE"
	GOVERNMENT_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "GOVERNMENT"
	FREAK_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "FREAK"
	HEALTH_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "HEALTH"
	EMOTION_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "EMOTION"
	COMICS_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "COMICS"
	LAWS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "LAWS"
	EDUCATION_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "EDUCATION"
	DIGITAL_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "DIGITAL"
	FASHION_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "FASHION"
	HISTORY_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "HISTORY"
	BUSINESS_ToolsBidSuggestV2ArticleCategory      ToolsBidSuggestV2ArticleCategory = "BUSINESS"
	SPORTS_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "SPORTS"
	PHOTOGRAPHY_ToolsBidSuggestV2ArticleCategory   ToolsBidSuggestV2ArticleCategory = "PHOTOGRAPHY"
	ANIMATION_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "ANIMATION"
	GRADUATES_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "GRADUATES"
	COLLECTION_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "COLLECTION"
	INTERNATIONAL_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "INTERNATIONAL"
	FINANCE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "FINANCE"
	CONSTELLATION_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "CONSTELLATION"
)

// All allowed values of ToolsBidSuggestV2ArticleCategory enum
var AllowedToolsBidSuggestV2ArticleCategoryEnumValues = []ToolsBidSuggestV2ArticleCategory{
	"STORIES",
	"TECHNOLOGY",
	"TIPS",
	"DESIGN",
	"ENTERTAINMENT",
	"LOCAL",
	"GAMES",
	"AMUSEMENT",
	"TRAVEL",
	"MILITARY",
	"SOCIETY",
	"CARS",
	"PARENTING",
	"HOME",
	"GOURMET",
	"PETS",
	"RUMOR_CRACKER",
	"WEIGHT_LOSING",
	"MOVIE",
	"ESTATE",
	"WORKPLACE",
	"ESSAY",
	"REGIMEN",
	"EXPLORE",
	"CULTURE",
	"SCIENCE",
	"GOVERNMENT",
	"FREAK",
	"HEALTH",
	"EMOTION",
	"COMICS",
	"LAWS",
	"EDUCATION",
	"DIGITAL",
	"FASHION",
	"HISTORY",
	"BUSINESS",
	"SPORTS",
	"PHOTOGRAPHY",
	"ANIMATION",
	"GRADUATES",
	"COLLECTION",
	"INTERNATIONAL",
	"FINANCE",
	"CONSTELLATION",
}

// NewToolsBidSuggestV2ArticleCategoryFromValue returns a pointer to a valid ToolsBidSuggestV2ArticleCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2ArticleCategoryFromValue(v string) (*ToolsBidSuggestV2ArticleCategory, error) {
	ev := ToolsBidSuggestV2ArticleCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2ArticleCategory: valid values are %v", v, AllowedToolsBidSuggestV2ArticleCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2ArticleCategory) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2ArticleCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_article_category value
func (v ToolsBidSuggestV2ArticleCategory) Ptr() *ToolsBidSuggestV2ArticleCategory {
	return &v
}
