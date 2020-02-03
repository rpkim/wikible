package util

import (
	"bytes"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"text/template"
)

func ParseProjectConfig(projectFilePath string) string {
	yamlFile, err := ioutil.ReadFile(projectFilePath)
	config := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(yamlFile), &config)
	if err != nil {
		log.Fatalf("yaml Unmarshal error: %v", err)
		return ""
	}

	templateFilePath := fmt.Sprintf("%v", config["template"])

	//Testing
	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, config["vars"]); err != nil {
		log.Fatalln(err)
		return ""
	}

	// remove the empty lines
	result := regexp.MustCompile(`[\t\r\n]+`).ReplaceAllString(strings.TrimSpace(tpl.String()), "\n")

	return result
}

// A sample config
// config := map[string]interface{}{
// 	"project_name": "Bixby",
// 	"country": []string{
// 		"KR","US",
// 	},
// 	"environment" : []string {
// 		"DEV","STG",
// 	},
// 	"module": []string {
// 		"ASR", "NLU", "TTS",
// 	},
// 	"organizations": []string {
// 		"HQ Devops", "SRI-B Operations", "SRC-B TechOps & Dev", "SRPOL TechOps & Dev",
// 	},
// }
