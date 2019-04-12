package userlib

import (
	"io/ioutil"
	"path"
)

type fileReader func(string, string)([]byte, error)

const (
	FILEERRORCODE = 404
	FILEERRORMSG = "File Read Error"
	TIMEOUTERRORCODE = 408
	SUCCESSCODE = 200
	CapacityString = "Cache status:  # of entries (%v)\ntotal bytes occupied by entries [%v]\nmax allowed capacity {%v}\n"
	TimeoutString = "The file request timed out!\n"
	CacheCloseMessage = "The cache has been cleared!\n"
	ContextType = "Content-Type"
)

var f fileReader = func(workingDir, filename string)(data []byte, err error){
	filepath := GetRealFilePath(workingDir, filename)
	data, err = ioutil.ReadFile(filepath)
	return
}

func ReadFile(workingDir, filename string)(data []byte, err error){
	data, err = f(workingDir, filename)
	return
}

func ReplaceReadFile(newfunc func(string, string)([]byte, error)){
	f = newfunc
}

func GetRealFilePath(workingDir, filename string) string {
	if filename[0:2] == "./" {
		filename = filename[2:]
	}
	if workingDir[len(workingDir) - 1:] != "/" {
		workingDir = workingDir + "/"
	}
	return workingDir + filename
}

func GetContentType(filename string) (string) {
	extension := path.Ext(filename)
	switch extension {
	case ".htm":
		fallthrough
	case ".html":
		return "text/html"
	case ".jpeg":
		fallthrough
	case ".jpg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".pdf":
		return "application/pdf"
	default:
		return "text/plain; charset=utf-8"
	}
}