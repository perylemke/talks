package main

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v4/go/pulumi"
	"github.com/stretchr/testify/assert"
)

func TestCreateS3Bucket(t *testing.T) {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := createS3Bucket(ctx, "my-test-bucket")
		if err != nil {
			t.Fatalf("Error on created the S3 bucket: %v", err)
		}

		assert.NotNil(t, bucket)
		assert.Equal(t, "my-test-bucket", bucket.ID())

		return nil
	})
}
