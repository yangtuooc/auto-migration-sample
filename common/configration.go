package common

type Configuration struct {
	Mysql  MysqlConfiguration `json:"mysql" yaml:"mysql"`
	Server Server             `json:"server" yaml:"server"`
}

type Server struct {
	Port    string `json:"port" yaml:"port"`
	Context string `json:"context" yaml:"context"`
}

type MysqlConfiguration struct {
	Host     string `json:"host" yaml:"host"`
	Option   string `json:"option" yaml:"option"`
	Database string `json:"database" yaml:"database"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (s Server) UseContext() string {
	if s.Context == "" {
		return "/api"
	}
	return s.Context
}

func (s Server) UsePort() string {
	return ":" + s.Port
}
