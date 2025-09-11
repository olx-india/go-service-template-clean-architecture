package dto

// CheckLimitRequest represents the request for checking limit.
type CheckLimitRequest struct {
	UserID int `json:"userID" binding:"required"`
}

// CheckLimitResponse represents the response for limit.
type CheckLimitResponse struct {
	UserID         int `json:"userID"`
	LimitAvailable int `json:"limitAvailable"`
}
