package structs

type Volunteer struct {
	EmailId     string `json:"email"`
	Name        string `json:"name"`
	UserId      string `json:"id"`
	Picture     string `json:"picture"`
	PhoneNumber string `json:"phoneNumber"`
	Allowed     bool   `json:"allowed"`
}
