package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewCloudBackendStack(scope constructs.Construct, name string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &name)

	cdktf.NewCloudBackend(stack, &cdktf.CloudBackendConfig{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("company"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String("my-app-prod")),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("dns-server"), &cdktf.TerraformOutputConfig{
		Value: "hello-world",
	})

	return stack
}
