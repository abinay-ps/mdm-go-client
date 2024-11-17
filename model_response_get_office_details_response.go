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

type ResponseGetOfficeDetailsResponse struct {
	AccountingOfficeId int32 `json:"accounting_office_id,omitempty"`
	AccountingOfficeName string `json:"accounting_office_name,omitempty"`
	AdminFlag bool `json:"admin_flag,omitempty"`
	AdminOfficeId int32 `json:"admin_office_id,omitempty"`
	ApprovalStatus string `json:"approval_status,omitempty"`
	ApprovedBy string `json:"approved_by,omitempty"`
	ApprovedDate string `json:"approved_date,omitempty"`
	AtmId string `json:"atm_id,omitempty"`
	BoId int32 `json:"bo_id,omitempty"`
	BoName string `json:"bo_name,omitempty"`
	CircleCode string `json:"circle_code,omitempty"`
	CircleName string `json:"circle_name,omitempty"`
	CircleOfficeId int32 `json:"circle_office_id,omitempty"`
	CityName string `json:"city_name,omitempty"`
	ClosedDate string `json:"closed_date,omitempty"`
	ContactNumber string `json:"contact_number,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
	CsiFacilityId string `json:"csi_facility_id,omitempty"`
	Dac string `json:"dac,omitempty"`
	DdoCode string `json:"ddo_code,omitempty"`
	DdoOfficeId int32 `json:"ddo_office_id,omitempty"`
	DdoOfficeName string `json:"ddo_office_name,omitempty"`
	DeletedFlag bool `json:"deleted_flag,omitempty"`
	DeliveryOfficeFlag bool `json:"delivery_office_flag,omitempty"`
	DivisionName string `json:"division_name,omitempty"`
	DivisionOfficeId int32 `json:"division_office_id,omitempty"`
	EmailId string `json:"email_id,omitempty"`
	GstnCode string `json:"gstn_code,omitempty"`
	HoId int32 `json:"ho_id,omitempty"`
	HoName string `json:"ho_name,omitempty"`
	HroId int32 `json:"hro_id,omitempty"`
	HroName string `json:"hro_name,omitempty"`
	Landmark string `json:"landmark,omitempty"`
	Latitude float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
	ModifiedFlag bool `json:"modified_flag,omitempty"`
	OfficeAddress1 string `json:"office_address1,omitempty"`
	OfficeAddress2 string `json:"office_address2,omitempty"`
	OfficeAddress3 string `json:"office_address3,omitempty"`
	OfficeClass string `json:"office_class,omitempty"`
	OfficeHierarchyId int32 `json:"office_hierarchy_id,omitempty"`
	OfficeId int32 `json:"office_id"`
	OfficeLevel string `json:"office_level,omitempty"`
	OfficeLocationId int32 `json:"office_location_id,omitempty"`
	OfficeName string `json:"office_name,omitempty"`
	OfficeStatus string `json:"office_status,omitempty"`
	OfficeStatusId int32 `json:"office_status_id,omitempty"`
	OfficeStreet string `json:"office_street,omitempty"`
	OfficeTypeCode string `json:"office_type_code,omitempty"`
	OfficeTypeId int32 `json:"office_type_id,omitempty"`
	OrderMemoNumber string `json:"order_memo_number,omitempty"`
	PaoCode string `json:"pao_code,omitempty"`
	PaoOfficeId string `json:"pao_office_id,omitempty"`
	Pincode int32 `json:"pincode,omitempty"`
	PliId string `json:"pli_id,omitempty"`
	QrTerminalId string `json:"qr_terminal_id,omitempty"`
	RegionName string `json:"region_name,omitempty"`
	RegionOfficeId int32 `json:"region_office_id,omitempty"`
	Remarks string `json:"remarks,omitempty"`
	ReportingOfficeId int32 `json:"reporting_office_id,omitempty"`
	ReportingOfficeName string `json:"reporting_office_name,omitempty"`
	SoId int32 `json:"so_id,omitempty"`
	SoName string `json:"so_name,omitempty"`
	SolId string `json:"sol_id,omitempty"`
	SroId int32 `json:"sro_id,omitempty"`
	SroName string `json:"sro_name,omitempty"`
	StateName string `json:"state_name,omitempty"`
	SubDivisionName string `json:"sub_division_name,omitempty"`
	SubDivisionOfficeId int32 `json:"sub_division_office_id,omitempty"`
	SupportedDocumentPath string `json:"supported_document_path,omitempty"`
	TalukName string `json:"taluk_name,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	UpdatedDate string `json:"updated_date,omitempty"`
	ValidFrom string `json:"valid_from,omitempty"`
	ValidTo string `json:"valid_to,omitempty"`
	VillageName string `json:"village_name,omitempty"`
	WegCode string `json:"weg_code,omitempty"`
	WorkingDays string `json:"working_days,omitempty"`
	WorkingHoursFrom string `json:"working_hours_from,omitempty"`
	WorkingHoursTo string `json:"working_hours_to,omitempty"`
}