package folder

import (
	"time"

	"github.com/grafana/grafana/pkg/util/errutil"
)

var ErrMaximumDepthReached = errutil.NewBase(errutil.StatusBadRequest, "folder.maximum-depth-reached", errutil.WithPublicMessage("Maximum nested folder depth reached"))
var ErrBadRequest = errutil.NewBase(errutil.StatusBadRequest, "folder.bad-request")
var ErrDatabaseError = errutil.NewBase(errutil.StatusInternal, "folder.database-error")
var ErrInternal = errutil.NewBase(errutil.StatusInternal, "folder.internal")

const (
	GeneralFolderUID     = "general"
	RootFolderUID        = ""
	MaxNestedFolderDepth = 8
)

var ErrFolderNotFound = errutil.NewBase(errutil.StatusNotFound, "folder.notFound")

type Folder struct {
	ID          int64  `xorm:"pk autoincr 'id'"`
	OrgID       int64  `xorm:"org_id"`
	UID         string `xorm:"uid"`
	ParentUID   string `xorm:"parent_uid"`
	Title       string
	Description string

	Created time.Time
	Updated time.Time

	// TODO: validate if this field is required/relevant to folders.
	// currently there is no such column
	// UpdatedBy int64
}

// NewFolder tales a title and returns a Folder with the Created and Updated
// fields set to the current time.
func NewFolder(title string, description string) *Folder {
	return &Folder{
		Title:       title,
		Description: description,
		Created:     time.Now(),
		Updated:     time.Now(),
	}
}

// CreateFolderCommand captures the information required by the folder service
// to create a folder.
type CreateFolderCommand struct {
	UID         string `json:"uid"`
	OrgID       int64  `json:"orgId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ParentUID   string `json:"parent_uid"`
}

// UpdateFolderCommand captures the information required by the folder service
// to update a folder. Use Move to update a folder's parent folder.
type UpdateFolderCommand struct {
	Folder         *Folder `json:"folder"` // The extant folder
	NewUID         *string `json:"uid" xorm:"uid"`
	NewParentUID   *string `json:"parent_uid" xorm:"parent_uid"`
	NewTitle       *string `json:"title"`
	NewDescription *string `json:"description"`
}

// MoveFolderCommand captures the information required by the folder service
// to move a folder.
type MoveFolderCommand struct {
	UID          string `json:"uid"`
	NewParentUID string `json:"new_parent_uid"`
	OrgID        int64  `json:"orgId"`
}

// DeleteFolderCommand captures the information required by the folder service
// to delete a folder.
type DeleteFolderCommand struct {
	UID   string `json:"uid" xorm:"uid"`
	OrgID int64  `json:"orgId" xorm:"org_id"`
}

// GetFolderQuery is used for all folder Get requests. Only one of UID, ID, or
// Title should be set; if multilpe fields are set by the caller the dashboard
// service will select the field with the most specificity, in order: ID, UID,
// Title.
type GetFolderQuery struct {
	UID   *string
	ID    *int64
	Title *string
	OrgID int64
}

// GetParentsQuery captures the information required by the folder service to
// return a list of all parent folders of a given folder.
type GetParentsQuery struct {
	UID   string `xorm:"uid"`
	OrgID int64  `xorm:"org_id"`
}

// GetTreeCommand captures the information required by the folder service to
// return a list of child folders of the given folder.

type GetTreeQuery struct {
	UID   string `xorm:"uid"`
	OrgID int64  `xorm:"org_id"`
	Depth int64

	// Pagination options
	Limit int64
	Page  int64
}
