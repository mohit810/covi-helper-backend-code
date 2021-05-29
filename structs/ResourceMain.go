package structs

type ResourceMain struct {
	ResponseCode   int               `json:"responseCode"`
	CityCode       string            `json:"cityCode"`
	StateName      string            `json:"stateName"`
	CityName       string            `json:"cityName"`
	ResourceMaster []*ResourceMaster `json:"resourceMaster"`
}
