/*
 * Postman Delivery Management
 *
 * This is Postman Delivery Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApierrorsAppError struct {
	Code string `json:"code,omitempty"`
	FieldErrors []ApierrorsFieldError `json:"field_errors,omitempty"`
	Message string `json:"message,omitempty"`
}
