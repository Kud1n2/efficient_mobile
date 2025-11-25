package main

// WithDate represents a subscription with dates
// @Description Subscription information with start and end dates
type WithDate struct {
	Service_name string `json:"service_name"`
	Price        int    `json:"price"`
	User_id      string `json:"user_id"`
	Start_date   string `json:"start_date"`
	Finish_date  string `json:"finish_date,omitempty"`
}

// FindByName represents request body for finding by service name
// @Description Request to find subscription by service name
type FindByName struct {
	Service_name string `json:"service_name"`
}

// UpdateByName represents request body for updating subscription
// @Description Request to update subscription information
type UpdateByName struct {
	Old_service_name string `json:"old_service_name"`
	Service_name     string `json:"service_name,omitempty"`
	Price            *int   `json:"price,omitempty"`
	User_id          string `json:"user_id,omitempty"`
	Start_date       string `json:"start_date,omitempty"`
	Finish_date      string `json:"finish_date,omitempty"`
}

// GetSum represents request body for calculating sum
// @Description Request to calculate total subscription costs with filters
type GetSum struct {
	Service_name string `json:"service_name,omitempty"`
	User_id      string `json:"user_id,omitempty"`
	Start_date   string `json:"start_date"`
	Finish_date  string `json:"finish_date"`
}

// SuccessResponse represents a successful operation response
// @Description Standard success response format
type SuccessResponse struct {
	Message string `json:"message" example:"Operation completed successfully"`
}

// ErrorResponse represents an error response
// @Description Standard error response format
type ErrorResponse struct {
	Error   string `json:"error" example:"Database error"`
	Message string `json:"message" example:"Could not connect to database"`
}

// SumResponse represents total cost calculation response
// @Description Response format for sum calculation
type SumResponse struct {
	Sum int `json:"sum" example:"4500"`
}
