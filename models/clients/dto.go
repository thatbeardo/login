package dto

// User is the data sent by the App to the backend
type User struct {
	FirstName string `json:"firstName" mapstructure:"name"`
	LastName  string `json:"lastName" binding:"required"`
	EmailID   string `json:"emailID" mapstructure:"context_id"`
}
