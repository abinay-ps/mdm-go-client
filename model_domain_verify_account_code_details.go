/*
 * Master Data Management
 *
 * This is Master Data Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DomainVerifyAccountCodeDetails struct {
	ApprovedBy string `json:"approved_by,omitempty"`
	GlCode int32 `json:"gl_code,omitempty"`
	GlStatus string `json:"gl_status,omitempty"`
}
