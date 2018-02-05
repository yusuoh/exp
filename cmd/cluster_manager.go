package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/user"
	"path/filepath"
	// "reflect"
)

const (
	configPath = ".kube/config"
)

func ReadKubeConfig(path string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(path, configPath))
}

func ParseConfig(config []byte, s interface{}) error {
	err := yaml.Unmarshal(config, s)
	if err != nil {
		return err
	}
	return nil
}

type ClusterInfo struct {
	Name    string      `yaml:"name"`
	Cluster interface{} `yaml:cluster`
}

type KubeConfig struct {
	Clusters       []ClusterInfo `yaml:"clusters"`
	CurrentContext string        `yaml:"current-context"`
}

func ListCmd() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	config, err := ReadKubeConfig(usr.HomeDir)
	if err != nil {
		return err
	}
	c := KubeConfig{}
	err = ParseConfig(config, &c)
	if err != nil {
		return err
	}
	for _, m := range c.Clusters {
		fmt.Printf("Cluster: %s\n", m.Name)
	}
	fmt.Printf("Current context: %v\n", c.CurrentContext)
	return nil
}

func main() {
	err := ListCmd()
	if err != nil {
		fmt.Println(err.Error())
	}
}
