// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package workspace

import (
	"fmt"
	"strings"

	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/workspace/buildconfig"
)

type Workspace struct {
	Id                  string                     `json:"id" validate:"required"`
	Name                string                     `json:"name" validate:"required"`
	State               WorkspaceState             `json:"state" validate:"required"`
	Image               string                     `json:"image" validate:"required"`
	User                string                     `json:"user" validate:"required"`
	BuildConfig         *buildconfig.BuildConfig   `json:"buildConfig,omitempty" validate:"optional"`
	Repository          *gitprovider.GitRepository `json:"repository" validate:"required"`
	EnvVars             map[string]string          `json:"envVars" validate:"required"`
	TargetId            string                     `json:"targetId" validate:"required"`
	ApiKey              string                     `json:"-"`
	Metadata            *WorkspaceMetadata         `json:"metadata,omitempty" validate:"optional"`
	GitProviderConfigId *string                    `json:"gitProviderConfigId,omitempty" validate:"optional"`
} // @name Workspace

type WorkspaceState string // @name WorkspaceState

const (
	WorkspaceStatePendingCreate       WorkspaceState = "pending-create"
	WorkspaceStateCreating            WorkspaceState = "creating"
	WorkspaceStatePendingStart        WorkspaceState = "pending-start"
	WorkspaceStateStarting            WorkspaceState = "starting"
	WorkspaceStateStarted             WorkspaceState = "started"
	WorkspaceStatePendingStop         WorkspaceState = "pending-stop"
	WorkspaceStateStopping            WorkspaceState = "stopping"
	WorkspaceStateStopped             WorkspaceState = "stopped"
	WorkspaceStatePendingRestart      WorkspaceState = "pending-restart"
	WorkspaceStateError               WorkspaceState = "error"
	WorkspaceStateUnresponsive        WorkspaceState = "unresponsive"
	WorkspaceStatePendingDelete       WorkspaceState = "pending-delete"
	WorkspaceStatePendingForcedDelete WorkspaceState = "pending-forced-delete"
	WorkspaceStateDeleting            WorkspaceState = "deleting"
)

func (w *Workspace) WorkspaceFolderName() string {
	if w.Repository != nil {
		return w.Repository.Name
	}
	return w.Name
}

type WorkspaceInfo struct {
	Name             string `json:"name" validate:"required"`
	Created          string `json:"created" validate:"required"`
	IsRunning        bool   `json:"isRunning" validate:"required"`
	ProviderMetadata string `json:"providerMetadata,omitempty" validate:"optional"`
	TargetId         string `json:"targetId" validate:"required"`
} // @name WorkspaceInfo

type WorkspaceMetadata struct {
	UpdatedAt string     `json:"updatedAt" validate:"required"`
	Uptime    uint64     `json:"uptime" validate:"required"`
	GitStatus *GitStatus `json:"gitStatus" validate:"required"`
} // @name WorkspaceMetadata

type GitStatus struct {
	CurrentBranch   string        `json:"currentBranch" validate:"required"`
	Files           []*FileStatus `json:"fileStatus" validate:"required"`
	BranchPublished bool          `json:"branchPublished" validate:"optional"`
	Ahead           int           `json:"ahead" validate:"optional"`
	Behind          int           `json:"behind" validate:"optional"`
} // @name GitStatus

type FileStatus struct {
	Name     string `json:"name" validate:"required"`
	Extra    string `json:"extra" validate:"required"`
	Staging  Status `json:"staging" validate:"required"`
	Worktree Status `json:"worktree" validate:"required"`
} // @name FileStatus

// Status status code of a file in the Worktree
type Status string // @name Status

const (
	Unmodified         Status = "Unmodified"
	Untracked          Status = "Untracked"
	Modified           Status = "Modified"
	Added              Status = "Added"
	Deleted            Status = "Deleted"
	Renamed            Status = "Renamed"
	Copied             Status = "Copied"
	UpdatedButUnmerged Status = "Updated but unmerged"
)

type WorkspaceEnvVarParams struct {
	ApiUrl        string
	ServerUrl     string
	ServerVersion string
	ClientId      string
}

func GetWorkspaceEnvVars(workspace *Workspace, params WorkspaceEnvVarParams, telemetryEnabled bool) map[string]string {
	envVars := map[string]string{
		"DAYTONA_TARGET_ID":                workspace.TargetId,
		"DAYTONA_WORKSPACE_ID":             workspace.Id,
		"DAYTONA_WORKSPACE_REPOSITORY_URL": workspace.Repository.Url,
		"DAYTONA_SERVER_API_KEY":           workspace.ApiKey,
		"DAYTONA_SERVER_VERSION":           params.ServerVersion,
		"DAYTONA_SERVER_URL":               params.ServerUrl,
		"DAYTONA_SERVER_API_URL":           params.ApiUrl,
		"DAYTONA_CLIENT_ID":                params.ClientId,
		// (HOME) will be replaced at runtime
		"DAYTONA_AGENT_LOG_FILE_PATH": "(HOME)/.daytona-agent.log",
	}

	if telemetryEnabled {
		envVars["DAYTONA_TELEMETRY_ENABLED"] = "true"
	}

	return envVars
}

func GetWorkspaceHostname(workspaceId string) string {
	// Replace special chars with hyphen to form valid hostname
	// String resulting in consecutive hyphens is also valid
	workspaceId = strings.ReplaceAll(workspaceId, "_", "-")
	workspaceId = strings.ReplaceAll(workspaceId, "*", "-")
	workspaceId = strings.ReplaceAll(workspaceId, ".", "-")

	hostname := fmt.Sprintf("ws-%s", workspaceId)

	if len(hostname) > 63 {
		return hostname[:63]
	}

	return hostname
}
