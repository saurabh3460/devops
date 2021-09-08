package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type User struct {
	Name string
	Occupation string
}

func main() {

	yfile, err := ioutil.ReadFile("example.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]User)
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
	}

	word := [5]string{"Falcon", "sky", "cloud", "fox"}

	d, err := yaml.Marshal(&word)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("generatedYaml.yaml", d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
