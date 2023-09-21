package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// START OMIT
func TestTerraformAwsS3Example(t *testing.T) {
	t.Parallel()

	expectedName := fmt.Sprintf("my-bucket-test-%s", strings.ToLower(random.UniqueId()))
	expectedEnvironment := "dev"

	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",

		Vars: map[string]interface{}{
			"tag_bucket_name":        expectedName,
			"tag_bucket_environment": expectedEnvironment,
			"with_policy":            "true",
			"region":                 awsRegion,
		},
	})
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketID := terraform.Output(t, terraformOptions, "bucket_id")

	actualStatus := aws.GetS3BucketVersioning(t, awsRegion, bucketID)
	expectedStatus := "Enabled"
	assert.Equal(t, expectedStatus, actualStatus)
}

// END OMIT
