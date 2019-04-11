package userlib

import (
	"io/ioutil"
	"path"
)

type fileReader func(string)([]byte, error)

const (
	FILEERRORCODE = 404
	FILEERRORMSG = "File Read Error"
	TIMEOUTERRORCODE = 408
	SUCCESSCODE = 200
	CapacityString = "Cache status:  # of entries %v\ntotal bytes occupied by entries %v\nmax allowed capacity %v\n"
	TimeoutString = "The file request timed out!"
	CacheCloseMessage = "The cache has been cleared!"
)

var F fileReader = func(filename string)(data []byte, err error){
	data, err = ioutil.ReadFile(filename)
	return
}

func ReadFile(filename string)(data []byte, err error){
	data, err = F(filename)
	return
}


func ReplaceReadFile(newfunc func(string)([]byte, error)){
	F = newfunc
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