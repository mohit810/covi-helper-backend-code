package structs

type VerifyResource struct {
	ResponseCode   int               `json:"responseCode"`
	ResourceMaster []*ResourceMaster `json:"resourceMaster"`
}
