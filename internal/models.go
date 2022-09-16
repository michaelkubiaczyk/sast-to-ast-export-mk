package internal

type Args struct {
	URL,
	Username,
	Password,
	OutputPath,
	ProductName string
	Export              []string
	ProjectsActiveSince int
	Debug               bool
	DBConnectionString,
	ProjectsIds,
	TeamName string
	QueryMappingFile string
}

type ReportJob struct {
	ProjectID  int
	ScanID     int
	ReportType string
}

type TriagedScan struct {
	ProjectID int
	ScanID    int
}

type PresetJob struct {
	PresetID int
}

type PresetConsumeOutput struct {
	Err      error
	PresetID int
}
