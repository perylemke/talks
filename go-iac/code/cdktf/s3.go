package main

import (
	"github.com/hashicorp/terraform-cdk-go/tfcdk"
	"github.com/zclconf/go-cty/cty"
)

func NewS3Bucket(stack tfcdk.TfcdkStack, name string) {
	stack.Resource(name, "aws_s3_bucket", "my_test_bucket",
		&map[string]cty.Value{
			"bucket": cty.StringVal("my-bucket-test"),
		})
}

func main() {
	app := tfcdk.NewApp(nil)

	stack := tfcdk.NewStack(app, "MyStack")

	NewS3Bucket(stack, "MyS3Bucket")

	app.Synth(nil)
}
