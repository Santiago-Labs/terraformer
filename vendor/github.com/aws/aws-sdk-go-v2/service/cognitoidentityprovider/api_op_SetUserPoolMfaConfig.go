// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SetUserPoolMfaConfigInput struct {
	_ struct{} `type:"structure"`

	// The MFA configuration. Valid values include:
	//
	//    * OFF MFA will not be used for any users.
	//
	//    * ON MFA is required for all users to sign in.
	//
	//    * OPTIONAL MFA will be required only for individual users who have an
	//    MFA factor enabled.
	MfaConfiguration UserPoolMfaType `type:"string" enum:"true"`

	// The SMS text message MFA configuration.
	SmsMfaConfiguration *SmsMfaConfigType `type:"structure"`

	// The software token MFA configuration.
	SoftwareTokenMfaConfiguration *SoftwareTokenMfaConfigType `type:"structure"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SetUserPoolMfaConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetUserPoolMfaConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetUserPoolMfaConfigInput"}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}
	if s.SmsMfaConfiguration != nil {
		if err := s.SmsMfaConfiguration.Validate(); err != nil {
			invalidParams.AddNested("SmsMfaConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetUserPoolMfaConfigOutput struct {
	_ struct{} `type:"structure"`

	// The MFA configuration. Valid values include:
	//
	//    * OFF MFA will not be used for any users.
	//
	//    * ON MFA is required for all users to sign in.
	//
	//    * OPTIONAL MFA will be required only for individual users who have an
	//    MFA factor enabled.
	MfaConfiguration UserPoolMfaType `type:"string" enum:"true"`

	// The SMS text message MFA configuration.
	SmsMfaConfiguration *SmsMfaConfigType `type:"structure"`

	// The software token MFA configuration.
	SoftwareTokenMfaConfiguration *SoftwareTokenMfaConfigType `type:"structure"`
}

// String returns the string representation
func (s SetUserPoolMfaConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetUserPoolMfaConfig = "SetUserPoolMfaConfig"

// SetUserPoolMfaConfigRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Set the user pool multi-factor authentication (MFA) configuration.
//
//    // Example sending a request using SetUserPoolMfaConfigRequest.
//    req := client.SetUserPoolMfaConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/SetUserPoolMfaConfig
func (c *Client) SetUserPoolMfaConfigRequest(input *SetUserPoolMfaConfigInput) SetUserPoolMfaConfigRequest {
	op := &aws.Operation{
		Name:       opSetUserPoolMfaConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetUserPoolMfaConfigInput{}
	}

	req := c.newRequest(op, input, &SetUserPoolMfaConfigOutput{})
	return SetUserPoolMfaConfigRequest{Request: req, Input: input, Copy: c.SetUserPoolMfaConfigRequest}
}

// SetUserPoolMfaConfigRequest is the request type for the
// SetUserPoolMfaConfig API operation.
type SetUserPoolMfaConfigRequest struct {
	*aws.Request
	Input *SetUserPoolMfaConfigInput
	Copy  func(*SetUserPoolMfaConfigInput) SetUserPoolMfaConfigRequest
}

// Send marshals and sends the SetUserPoolMfaConfig API request.
func (r SetUserPoolMfaConfigRequest) Send(ctx context.Context) (*SetUserPoolMfaConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetUserPoolMfaConfigResponse{
		SetUserPoolMfaConfigOutput: r.Request.Data.(*SetUserPoolMfaConfigOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetUserPoolMfaConfigResponse is the response type for the
// SetUserPoolMfaConfig API operation.
type SetUserPoolMfaConfigResponse struct {
	*SetUserPoolMfaConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetUserPoolMfaConfig request.
func (r *SetUserPoolMfaConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
