resource "aws_s3_bucket" "test" {
  bucket = "my-bucket-test"

  tags = {
    Name        = "My test bucket"
    Environment = "Dev"
  }
}