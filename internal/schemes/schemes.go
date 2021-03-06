package schemes

import (
	"net/http"
)

// UserStruct : user information auth implicated
type UserStruct struct {
	FullName string
	Username string
	Mail     string
	Power    int
	Password string
	Token    string
}

// VerboseStruct : verbosity params
type VerboseStruct struct {
	Verbose bool
	Level   int
}

// ServerAPIStruct : connect to api
type ServerAPIStruct struct {
	Server string
	Port   int
}

// TrackingStruct : logging events from client actions
type TrackingStruct struct {
	Path string
	Cli  string
}

// ConfigStruct : global config sections
type ConfigStruct struct {
	ServerAPI ServerAPIStruct
	Tracking  TrackingStruct
}

// UserSysStruct : user info
type UserSysStruct struct {
	Username string
	Uid      string
	GroupIds []string
	HomeDir  string
}

type ResponseStruct struct {
	Res   *http.Response
	Err   error
	Alert bool
}

type AuthStruct struct {
	Path  string
	Token string
}

var User = &UserStruct{}
var Verbose = &VerboseStruct{}
var ServerApi = &ServerAPIStruct{}
var Config = &ConfigStruct{}
var UserSys = &UserSysStruct{}
var Response = ResponseStruct{}
var Auth = &AuthStruct{}
