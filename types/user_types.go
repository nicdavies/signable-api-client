package types

type CreateUserResponse struct {
	Http      int    `json:"http"`
	Message   string `json:"message"`
	Href      string `json:"href"`
	UserId    int    `json:"user_id"`
	TeamId    int    `json:"team_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	UserAdded string `json:"user_added"`
}

type ReadUserResponse struct {
	Http            int         `json:"http"`
	UserId          int         `json:"user_id"`
	TeamId          int         `json:"team_id"`
	UserName        string      `json:"user_name"`
	UserEmail       string      `json:"user_email"`
	UserAdded       string      `json:"user_added"`
	UserLastUpdated interface{} `json:"user_last_updated"`
}

type ListUsersResponse struct {
	Http       int    `json:"http"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	TotalUsers string `json:"total_users"`
	Next       string `json:"next"`
	Users      []struct {
		UserId          int    `json:"user_id"`
		TeamId          int    `json:"team_id"`
		UserName        string `json:"user_name"`
		UserEmail       string `json:"user_email"`
		UserAdded       string `json:"user_added"`
		UserLastUpdated string `json:"user_last_updated"`
	} `json:"users"`
}

type UpdateUserResponse struct {
	Http                int    `json:"http"`
	Message             string `json:"message"`
	Href                string `json:"href"`
	UserId              int    `json:"user_id"`
	TeamId              int    `json:"team_id"`
	UserName            string `json:"user_name"`
	UserEmail           string `json:"user_email"`
	UserPasswordChanged bool   `json:"user_password_changed"`
	UserUpdated         string `json:"user_updated"`
}

type DeleteUserResponse struct {
	Http      int    `json:"http"`
	Message   string `json:"message"`
	UserId    int    `json:"user_id"`
	TeamId    int    `json:"team_id"`
	UserEmail string `json:"user_email"`
	UserName  string `json:"user_name"`
}
