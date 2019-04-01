// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package applicationdiscoveryservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice"
)

var _ aws.Config
var _ awserr.Error

func TestInteg_00_DescribeAgents(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-west-2")
	svc := applicationdiscoveryservice.New(cfg)
	params := &applicationdiscoveryservice.DescribeAgentsInput{}

	req := svc.DescribeAgentsRequest(params)
	req.SetContext(ctx)

	_, err := req.Send()
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
