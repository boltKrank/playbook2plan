package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Input file:

	dat, err := ioutil.ReadFile("./playbookExamples/nginx.yaml")
	check(err)
	fmt.Print(string(dat))

	playbookFile, err := os.Open("./playbookExamples/nginx.yaml")
	check(err)

	// Check https://github.com/go-yaml/yaml
	// https://godoc.org/gopkg.in/yaml.v2
	// YAML Syntax: https://docs.ansible.com/ansible/latest/reference_appendices/YAMLSyntax.html
	// Other parsers: http://sweetohm.net/article/go-yaml-parsers.en.html
	// Ansible example: https://www.tastycidr.net/yaml-in-go-parsing-multi-level-yaml-using-the-ghodss-yaml-library/

	// Ansible best practices: https://docs.ansible.com/ansible/latest/user_guide/playbooks_best_practices.html

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
