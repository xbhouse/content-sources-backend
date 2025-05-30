package api

import "github.com/lib/pq"

type RepositoryRpm struct {
	UUID     string `json:"uuid"`     // Identifier of the rpm
	Name     string `json:"name"`     // The rpm package name
	Arch     string `json:"arch"`     // The architecture of the rpm
	Version  string `json:"version"`  // The version of the  rpm
	Release  string `json:"release"`  // The release of the rpm
	Epoch    int32  `json:"epoch"`    // The epoch of the rpm
	Summary  string `json:"summary"`  // The summary of the rpm
	Checksum string `json:"checksum"` // The checksum of the rpm
}

type SnapshotRpm struct {
	Name    string `json:"name"`    // The rpm package name
	Arch    string `json:"arch"`    // The architecture of the rpm
	Version string `json:"version"` // The version of the  rpm
	Release string `json:"release"` // The release of the rpm
	Epoch   string `json:"epoch"`   // The epoch of the rpm
	Summary string `json:"summary"` // The summary of the rpm
}

type RepositoryRpmCollectionResponse struct {
	Data  []RepositoryRpm  `json:"data"`  // List of rpms
	Meta  ResponseMetadata `json:"meta"`  // Metadata about the request
	Links Links            `json:"links"` // Links to other pages of results
}

type SnapshotErrataListRequest struct {
	UUID     string   `param:"uuid" validate:"required"` // Identifier of the repository
	Search   string   `query:"search"`                   // Errata Id to optionally filter-on
	Type     []string `query:"type"`                     // Type to optionally filter-on
	Severity []string `query:"severity"`                 // Severity to optionally filter-on
}

type SnapshotErrata struct {
	Id              string   `json:"id"`
	ErrataId        string   `json:"errata_id"`        // ID of the errata
	Title           string   `json:"title"`            // Title of the errata
	Summary         string   `json:"summary"`          // Summary of the errata
	Description     string   `json:"description"`      // Description of the errata
	IssuedDate      string   `json:"issued_date"`      // IssuedDate of the errata
	UpdateDate      string   `json:"updated_date"`     // UpdateDate of the errata
	Type            string   `json:"type"`             // Type of the errata
	Severity        string   `json:"severity"`         // Severity of the errata
	RebootSuggested bool     `json:"reboot_suggested"` // Whether a reboot is suggested
	CVEs            []string `json:"cves"`             // List of CVEs
}

type SnapshotErrataCollectionResponse struct {
	Data  []SnapshotErrata `json:"data"`  // List of errata
	Meta  ResponseMetadata `json:"meta"`  // Metadata about the request
	Links Links            `json:"links"` // Links to other pages of results
}
type SnapshotRpmCollectionResponse struct {
	Data  []SnapshotRpm    `json:"data"`  // List of rpms
	Meta  ResponseMetadata `json:"meta"`  // Metadata about the request
	Links Links            `json:"links"` // Links to other pages of results
}

type RepositoryRpmRequest struct {
	UUID   string `param:"uuid" validate:"required"` // Identifier of the repository
	Search string `query:"search"`                   // Search string based query to optionally filter-on
	SortBy string `query:"sort_by"`                  // SortBy sets the sort order of the result
}

type SearchRpmRequest struct {
	URLs   []string `json:"urls,omitempty"`  // URLs of repositories to search
	UUIDs  []string `json:"uuids,omitempty"` // List of repository UUIDs to search
	Search string   `json:"search"`          // Search string to search rpm names
	Limit  *int     `json:"limit,omitempty"` // Maximum number of records to return for the search
}

type SnapshotSearchRpmRequest struct {
	UUIDs                 []string `json:"uuids,omitempty"`                   // List of Snapshot UUIDs to search
	Search                string   `json:"search"`                            // Search string to search rpm names
	Limit                 *int     `json:"limit,omitempty"`                   // Maximum number of records to return for the search
	IncludePackageSources bool     `json:"include_package_sources,omitempty"` // Whether to include module information
}

type DetectRpmsRequest struct {
	URLs     []string `json:"urls,omitempty"`  // URLs of repositories to search
	UUIDs    []string `json:"uuids,omitempty"` // List of repository UUIDs to search
	RpmNames []string `json:"rpm_names"`       // List of rpm names to search
	Limit    *int     `json:"limit,omitempty"` // Maximum number of records to return for the search
}

const SearchRpmRequestLimitDefault int = 100
const SearchRpmRequestLimitMaximum int = 500

type SearchRpmResponse struct {
	PackageName    string                   `json:"package_name"`                       // Package name found
	Summary        string                   `json:"summary"`                            // Summary of the package found
	PackageSources []PackageSourcesResponse `json:"package_sources,omitempty" gorm:"-"` // List of the module streams for the package
}

type PackageSourcesResponse struct {
	Type         string         `json:"type"`                  // Type of rpm (can be either 'package' or 'module')
	Name         string         `json:"name,omitempty"`        // Name of the module
	Stream       string         `json:"stream,omitempty"`      // Stream of the module
	Context      string         `json:"context,omitempty"`     // Context of the module
	Arch         string         `json:"arch,omitempty"`        // Architecture of the module
	Version      string         `json:"version,omitempty"`     // Version of the module
	Description  string         `json:"description,omitempty"` // Description of the module
	PackageNames pq.StringArray `json:"-" gorm:"type:text"`    // Packages included in the module
	StartDate    string         `json:"start_date,omitempty"`  // Start date of the lifecycle
	EndDate      string         `json:"end_date,omitempty"`    // End date of the lifecycle
}

type DetectRpmsResponse struct {
	Found   []string `json:"found"`   // List of rpm names found in given repositories
	Missing []string `json:"missing"` // List of rpm names not found in given repositories
}

// SetMetadata Map metadata to the collection.
// meta Metadata about the request.
// links Links to other pages of results.
func (r *RepositoryRpmCollectionResponse) SetMetadata(meta ResponseMetadata, links Links) {
	r.Meta = meta
	r.Links = links
}

func (r *SnapshotErrataCollectionResponse) SetMetadata(meta ResponseMetadata, links Links) {
	r.Meta = meta
	r.Links = links
}

func (r *SnapshotRpmCollectionResponse) SetMetadata(meta ResponseMetadata, links Links) {
	r.Meta = meta
	r.Links = links
}
