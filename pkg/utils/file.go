package utils

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func CreateFile(fileName string) (f *os.File, e error) {
	exist, e := PathExists(fileName)
	if e != nil {
		return
	}
	if exist == false {
		f, e := os.Create(fileName)
		if e != nil {
			return f, e
		}
	}
	f, e = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if e != nil {
		return
	}
	return
}
func LogDir(savePath string) string {
	return savePath + "/" + GetCurrentDate() + "/"
}
