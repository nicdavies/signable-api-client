package types

type CreateTeamResponse struct {
	Http                   int    `json:"http"`
	Message                string `json:"message"`
	Href                   string `json:"href"`
	TeamId                 string `json:"team_id"`
	TeamName               string `json:"team_name"`
	TeamPermissionOwn      bool   `json:"team_permission_own"`
	TeamPermissionUsers    bool   `json:"team_permission_users"`
	TeamPermissionBranding bool   `json:"team_permission_branding"`
	TeamPermissionApps     bool   `json:"team_permission_apps"`
	TeamPermissionSettings bool   `json:"team_permission_settings"`
	TeamPermissionCompany  bool   `json:"team_permission_company"`
	TeamCreated            string `json:"team_created"`
}

type ReadTeamResponse struct {
	Http                   int         `json:"http"`
	TeamId                 int         `json:"team_id"`
	TeamName               string      `json:"team_name"`
	TeamUsers              string      `json:"team_users"`
	TeamPermissionOwn      bool        `json:"team_permission_own"`
	TeamPermissionUsers    bool        `json:"team_permission_users"`
	TeamPermissionBranding bool        `json:"team_permission_branding"`
	TeamPermissionApps     bool        `json:"team_permission_apps"`
	TeamPermissionSettings bool        `json:"team_permission_settings"`
	TeamPermissionCompany  bool        `json:"team_permission_company"`
	TeamUpdated            interface{} `json:"team_updated"`
	TeamCreated            string      `json:"team_created"`
}

type ListTeamsResponse struct {
	Http       int    `json:"http"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	TotalTeams string `json:"total_teams"`
	Teams      []struct {
		TeamId                 string      `json:"team_id"`
		TeamName               string      `json:"team_name"`
		TeamUsers              string      `json:"team_users"`
		TeamPermissionOwn      bool        `json:"team_permission_own"`
		TeamPermissionUsers    bool        `json:"team_permission_users"`
		TeamPermissionBranding bool        `json:"team_permission_branding"`
		TeamPermissionApps     bool        `json:"team_permission_apps"`
		TeamPermissionSettings bool        `json:"team_permission_settings"`
		TeamPermissionCompany  bool        `json:"team_permission_company"`
		TeamUpdated            interface{} `json:"team_updated"`
		TeamCreated            string      `json:"team_created"`
	} `json:"teams"`
}

type UpdateTeamResponse struct {
	Http                   int    `json:"http"`
	Message                string `json:"message"`
	Href                   string `json:"href"`
	TeamId                 int    `json:"team_id"`
	TeamName               string `json:"team_name"`
	TeamUsers              string `json:"team_users"`
	TeamPermissionOwn      bool   `json:"team_permission_own"`
	TeamPermissionUsers    bool   `json:"team_permission_users"`
	TeamPermissionBranding bool   `json:"team_permission_branding"`
	TeamPermissionApps     bool   `json:"team_permission_apps"`
	TeamPermissionSettings bool   `json:"team_permission_settings"`
	TeamPermissionCompany  bool   `json:"team_permission_company"`
	TeamUpdated            string `json:"team_updated"`
}

type DeleteTeamResponse struct {
	Http        int    `json:"http"`
	Message     string `json:"message"`
	TeamId      int    `json:"team_id"`
	TeamName    string `json:"team_name"`
	TeamRemoved string `json:"team_removed"`
}
