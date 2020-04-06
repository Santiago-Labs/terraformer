// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateIdentityProviderInput struct {
	_ struct{} `type:"structure"`

	// The identity provider attribute mapping to be changed.
	AttributeMapping map[string]string `type:"map"`

	// A list of identity provider identifiers.
	IdpIdentifiers []string `type:"list"`

	// The identity provider details to be updated, such as MetadataURL and MetadataFile.
	ProviderDetails map[string]string `type:"map"`

	// The identity provider name.
	//
	// ProviderName is a required field
	ProviderName *string `min:"1" type:"string" required:"true"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateIdentityProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateIdentityProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateIdentityProviderInput"}

	if s.ProviderName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProviderName"))
	}
	if s.ProviderName != nil && len(*s.ProviderName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProviderName", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateIdentityProviderOutput struct {
	_ struct{} `type:"structure"`

	// The identity provider object.
	//
	// IdentityProvider is a required field
	IdentityProvider *IdentityProviderType `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateIdentityProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateIdentityProvider = "UpdateIdentityProvider"

// UpdateIdentityProviderRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Updates identity provider information for a user pool.
//
//    // Example sending a request using UpdateIdentityProviderRequest.
//    req := client.UpdateIdentityProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/UpdateIdentityProvider
func (c *Client) UpdateIdentityProviderRequest(input *UpdateIdentityProviderInput) UpdateIdentityProviderRequest {
	op := &aws.Operation{
		Name:       opUpdateIdentityProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateIdentityProviderInput{}
	}

	req := c.newRequest(op, input, &UpdateIdentityProviderOutput{})
	return UpdateIdentityProviderRequest{Request: req, Input: input, Copy: c.UpdateIdentityProviderRequest}
}

// UpdateIdentityProviderRequest is the request type for the
// UpdateIdentityProvider API operation.
type UpdateIdentityProviderRequest struct {
	*aws.Request
	Input *UpdateIdentityProviderInput
	Copy  func(*UpdateIdentityProviderInput) UpdateIdentityProviderRequest
}

// Send marshals and sends the UpdateIdentityProvider API request.
func (r UpdateIdentityProviderRequest) Send(ctx context.Context) (*UpdateIdentityProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateIdentityProviderResponse{
		UpdateIdentityProviderOutput: r.Request.Data.(*UpdateIdentityProviderOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateIdentityProviderResponse is the response type for the
// UpdateIdentityProvider API operation.
type UpdateIdentityProviderResponse struct {
	*UpdateIdentityProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateIdentityProvider request.
func (r *UpdateIdentityProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
