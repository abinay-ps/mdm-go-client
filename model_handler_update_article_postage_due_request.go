/*
 * Postman Delivery Management
 *
 * This is Postman Delivery Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type HandlerUpdateArticlePostageDueRequest struct {
	OfficeId int32 `json:"office_id"`
	PostageDue float32 `json:"postage_due"`
	UpdatedBy string `json:"updated_by"`
}