package vsax

type Metadata struct {
	TotalCount    int
	NextQueryLink string
	ResponseCode  int
}

type Device struct {
	Identifier       string
	Name             string
	GroupId          int
	GroupName        string
	IsAgentInstalled bool
	IsMdmEnrolled    bool
	SiteId           int
	SiteName         string
	OrganizationId   int
	OrganizationName string
	HasCustomFields  bool
}

type DeviceResult struct {
	Data Device
	Meta Metadata
}

type AllDevicesResult struct {
	Data []Device
	Meta Metadata
}

type AssetSecurity struct {
	Type     string
	Name     string
	Enabled  bool
	UpToDate bool
}

type AssetLocalIpAddress struct {
	Name            string
	PhysicalAddress string
	DhcpEnabled     bool
	Gateways        []string
	DnsServers      []string
	SubnetMask      string
	IpV4            string
	IpV6            string
}

type AssetInstalledSoftware struct {
	Name      string
	Publisher string
	Version   string
}

type AssetDisk struct {
	Name           string
	System         bool
	FreePercentage int
	TotalValue     int
}

type AssetIpAddressIp struct {
	IP       string
	V6       bool
	Download int
	Upload   int
}

type AssetIpAddress struct {
	Name string
	MAC  string
	IPs  []AssetIpAddressIp
}

type AssetAssetInfo struct {
	CategoryName string
	CategoryData map[string]string
}

type AssetAvailableUpdates struct {
	UpdateId       string
	RevisionNumber int
	Title          string
	Description    string
	KbArticle      string
	Severity       string
	CvssScore      float32
	CveCodes       []string
	Category       string
	ReleaseDate    string
}

type AssetUpdates struct {
	Critical    int
	Important   int
	Unspecified int
}

type AssetEventLogs struct {
	Error       int
	Warning     int
	Information int
}

type Asset struct {
	Identifier        string
	Name              string
	GroupName         string
	Description       string
	Tags              []string
	Type              string
	ClientVersion     string
	LastSeenOnline    string
	ExternalUrl       string
	CpuUsage          int
	MemoryUsage       int
	MemoryTotal       int
	FirewallEnabled   bool
	AntivirusEnabled  string
	AntiVirusUpToDate string
	UacEnabled        bool
	Updates           AssetUpdates
	AvailableUpdates  AssetAvailableUpdates
	AssetInfo         []AssetAssetInfo
	PublicIpAddress   string
	ComputerId        int
	OrganizationId    int
	SiteId            int
	IpAddresses       []AssetIpAddress
	Disks             []AssetDisk
	InstalledSoftware []AssetInstalledSoftware
	LocalIpAddresses  []AssetLocalIpAddress
	Security          []AssetSecurity
}

type AllAssetsResult struct {
	Data []Asset
	Meta Metadata
}

type AssetResult struct {
	Data Asset
	Meta Metadata
}
