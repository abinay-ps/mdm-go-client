/*
 * Postman Delivery Management
 *
 * This is Postman Delivery Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ResponseListOrdinaryArticlesApiResponse struct {
	Data *ResponseListOrdinaryArticlesResponse `json:"data,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	Message string `json:"message,omitempty"`
	OrderBy string `json:"order_by,omitempty"`
	ReturnedRecordsCount int32 `json:"returned_records_count,omitempty"`
	Skip int32 `json:"skip,omitempty"`
	SortType string `json:"sort_type,omitempty"`
	StatusCode int32 `json:"status_code,omitempty"`
	TotalRecordsCount int32 `json:"total_records_count,omitempty"`
}
