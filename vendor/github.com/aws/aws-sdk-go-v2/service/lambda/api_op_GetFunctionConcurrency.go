// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetFunctionConcurrencyInput struct {
	_ struct{} `type:"structure"`

	// The name of the Lambda function.
	//
	// Name formats
	//
	//    * Function name - my-function.
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//    * Partial ARN - 123456789012:function:my-function.
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFunctionConcurrencyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFunctionConcurrencyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFunctionConcurrencyInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFunctionConcurrencyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetFunctionConcurrencyOutput struct {
	_ struct{} `type:"structure"`

	// The number of simultaneous executions that are reserved for the function.
	ReservedConcurrentExecutions *int64 `type:"integer"`
}

// String returns the string representation
func (s GetFunctionConcurrencyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFunctionConcurrencyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ReservedConcurrentExecutions != nil {
		v := *s.ReservedConcurrentExecutions

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ReservedConcurrentExecutions", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opGetFunctionConcurrency = "GetFunctionConcurrency"

// GetFunctionConcurrencyRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns details about the reserved concurrency configuration for a function.
// To set a concurrency limit for a function, use PutFunctionConcurrency.
//
//    // Example sending a request using GetFunctionConcurrencyRequest.
//    req := client.GetFunctionConcurrencyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetFunctionConcurrency
func (c *Client) GetFunctionConcurrencyRequest(input *GetFunctionConcurrencyInput) GetFunctionConcurrencyRequest {
	op := &aws.Operation{
		Name:       opGetFunctionConcurrency,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-09-30/functions/{FunctionName}/concurrency",
	}

	if input == nil {
		input = &GetFunctionConcurrencyInput{}
	}

	req := c.newRequest(op, input, &GetFunctionConcurrencyOutput{})
	return GetFunctionConcurrencyRequest{Request: req, Input: input, Copy: c.GetFunctionConcurrencyRequest}
}

// GetFunctionConcurrencyRequest is the request type for the
// GetFunctionConcurrency API operation.
type GetFunctionConcurrencyRequest struct {
	*aws.Request
	Input *GetFunctionConcurrencyInput
	Copy  func(*GetFunctionConcurrencyInput) GetFunctionConcurrencyRequest
}

// Send marshals and sends the GetFunctionConcurrency API request.
func (r GetFunctionConcurrencyRequest) Send(ctx context.Context) (*GetFunctionConcurrencyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFunctionConcurrencyResponse{
		GetFunctionConcurrencyOutput: r.Request.Data.(*GetFunctionConcurrencyOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFunctionConcurrencyResponse is the response type for the
// GetFunctionConcurrency API operation.
type GetFunctionConcurrencyResponse struct {
	*GetFunctionConcurrencyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFunctionConcurrency request.
func (r *GetFunctionConcurrencyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
