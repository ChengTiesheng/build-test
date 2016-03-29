package main

// PassedParams contains the parameters needed by docker build.
type PassedParams struct {
	ImageName     string
	Username      string
	Password      string
	Email         string
	Dockerfile    string
	DockerfileURL string
	TarURL        string
	TarFile       []byte
	GitUsr        string
	GitRepo       string
	GitTag        string
}

// PushAuth ...
type PushAuth struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Serveraddress string `json:"serveraddress"`
	Email         string `json:"email"`
}

// StreamCatcher ...
type StreamCatcher struct {
	ErrorDetail ErrorCatcher `json:"errorDetail"`
	Stream      string       `json:"stream"`
}

// ErrorCatcher ...
type ErrorCatcher struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// MessageStream ...
type MessageStream struct {
	Error       string `json:"error"`
	ErrorDetail struct {
		Message string `json:"message"`
	} `json:"errorDetail"`
	ID             string `json:"id"`
	Progress       string `json:"progress"`
	ProgressDetail struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"progressDetail"`
	Status string `json:"status"`
	Stream string `json:"stream"`
}
