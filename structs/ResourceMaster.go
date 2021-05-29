package structs

type ResourceMaster struct {
	Oxygen      []*ResourceLead `json:"oxygen"`
	Plasma      []*ResourceLead `json:"plasma"`
	Medicine    []*ResourceLead `json:"medicine"`
	Bed         []*ResourceLead `json:"bed"`
	Ambulance   []*ResourceLead `json:"ambulance"`
	HelpingHand []*ResourceLead `json:"helpingHand"`
}
