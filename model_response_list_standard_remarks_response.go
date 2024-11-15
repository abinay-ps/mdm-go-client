/*
 * Postman Delivery Management
 *
 * This is Postman Delivery Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ResponseListStandardRemarksResponse struct {
	ActionCode int32 `json:"action_code,omitempty"`
	ActionToBeTaken string `json:"action_to_be_taken,omitempty"`
	BoCategory int32 `json:"bo_category,omitempty"`
	Category int32 `json:"category,omitempty"`
	Remarks string `json:"remarks,omitempty"`
	RemarksId string `json:"remarks_id,omitempty"`
	TreatmentOfTheArticleForThatDay string `json:"treatment_of_the_article_for_that_day,omitempty"`
}