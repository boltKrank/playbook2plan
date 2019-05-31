package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// PlayBook parent struct
type PlayBook struct {
	Name   string
	Hosts  []string
	Become bool
	//Tasks  []Task
}

func main() {

	// Input file:

	dat, err := ioutil.ReadFile("./playbookExamples/nginx.yaml")
	check(err)
	fmt.Print(string(dat))

	var playbook PlayBook
	err = yaml.Unmarshal(dat, &playbook)
	check(err)

	playbookFile, err := os.Open("./playbookExamples/nginx.yaml")
	check(err)

	// Ansible example Playbooks: https://github.com/ansible/ansible-examples

	// Check https://github.com/go-yaml/yaml
	// https://godoc.org/gopkg.in/yaml.v2
	// YAML Syntax: https://docs.ansible.com/ansible/latest/reference_appendices/YAMLSyntax.html
	// Other parsers: http://sweetohm.net/article/go-yaml-parsers.en.html
	// Ansible example: https://www.tastycidr.net/yaml-in-go-parsing-multi-level-yaml-using-the-ghodss-yaml-library/

	// Ansible best practices: https://docs.ansible.com/ansible/latest/user_guide/playbooks_best_practices.html

	//Seperating structs into another file: https://stackoverflow.com/questions/14155122/how-to-call-function-from-another-file-in-go-language

	// Iterating through a struct https://stackoverflow.com/questions/18926303/iterate-through-the-fields-of-a-struct-in-go

	// parse through struct playbook

	playbookFile.Close()

	// Output file:

	// inputPtr := flag.String("input", "noinput", "a string")
	// outputPtr := flag.String("output", "nooutput", "a string")

	// fmt.Println("input:", *inputPtr)
	// fmt.Println("output:", *outputPtr)

	// if *inputPtr == "noinput" || *outputPtr == "nooutput" {

	// 	fmt.Println("Usage: playbook2plan -input=<input_playbook> -output=<output_plan>")

	// }

}
