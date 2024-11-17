/*
 * Master Data Management API
 *
 * A comprehensive API suite for managing various aspects of the postal system, including office management, product management, tariffs (domestic and international), product country mapping, local pin codes, country details, and currency exchange rates. Additionally, the API supports insurance quote generation for PLI/RPLI, and interest rate calculations for savings schemes such as SB, SSA, PPF, MIS, TD, NSC/KVP, SCSS, and MSS.
 *
 * API version: 1.0
 * Contact: support_cept@indiapost.gov.in
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ResponseListLocalPincodesResponse struct {
	Latitude float32 `json:"latitude,omitempty"`
	LocalPinId int32 `json:"local_pin_id,omitempty"`
	LocalPincode float32 `json:"local_pincode,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
	OfficeId int32 `json:"office_id,omitempty"`
	OfficeName string `json:"office_name,omitempty"`
	OfficeTypeCode string `json:"office_type_code,omitempty"`
	SourcePincode float32 `json:"source_pincode,omitempty"`
}