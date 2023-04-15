package altutils

import "path/filepath"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
func CheckExtension(path string, extension string) {
	if filepath.Ext(path) != extension {
		panic("Invalid file")
	}
}
