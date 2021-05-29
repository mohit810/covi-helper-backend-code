package structs

type ResourceLead struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	URL            string `json:"url"`
	PhoneNumber    string `json:"phoneNumber"`
	LeadType       string `json:"lead_type"`
	WhatsappNumber string `json:"whatsappNumber"`
	Website        string `json:"website"`
	Notes          string `json:"notes"`
	CreationTime   string `json:"creationTime"`
	Statecode      int    `json:"statecode"`
	Citycode       string `json:"citycode"`
	VerifiedBy     string `json:"verified_by"`
	VerfiedTime    string `json:"verfied_time"`
}
