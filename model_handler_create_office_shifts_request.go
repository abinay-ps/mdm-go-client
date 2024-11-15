/*
 * Postman Delivery Management
 *
 * This is Postman Delivery Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type HandlerCreateOfficeShiftsRequest struct {
	CreatedBy string `json:"created_by,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
	IsShiftBegin bool `json:"is_shift_begin,omitempty"`
	IsShiftEnd bool `json:"is_shift_end,omitempty"`
	OfficeId int32 `json:"office_id,omitempty"`
	OfficeName string `json:"office_name,omitempty"`
	ReasonForDelayShiftEnd string `json:"reason_for_delay_shift_end,omitempty"`
	ShiftEndDate string `json:"shift_end_date,omitempty"`
	ShiftStartDate string `json:"shift_start_date,omitempty"`
	ShiftType *HandlerCreateOfficeShiftsRequestShiftType `json:"shift_type"`
}