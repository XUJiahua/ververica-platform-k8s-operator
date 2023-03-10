package main

import (
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	nativeconverters "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/native_converters"
	"github.com/ghodss/yaml"
	"log"
	"os"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	inputFile := os.Args[1]
	data, err := os.ReadFile(inputFile)
	handleErr(err)

	var vpDeployment v1beta2.VpDeployment
	err = yaml.Unmarshal(data, &vpDeployment)
	handleErr(err)

	deployment, err := nativeconverters.DeploymentFromNative(vpDeployment)
	handleErr(err)

	data, err = yaml.Marshal(deployment)
	handleErr(err)

	err = os.WriteFile("output.yaml", data, 0644)
	handleErr(err)
}
