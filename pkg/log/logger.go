package log

import log "github.com/grestful/logs"

var SqlLogger log.Logger

func SetSqlLogger(path string) {
	log.SetFile(log.FileConfig{

		Category: "file",
		Level:    "DEBUG",
		Filename: path + "/sql.log",
		Rotate:   true,
		Daily:    true,
	})
	writer := log.GetLogger("file")
	log.SetDefaultLog(writer)
}
