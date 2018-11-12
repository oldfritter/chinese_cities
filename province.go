package chinese_cities

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type Province struct {
	Id   int    `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
}

var AllProvinces []Province

func InitProvinces() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	path_str, _ := filepath.Abs(path.Dir(filename) + "/database/provinces.yml")
	content, err := ioutil.ReadFile(path_str)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(content, &AllProvinces)
}

func FindProvinceById(id int) (*Province, error) {
	if len(AllProvinces) == 0 {
		InitProvinces()
	}
	for _, province := range AllProvinces {
		if province.Id == id {
			return &province, nil
		}
	}
	return &Province{}, fmt.Errorf("No province can be found.")
}

func FindAllProvinces() *[]Province {
	if len(AllProvinces) == 0 {
		InitProvinces()
	}
	return &AllProvinces
}

func (province *Province) Cities() *[]City {
	return FindCitiesByProvinceId(province.Id)
}
