package api

import (
	"os"
	"path/filepath"

	"github.com/content-services/content-sources-backend/pkg/config"
)

const IdentityHeader = "x-rh-identity"
const ApiVersion = "1.0"
const ApiVersionMajor = "1"

// CollectionMetadataSettable a collection response with settable metadata
type CollectionMetadataSettable interface {
	SetMetadata(meta ResponseMetadata, links Links)
}

type PaginationData struct {
	Limit  int    `query:"limit" json:"limit" `    // Number of results to return
	Offset int    `query:"offset" json:"offset"`   // Offset into the total results
	SortBy string `query:"sort_by" json:"sort_by"` // SortBy sets the sort order of the results
}

type FilterData struct {
	Search              string `query:"search" json:"search" `                              // Search string based query to optionally filter on
	Arch                string `query:"arch" json:"arch" `                                  // Comma separated list of architecture to optionally filter on (e.g. 'x86_64,s390x' would return Repositories with x86_64 or s390x only)
	Version             string `query:"version" json:"version"`                             // Comma separated list of versions to optionally filter on  (e.g. '7,8' would return Repositories with versions 7 or 8 only)
	AvailableForArch    string `query:"available_for_arch" json:"available_for_arch"`       // Filter by compatible arch (e.g. 'x86_64' would return Repositories with the 'x86_64' arch and Repositories where arch is not set)
	AvailableForVersion string `query:"available_for_version" json:"available_for_version"` // Filter by compatible version (e.g. 7 would return Repositories with the version 7 or where version is not set)
	Name                string `query:"name" json:"name"`                                   // Filter repositories by name using an exact match.
	URL                 string `query:"url" json:"url"`                                     // Comma separated list of urls to optionally filter on.
	UUID                string `query:"uuid" json:"uuid"`                                   // Comma separated list of uuids to optionally filter on.
	Status              string `query:"status" json:"status"`                               // Comma separated list of statuses to optionally filter on.
	Origin              string `query:"origin" json:"origin"`                               // Comma separated list of origins to filter on (e.g. external, red_hat, upload)
	ContentType         string `query:"content_type" json:"content_type"`                   // Filter repositories by content type (e.g. rpm)
}

type ResponseMetadata struct {
	Limit  int   `query:"limit" json:"limit"`   // Limit of results used for the request
	Offset int   `query:"offset" json:"offset"` // Offset into results used for the request
	Count  int64 `json:"count"`                 // Total count of results
}

type Links struct {
	First string `json:"first"`          // Path to first page of results
	Next  string `json:"next,omitempty"` // Path to next page of results
	Prev  string `json:"prev,omitempty"` // Path to previous page of results
	Last  string `json:"last"`           // Path to last page of results
}

type UUIDListRequest struct {
	UUIDs []string `json:"uuids" validate:"required"`
}

type AdminTaskFilterData struct {
	Status    string `json:"status"` // Comma separated list of statuses to optionally filter on.
	OrgId     string `json:"org_id"`
	AccountId string `json:"account_id"`
	Typename  string `json:"type"`
}

type FeatureStatus struct {
	OrgID       string   `json:"org_id"`
	FeatureList []string `json:"feature_list"`
}

func RootPrefix() string {
	pathPrefix, present := os.LookupEnv("PATH_PREFIX")
	if !present {
		pathPrefix = "api"
	}

	appName, present := os.LookupEnv("APP_NAME")
	if !present {
		appName = config.DefaultAppName
	}
	return filepath.Join("/", pathPrefix, appName)
}

func FullRootPath() string {
	return filepath.Join(RootPrefix(), "v"+ApiVersion)
}
func MajorRootPath() string {
	return filepath.Join(RootPrefix(), "v"+ApiVersionMajor)
}
