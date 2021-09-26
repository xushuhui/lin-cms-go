package conf

import "time"

type Config struct {
	Server Server
	Data   Data
	Log    Log
}
type Server struct {
	Http Http
}
type Data struct {
	Database Database
	Redis    Redis
}
type Http struct {
	Addr    string        `json:"addr"`
	Timeout time.Duration `json:"timeout"`
}
type Database struct {
	Driver string
	Source string
}
type Redis struct {
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}
type Log struct {
	Path string
}
