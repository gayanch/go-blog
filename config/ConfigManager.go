package config

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

const (
	dbConfigFile = "config/dbconfig.xml"
	blogConfigFile = "config/blogconfig.xml"
)

//DB Config
type XMLConfig struct {
	XMLName xml.Name `xml:"config"`
	Key string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

type XMLDbConfigs struct {
	XMLName xml.Name `xml:"db"`
	Configs []XMLConfig `xml:"config"`
}
//End DB Config

//Blog Config
type XMLBlogConfigs struct {
	XMLName xml.Name `xml:"blog"`
	Configs []XMLConfig `xml:"config"`	
}
//End Blog Config

func ReadDbConfig() (XMLDbConfigs, error){
	file, err := os.Open(dbConfigFile)
	if err != nil {
		return XMLDbConfigs{}, err
	}
	defer file.Close()
	
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return XMLDbConfigs{}, err
	}
	
	var config XMLDbConfigs
	err = xml.Unmarshal(data, &config)
	if err != nil {
		return XMLDbConfigs{}, err
	}
	
	return config, nil
}

func ReadBlogConfig() (XMLBlogConfigs, error) {
	file, err := os.Open(blogConfigFile)
	if err != nil {
		return XMLBlogConfigs{}, err
	}
	defer file.Close()
	
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return XMLBlogConfigs{}, err
	}
	
	var config XMLBlogConfigs
	err = xml.Unmarshal(data, &config)
	if err != nil {
		return XMLBlogConfigs{}, err
	}
	
	return config, nil
}
