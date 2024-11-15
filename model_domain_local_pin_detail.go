/*
 * Master Data Management
 *
 * This is Master Data Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DomainLocalPinDetail struct {
	LocalFlag string `json:"local_flag,omitempty"`
	LocalPincode int32 `json:"local_pincode,omitempty"`
	SourcePincode int32 `json:"source_pincode,omitempty"`
}
