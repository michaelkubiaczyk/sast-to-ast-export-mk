package internal

import "time"

type Args struct {
	URL,
	Username,
	Password,
	OutputPath,
	ProductName string
	Export                      []string
	ProjectsActiveSince         int
	IsDefaultProjectActiveSince bool
	Debug                       bool
	DBConnectionString,
	ProjectsIds,
	OutputFolder,
	InputFolder,
	TeamName string
	RunTime          time.Time
	QueryMappingFile string
	NestedTeams      bool
	ProjectNameMap   string
	SimIDVersion     int
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
