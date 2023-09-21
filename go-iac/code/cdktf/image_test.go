// START OMIT
package main

import (
	"testing"

	"cdk.tf/go/stack/generated/kreuzwerker/docker/container"
	"cdk.tf/go/stack/generated/kreuzwerker/docker/image"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

var stack = NewMyApplicationsAbstraction(cdktf.Testing_App(nil), "stack")
var synth = cdktf.Testing_Synth(stack)

func TestShouldContainContainer(t *testing.T) {
	assertion := cdktf.Testing_ToHaveResource(synth, container.Container_TfResourceType())

	if !*assertion {
		t.Error("Assertion Failed")
	}
}

// END OMIT
func TestShouldUseUbuntuImage(t *testing.T) {
	properties := map[string]interface{}{
		"name": "ubuntu:latest",
	}
	assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, image.Image_TfResourceType(), &properties)

	if !*assertion {
		t.Error("Assertion Failed")
	}
}
