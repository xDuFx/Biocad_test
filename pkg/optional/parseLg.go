package optional

import (
	"os"
	"strings"
	"testB/pkg/repository"
)

func ReadConf(filename string) (connStr, fileDirectPdf, fileDirect string, err error) {
	conn, err := os.ReadFile(filename)
	if err != nil {
		return "", "", "", err
	}
	str := strings.Split(string(conn), "\r")
	connStr = str[0]
	fileDirectPdf = str[1][1:]
	fileDirect = str[2][1:]
	return connStr, fileDirectPdf, fileDirect, err
}

func NameSplit(path string) (cname string) {
	splitName := []byte(path)
	var pos int
	for i := len(splitName) - 1; ; i-- {
		if i == 0 || splitName[i] == '\\' || splitName[i] == '/' {
			break
		}
		pos = i
	}
	return string(splitName[pos:])
}

func CheckResol(path string) bool {
	name := NameSplit(path)
	splitName := []byte(name)
	var pos int
	for i := len(splitName) - 1; ; i-- {
		if splitName[i] == '.' {
			break
		}
		pos = i
	}
	return string(splitName[pos:]) == "tsv"
}

func RecordingError(filename string, err error, db *repository.PGRepo) {
	f, _ := os.OpenFile("errors.txt", os.O_WRONLY, 0666)
	f.WriteString(filename + ": ")
	f.WriteString("\n")
	db.Errors(filename )
}
