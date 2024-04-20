package response

import db "myclass/db/sqlc"

type AddAvailabilityResponse struct {
	Date        string `json:"date"`
	UserId      int64  `json:"user_id"`
	Username    string `json:"username"`
	Start       string `json:"start"`
	EndTime     string `json:"end"`
	IsAvailable bool   `json:"is_available"`
}

func NewAvaialibilityResponse(professionalAvailability db.Availability) AddAvailabilityResponse {
	return AddAvailabilityResponse{
		Date:        professionalAvailability.Date,
		UserId:      professionalAvailability.UserID,
		Username:    professionalAvailability.Username,
		Start:       professionalAvailability.Start,
		EndTime:     professionalAvailability.EndTime,
		IsAvailable: professionalAvailability.IsAvailable,
	}
}
