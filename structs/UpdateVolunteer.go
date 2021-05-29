package structs

type UpdateVolunteer struct {
	ResponseCode int    `json:"responseCode"`
	Email        string `json:"email"`
	Allowed      bool   `json:"allowed"`
}
