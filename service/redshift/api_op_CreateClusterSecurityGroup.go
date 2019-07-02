// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateClusterSecurityGroupMessage
type CreateClusterSecurityGroupInput struct {
	_ struct{} `type:"structure"`

	// The name for the security group. Amazon Redshift stores the value as a lowercase
	// string.
	//
	// Constraints:
	//
	//    * Must contain no more than 255 alphanumeric characters or hyphens.
	//
	//    * Must not be "Default".
	//
	//    * Must be unique for all security groups that are created by your AWS
	//    account.
	//
	// Example: examplesecuritygroup
	//
	// ClusterSecurityGroupName is a required field
	ClusterSecurityGroupName *string `type:"string" required:"true"`

	// A description for the security group.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// A list of tag instances.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateClusterSecurityGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterSecurityGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClusterSecurityGroupInput"}

	if s.ClusterSecurityGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterSecurityGroupName"))
	}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateClusterSecurityGroupResult
type CreateClusterSecurityGroupOutput struct {
	_ struct{} `type:"structure"`

	// Describes a security group.
	ClusterSecurityGroup *ClusterSecurityGroup `type:"structure"`
}

// String returns the string representation
func (s CreateClusterSecurityGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateClusterSecurityGroup = "CreateClusterSecurityGroup"

// CreateClusterSecurityGroupRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Creates a new Amazon Redshift security group. You use security groups to
// control access to non-VPC clusters.
//
// For information about managing security groups, go to Amazon Redshift Cluster
// Security Groups (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using CreateClusterSecurityGroupRequest.
//    req := client.CreateClusterSecurityGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateClusterSecurityGroup
func (c *Client) CreateClusterSecurityGroupRequest(input *CreateClusterSecurityGroupInput) CreateClusterSecurityGroupRequest {
	op := &aws.Operation{
		Name:       opCreateClusterSecurityGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterSecurityGroupInput{}
	}

	req := c.newRequest(op, input, &CreateClusterSecurityGroupOutput{})
	return CreateClusterSecurityGroupRequest{Request: req, Input: input, Copy: c.CreateClusterSecurityGroupRequest}
}

// CreateClusterSecurityGroupRequest is the request type for the
// CreateClusterSecurityGroup API operation.
type CreateClusterSecurityGroupRequest struct {
	*aws.Request
	Input *CreateClusterSecurityGroupInput
	Copy  func(*CreateClusterSecurityGroupInput) CreateClusterSecurityGroupRequest
}

// Send marshals and sends the CreateClusterSecurityGroup API request.
func (r CreateClusterSecurityGroupRequest) Send(ctx context.Context) (*CreateClusterSecurityGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClusterSecurityGroupResponse{
		CreateClusterSecurityGroupOutput: r.Request.Data.(*CreateClusterSecurityGroupOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClusterSecurityGroupResponse is the response type for the
// CreateClusterSecurityGroup API operation.
type CreateClusterSecurityGroupResponse struct {
	*CreateClusterSecurityGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateClusterSecurityGroup request.
func (r *CreateClusterSecurityGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}