package file

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

//ReadFromYAML reads the YAML file and pass to the object
//args:
//	path: file path location
//	target: object which will hold the value
//returns:
//	error: operation state error
func ReadFromYAML(path string, target interface{}) error {
	yf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yf, target)
}

// PathExist check the path directory if exist
func PathExist(p string) bool {
	if stat, err := os.Stat(p); err == nil && stat.IsDir() {
		return true
	}
	return false
}
