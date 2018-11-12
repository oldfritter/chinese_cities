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

type City struct {
	Id         string `yaml:"id" json:"id"`
	ProvinceId string `yaml:"province_id" json:"province_id"`
	Name       string `yaml:"name" json:"name"`
}

var AllCities []City

func FindAllCities() *[]City {
	if len(AllCities) == 0 {
		InitCities()
	}
	return &AllCities
}

func InitCities() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	path_str, _ := filepath.Abs(path.Dir(filename) + "/database/cities.yml")
	content, err := ioutil.ReadFile(path_str)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(content, &AllCities)
}

func FindCityById(id string) (*City, error) {
	if len(AllCities) == 0 {
		InitCities()
	}
	for _, city := range AllCities {
		if city.Id == id {
			return &city, nil
		}
	}
	return &City{}, fmt.Errorf("No city can be found.")
}

func FindCitiesByProvinceId(province_id string) *[]City {
	if len(AllCities) == 0 {
		InitCities()
	}
	var cities []City
	for _, city := range AllCities {
		if city.ProvinceId == province_id {
			cities = append(cities, city)
		}
	}
	return &cities
}
