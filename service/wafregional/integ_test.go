// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package wafregional_test

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

var _ aws.Config
var _ awserr.Error

func TestInteg_00_ListRules(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-east-1")
	svc := wafregional.New(cfg)
	params := &wafregional.ListRulesInput{
		Limit: aws.Int64(20),
	}

	req := svc.ListRulesRequest(params)
	req.SetContext(ctx)

	_, err := req.Send()
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_CreateSqlInjectionMatchSet(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-east-1")
	svc := wafregional.New(cfg)
	params := &wafregional.CreateSqlInjectionMatchSetInput{
		ChangeToken: aws.String("fake_token"),
		Name:        aws.String("fake_name"),
	}

	req := svc.CreateSqlInjectionMatchSetRequest(params)
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
