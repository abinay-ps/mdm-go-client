/*
 * Master Data Management
 *
 * This is Master Data Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type HandlerMergeOfficeRequest struct {
	SourceOfficeDetails *HandlerMergeOfficeRequestSourceOfficeDetails `json:"source_office_details,omitempty"`
	TargetOfficeDetails *HandlerMergeOfficeRequestSourceOfficeDetails `json:"target_office_details,omitempty"`
}
