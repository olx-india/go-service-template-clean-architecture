package dto

// CheckLimitRequest represents the request for checking limit.
type CheckLimitRequest struct {
	UserID int `json:"userID" binding:"required" example:"1" description:"User identity"`
}

// CheckLimitResponse represents the response for limit.
type CheckLimitResponse struct {
	UserID         int `json:"userID" example:"1" description:"User identity"`
	LimitAvailable int `json:"limitAvailable" example:"10" description:"Remaining limit units"`
}
