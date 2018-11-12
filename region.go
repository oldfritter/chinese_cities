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

type Region struct {
	Id     string `yaml:"id" json:"id"`
	CityId string `yaml:"city_id" json:"city_id"`
	Name   string `yaml:"name" json:"name"`
}

var AllRegions []Region

func InitRegions() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	path_str, _ := filepath.Abs(path.Dir(filename) + "/database/regions.yml")
	content, err := ioutil.ReadFile(path_str)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(content, &AllRegions)
}

func FindRegionById(id string) (*Region, error) {
	if len(AllProvinces) == 0 {
		InitProvinces()
	}
	for _, region := range AllRegions {
		if region.Id == id {
			return &region, nil
		}
	}
	return &Region{}, fmt.Errorf("No region can be found.")
}

func (region *Region) City() *City {
	city, err := FindCityById(region.CityId)
	if err != nil {
		panic(err)
	}
	return city
}

func (region *Region) CityName() string {
	city := region.City()
	return (*city).Name
}

func (region *Region) Province() *Province {
	city := region.City()
	province, err := FindProvinceById(city.ProvinceId)
	if err != nil {
		panic(err)
	}
	return province
}

func (region *Region) ProvinceName() string {
	province := region.Province()
	return (*province).Name
}

func FindCitiesByCityId(city_id string) *[]Region {
	if len(AllRegions) == 0 {
		InitRegions()
	}
	var regions []Region
	for _, region := range AllRegions {
		if region.CityId == city_id {
			regions = append(regions, region)
		}
	}
	return &regions
}
