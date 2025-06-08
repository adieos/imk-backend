package dto

import (
	"errors"
)

const (
	MESSAGE_FAILED_GET_DATA_FROM_PATH     = "Failed to get data from request path"
	MESSAGE_FAILED_GET_DATA_FROM_QUERY    = "Failed to get data from request query"
	MESSAGE_SUCCESS_CREATE_BS             = "Bank Sampah created successfully"
	MESSAGE_FAILED_CREATE_BS              = "Failed to create Bank Sampah"
	MESSAGE_SUCCESS_GET_BS_BY_ID          = "Bank Sampah retrieved successfully"
	MESSAGE_FAILED_GET_BS_BY_ID           = "Failed to retrieve Bank Sampah by ID"
	MESSAGE_SUCCESS_GET_ALL_BS            = "All Bank Sampah retrieved successfully"
	MESSAGE_FAILED_GET_ALL_BS             = "Failed to retrieve all Bank Sampah"
	MESSAGE_SUCCESS_GET_ALL_BS_BY_USER_ID = "All Bank Sampah by user ID retrieved successfully"
	MESSAGE_FAILED_GET_ALL_BS_BY_USER_ID  = "Failed to retrieve all Bank Sampah by user ID"
	MESSAGE_SUCCESS_UPDATE_BS             = "Bank Sampah updated successfully"
	MESSAGE_FAILED_UPDATE_BS              = "Failed to update Bank Sampah"
	MESSAGE_SUCCESS_DELETE_BS             = "Bank Sampah deleted successfully"
	MESSAGE_FAILED_DELETE_BS              = "Failed to delete Bank Sampah"
	MESSAGE_SUCCESS_GET_ALL_WASTE_TYPES   = "All waste types retrieved successfully"
	MESSAGE_FAILED_GET_ALL_WASTE_TYPES    = "Failed to retrieve all waste types"
	MESSAGE_SUCCESS_GET_WASTE_TYPE_BY_ID  = "Waste type retrieved successfully"
	MESSAGE_FAILED_GET_WASTE_TYPE_BY_ID   = "Failed to retrieve waste type by ID"
	MESSAGE_SUCCESS_CREATE_WASTE_TYPE     = "Waste type created successfully"
	MESSAGE_FAILED_CREATE_WASTE_TYPE      = "Failed to create waste type"
	MESSAGE_SUCCESS_UPDATE_WASTE_TYPE     = "Waste type updated successfully"
	MESSAGE_FAILED_UPDATE_WASTE_TYPE      = "Failed to update waste type"
	MESSAGE_SUCCESS_DELETE_WASTE_TYPE     = "Waste type deleted successfully"
	MESSAGE_FAILED_DELETE_WASTE_TYPE      = "Failed to delete waste type"
	MESSAGE_SUCCESS_GET_BS_ACCEPTS        = "Bank Sampah accepts retrieved successfully"
	MESSAGE_FAILED_GET_BS_ACCEPTS         = "Failed to retrieve Bank Sampah accepts"
	MESSAGE_SUCCESS_CHANGE_STATUS_BS      = "Bank Sampah status changed successfully"
	MESSAGE_FAILED_CHANGE_STATUS_BS       = "Failed to change Bank Sampah status"
	MESSAGE_SUCCESS_GET_BS_ACCEPT_BY_ID   = "Bank Sampah accept retrieved successfully"
	MESSAGE_FAILED_GET_BS_ACCEPT_BY_ID    = "Failed to retrieve Bank Sampah accept by ID"
	MESSAGE_SUCCESS_CREATE_BS_ACCEPT      = "Bank Sampah accept created successfully"
	MESSAGE_FAILED_CREATE_BS_ACCEPT       = "Failed to create Bank Sampah accept"
	MESSAGE_SUCCESS_UPDATE_BS_ACCEPT      = "Bank Sampah accept updated successfully"
	MESSAGE_FAILED_UPDATE_BS_ACCEPT       = "Failed to update Bank Sampah accept"
	MESSAGE_SUCCESS_DELETE_BS_ACCEPT      = "Bank Sampah accept deleted successfully"
)

var (
	ErrBankSampahNotFound = errors.New("bank sampah not found")
)

type (
	BSCreateRequest struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		City    string `json:"city"`
		Phone   string `json:"phone"` // Phone number should be validated
	}

	BSWasteType struct {
		CategoryID string `json:"category_id"` // Category ID should be validated
		Price      int    `json:"price"`       // Price should be a positive integer
		Quota      int    `json:"quota"`       // Quota should be a positive integer
		Filled     int    `json:"filled"`      // Filled should be a boolean
	}

	BSUpdateRequest struct {
		Id                 string        `json:"id"`
		Name               string        `json:"name"`
		OpenHour           string        `json:"open_hour"`            // Open hour should be in HH:MM forma
		Phone              string        `json:"phone"`                // Phone number should be validated
		Description        string        `json:"description"`          // Description should be optional
		AcceptAll          bool          `json:"accept_all"`           // Accept all types of waste
		AcceptedWasteTypes []BSWasteType `json:"accepted_waste_types"` // Accepted waste types should be a list of strings
	}

	BSResponse struct {
		Id                 string        `json:"id"`
		Name               string        `json:"name"`
		Address            string        `json:"address"`
		City               string        `json:"city"`
		Phone              string        `json:"phone"`                // Phone number should be validated
		OpenHour           string        `json:"open_hour"`            // Open hour should be in HH:MM format
		Description        string        `json:"description"`          // Description should be optional
		AcceptAll          bool          `json:"accept_all"`           // Accept all types of waste
		AcceptedWasteTypes []BSWasteType `json:"accepted_waste_types"` // Accepted waste types should be a list of strings
		Status             string        `json:"status"`               // Status should be a string
	}
)
