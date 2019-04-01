// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package codecommit_test

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

var _ aws.Config
var _ awserr.Error

func TestInteg_00_ListRepositories(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-west-2")
	svc := codecommit.New(cfg)
	params := &codecommit.ListRepositoriesInput{}

	req := svc.ListRepositoriesRequest(params)
	req.SetContext(ctx)

	_, err := req.Send()
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_ListBranches(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-west-2")
	svc := codecommit.New(cfg)
	params := &codecommit.ListBranchesInput{
		RepositoryName: aws.String("fake-repo"),
	}

	req := svc.ListBranchesRequest(params)
	req.SetContext(ctx)

	_, err := req.Send()
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if v := aerr.Code(); v == aws.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}
