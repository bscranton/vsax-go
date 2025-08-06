package vsax

type Metadata struct {
	TotalCount int
	NextQueryLink string
	ResponseCode int
}

type Device struct {
	Identifier string
	Name string
	GroupId int
	GroupName string
	IsAgentInstalled bool
	IsMdmEnrolled bool
	SiteId int
	SiteName string
	OrganizationId int
	OrganizationName string
	HasCustomFields bool
}

type AllDevices struct {
	Data []Device
	Meta Metadata
}