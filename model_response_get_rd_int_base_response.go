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

type ResponseGetRdIntBaseResponse struct {
	AcOpenedDuring string `json:"ac_opened_during,omitempty"`
	Duration int32 `json:"duration,omitempty"`
	Interest float32 `json:"interest,omitempty"`
	MaturityAmount float32 `json:"maturity_amount,omitempty"`
	Principle float32 `json:"principle,omitempty"`
}