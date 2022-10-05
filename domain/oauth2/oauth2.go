package oauth2

import (
	"time"
)

type Client struct {
	ClientID int64  `json:"clientId"`
	Secret   string `json:"secret"`
	Name     string `json:"name"`
	WebSite  string `json:"webSite"`
	Email    string `json:"email"`
	Enabled  bool   `json:"enabled"`
	Paid     bool   `json:"paid"`
}

type ClientRedirectURI struct {
	ID       int64
	URI      string
	ClientID int64
}

type ClientAllowedURI struct {
	ID       int64
	URI      string
	ClientID int64
	//Super    bool
}

type ClientRole struct {
	ID       int64
	Role     string
	ClientID int64
	//Super    bool
}

type ClientScope struct {
	ID       int64
	Scope    string
	ClientID int64
}

type ClientRoleURI struct {
	ClientRoleID       int64
	ClientAllowedURIID int64
}

type RoleURI struct {
	ClientRoleID       int64
	Role               string
	ClientAllowedURIID int64
	ClientAllowedURI   string
	ClientID           int64
}

type RefreshToken struct {
	ID    int64
	Token string
}

type AccessToken struct {
	ID             int64
	Token          string
	Expires        time.Time
	RefreshTokenID int64
}


type AuthorizationCode struct {
	AuthorizationCode int64
	ClientID          int64
	UserID            string
	Expires           time.Time
	AccessTokenID     int64
	RandonAuthCode    string
	AlreadyUsed       bool
	Scope             string
}

type AuthCodeScope struct {
	ID                int64
	Scope             string
	AuthorizationCode int64
}

type AuthCodeRevolk struct {
	ID                int64
	AuthorizationCode int64
}

type ClientGrantType struct {
	ID        int64
	GrantType string
	ClientID  int64
}

type ImplicitGrant struct {
	ID            int64
	ClientID      int64
	UserID        string
	AccessTokenID int64
	Scope         string
}


type ImplicitScope struct {
	ID              int64
	Scope           string
	ImplicitGrantID int64
}

type PasswordGrant struct {
	ID            int64
	ClientID      int64
	UserID        string
	AccessTokenID int64
}

type CredentialsGrant struct {
	ID            int64
	ClientID      int64
	AccessTokenID int64
}