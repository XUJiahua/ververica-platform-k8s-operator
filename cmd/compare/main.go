package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
	nativeconverters "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/native_converters"
	"github.com/ghodss/yaml"
	"log"
	"os"
	"reflect"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	inputFile := os.Args[1]
	expectOutputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	handleErr(err)

	var vpDeployment v1beta2.VpDeployment
	err = yaml.Unmarshal(data, &vpDeployment)
	handleErr(err)

	deployment, err := nativeconverters.DeploymentFromNative(vpDeployment)
	handleErr(err)

	data, err = os.ReadFile(expectOutputFile)
	handleErr(err)

	var expectDep appmanagerapi.Deployment
	err = yaml.Unmarshal(data, &expectDep)
	handleErr(err)

	f1, err := os.Create("output.dump")
	handleErr(err)

	f2, err := os.Create("expect_output.dump")
	handleErr(err)
	if reflect.DeepEqual(deployment, expectDep) {
		log.Println("Success!")
	} else {
		spew.Fdump(f1, deployment)
		spew.Fdump(f2, expectDep)
		log.Println("Fail!")
	}

	//data, err = yaml.Marshal(deployment)
	//handleErr(err)
	//
	//err = os.WriteFile("output.yaml", data, 0644)
	//handleErr(err)
}
