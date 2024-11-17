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

type ResponseListPartialOfficeDetailsResponse struct {
	CityName string `json:"city_name,omitempty"`
	DeliveryOfficeFlag bool `json:"delivery_office_flag,omitempty"`
	OfficeId int32 `json:"office_id"`
	OfficeName string `json:"office_name"`
	OfficeTypeCode string `json:"office_type_code"`
	Pincode int32 `json:"pincode"`
	ReportingOfficeId int32 `json:"reporting_office_id,omitempty"`
	ReportingOfficeName string `json:"reporting_office_name,omitempty"`
	StateName string `json:"state_name"`
	TalukName string `json:"taluk_name,omitempty"`
	VillageName string `json:"village_name,omitempty"`
}