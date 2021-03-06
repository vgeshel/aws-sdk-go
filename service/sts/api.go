package sts

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// AssumeRoleRequest generates a request for the AssumeRole operation.
func (c *STS) AssumeRoleRequest(input *AssumeRoleInput) (req *aws.Request, output *AssumeRoleOutput) {
	if opAssumeRole == nil {
		opAssumeRole = &aws.Operation{
			Name:       "AssumeRole",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opAssumeRole, input, output)
	output = &AssumeRoleOutput{}
	req.Data = output
	return
}

func (c *STS) AssumeRole(input *AssumeRoleInput) (output *AssumeRoleOutput, err error) {
	req, out := c.AssumeRoleRequest(input)
	output = out
	err = req.Send()
	return
}

var opAssumeRole *aws.Operation

// AssumeRoleWithSAMLRequest generates a request for the AssumeRoleWithSAML operation.
func (c *STS) AssumeRoleWithSAMLRequest(input *AssumeRoleWithSAMLInput) (req *aws.Request, output *AssumeRoleWithSAMLOutput) {
	if opAssumeRoleWithSAML == nil {
		opAssumeRoleWithSAML = &aws.Operation{
			Name:       "AssumeRoleWithSAML",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opAssumeRoleWithSAML, input, output)
	output = &AssumeRoleWithSAMLOutput{}
	req.Data = output
	return
}

func (c *STS) AssumeRoleWithSAML(input *AssumeRoleWithSAMLInput) (output *AssumeRoleWithSAMLOutput, err error) {
	req, out := c.AssumeRoleWithSAMLRequest(input)
	output = out
	err = req.Send()
	return
}

var opAssumeRoleWithSAML *aws.Operation

// AssumeRoleWithWebIdentityRequest generates a request for the AssumeRoleWithWebIdentity operation.
func (c *STS) AssumeRoleWithWebIdentityRequest(input *AssumeRoleWithWebIdentityInput) (req *aws.Request, output *AssumeRoleWithWebIdentityOutput) {
	if opAssumeRoleWithWebIdentity == nil {
		opAssumeRoleWithWebIdentity = &aws.Operation{
			Name:       "AssumeRoleWithWebIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opAssumeRoleWithWebIdentity, input, output)
	output = &AssumeRoleWithWebIdentityOutput{}
	req.Data = output
	return
}

func (c *STS) AssumeRoleWithWebIdentity(input *AssumeRoleWithWebIdentityInput) (output *AssumeRoleWithWebIdentityOutput, err error) {
	req, out := c.AssumeRoleWithWebIdentityRequest(input)
	output = out
	err = req.Send()
	return
}

var opAssumeRoleWithWebIdentity *aws.Operation

// DecodeAuthorizationMessageRequest generates a request for the DecodeAuthorizationMessage operation.
func (c *STS) DecodeAuthorizationMessageRequest(input *DecodeAuthorizationMessageInput) (req *aws.Request, output *DecodeAuthorizationMessageOutput) {
	if opDecodeAuthorizationMessage == nil {
		opDecodeAuthorizationMessage = &aws.Operation{
			Name:       "DecodeAuthorizationMessage",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDecodeAuthorizationMessage, input, output)
	output = &DecodeAuthorizationMessageOutput{}
	req.Data = output
	return
}

func (c *STS) DecodeAuthorizationMessage(input *DecodeAuthorizationMessageInput) (output *DecodeAuthorizationMessageOutput, err error) {
	req, out := c.DecodeAuthorizationMessageRequest(input)
	output = out
	err = req.Send()
	return
}

var opDecodeAuthorizationMessage *aws.Operation

// GetFederationTokenRequest generates a request for the GetFederationToken operation.
func (c *STS) GetFederationTokenRequest(input *GetFederationTokenInput) (req *aws.Request, output *GetFederationTokenOutput) {
	if opGetFederationToken == nil {
		opGetFederationToken = &aws.Operation{
			Name:       "GetFederationToken",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetFederationToken, input, output)
	output = &GetFederationTokenOutput{}
	req.Data = output
	return
}

func (c *STS) GetFederationToken(input *GetFederationTokenInput) (output *GetFederationTokenOutput, err error) {
	req, out := c.GetFederationTokenRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetFederationToken *aws.Operation

// GetSessionTokenRequest generates a request for the GetSessionToken operation.
func (c *STS) GetSessionTokenRequest(input *GetSessionTokenInput) (req *aws.Request, output *GetSessionTokenOutput) {
	if opGetSessionToken == nil {
		opGetSessionToken = &aws.Operation{
			Name:       "GetSessionToken",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetSessionToken, input, output)
	output = &GetSessionTokenOutput{}
	req.Data = output
	return
}

func (c *STS) GetSessionToken(input *GetSessionTokenInput) (output *GetSessionTokenOutput, err error) {
	req, out := c.GetSessionTokenRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetSessionToken *aws.Operation

type AssumeRoleInput struct {
	DurationSeconds *int64  `type:"integer"`
	ExternalID      *string `locationName:"ExternalId" type:"string"`
	Policy          *string `type:"string"`
	RoleARN         *string `locationName:"RoleArn" type:"string" required:"true"`
	RoleSessionName *string `type:"string" required:"true"`
	SerialNumber    *string `type:"string"`
	TokenCode       *string `type:"string"`

	metadataAssumeRoleInput `json:"-", xml:"-"`
}

type metadataAssumeRoleInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type AssumeRoleOutput struct {
	AssumedRoleUser  *AssumedRoleUser `type:"structure"`
	Credentials      *Credentials     `type:"structure"`
	PackedPolicySize *int64           `type:"integer"`

	metadataAssumeRoleOutput `json:"-", xml:"-"`
}

type metadataAssumeRoleOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type AssumeRoleWithSAMLInput struct {
	DurationSeconds *int64  `type:"integer"`
	Policy          *string `type:"string"`
	PrincipalARN    *string `locationName:"PrincipalArn" type:"string" required:"true"`
	RoleARN         *string `locationName:"RoleArn" type:"string" required:"true"`
	SAMLAssertion   *string `type:"string" required:"true"`

	metadataAssumeRoleWithSAMLInput `json:"-", xml:"-"`
}

type metadataAssumeRoleWithSAMLInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type AssumeRoleWithSAMLOutput struct {
	AssumedRoleUser  *AssumedRoleUser `type:"structure"`
	Audience         *string          `type:"string"`
	Credentials      *Credentials     `type:"structure"`
	Issuer           *string          `type:"string"`
	NameQualifier    *string          `type:"string"`
	PackedPolicySize *int64           `type:"integer"`
	Subject          *string          `type:"string"`
	SubjectType      *string          `type:"string"`

	metadataAssumeRoleWithSAMLOutput `json:"-", xml:"-"`
}

type metadataAssumeRoleWithSAMLOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type AssumeRoleWithWebIdentityInput struct {
	DurationSeconds  *int64  `type:"integer"`
	Policy           *string `type:"string"`
	ProviderID       *string `locationName:"ProviderId" type:"string"`
	RoleARN          *string `locationName:"RoleArn" type:"string" required:"true"`
	RoleSessionName  *string `type:"string" required:"true"`
	WebIdentityToken *string `type:"string" required:"true"`

	metadataAssumeRoleWithWebIdentityInput `json:"-", xml:"-"`
}

type metadataAssumeRoleWithWebIdentityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type AssumeRoleWithWebIdentityOutput struct {
	AssumedRoleUser             *AssumedRoleUser `type:"structure"`
	Audience                    *string          `type:"string"`
	Credentials                 *Credentials     `type:"structure"`
	PackedPolicySize            *int64           `type:"integer"`
	Provider                    *string          `type:"string"`
	SubjectFromWebIdentityToken *string          `type:"string"`

	metadataAssumeRoleWithWebIdentityOutput `json:"-", xml:"-"`
}

type metadataAssumeRoleWithWebIdentityOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type AssumedRoleUser struct {
	ARN           *string `locationName:"Arn" type:"string" required:"true"`
	AssumedRoleID *string `locationName:"AssumedRoleId" type:"string" required:"true"`

	metadataAssumedRoleUser `json:"-", xml:"-"`
}

type metadataAssumedRoleUser struct {
	SDKShapeTraits bool `type:"structure"`
}

type Credentials struct {
	AccessKeyID     *string    `locationName:"AccessKeyId" type:"string" required:"true"`
	Expiration      *time.Time `type:"timestamp" timestampFormat:"iso8601" required:"true"`
	SecretAccessKey *string    `type:"string" required:"true"`
	SessionToken    *string    `type:"string" required:"true"`

	metadataCredentials `json:"-", xml:"-"`
}

type metadataCredentials struct {
	SDKShapeTraits bool `type:"structure"`
}

type DecodeAuthorizationMessageInput struct {
	EncodedMessage *string `type:"string" required:"true"`

	metadataDecodeAuthorizationMessageInput `json:"-", xml:"-"`
}

type metadataDecodeAuthorizationMessageInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DecodeAuthorizationMessageOutput struct {
	DecodedMessage *string `type:"string"`

	metadataDecodeAuthorizationMessageOutput `json:"-", xml:"-"`
}

type metadataDecodeAuthorizationMessageOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type FederatedUser struct {
	ARN             *string `locationName:"Arn" type:"string" required:"true"`
	FederatedUserID *string `locationName:"FederatedUserId" type:"string" required:"true"`

	metadataFederatedUser `json:"-", xml:"-"`
}

type metadataFederatedUser struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetFederationTokenInput struct {
	DurationSeconds *int64  `type:"integer"`
	Name            *string `type:"string" required:"true"`
	Policy          *string `type:"string"`

	metadataGetFederationTokenInput `json:"-", xml:"-"`
}

type metadataGetFederationTokenInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetFederationTokenOutput struct {
	Credentials      *Credentials   `type:"structure"`
	FederatedUser    *FederatedUser `type:"structure"`
	PackedPolicySize *int64         `type:"integer"`

	metadataGetFederationTokenOutput `json:"-", xml:"-"`
}

type metadataGetFederationTokenOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetSessionTokenInput struct {
	DurationSeconds *int64  `type:"integer"`
	SerialNumber    *string `type:"string"`
	TokenCode       *string `type:"string"`

	metadataGetSessionTokenInput `json:"-", xml:"-"`
}

type metadataGetSessionTokenInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetSessionTokenOutput struct {
	Credentials *Credentials `type:"structure"`

	metadataGetSessionTokenOutput `json:"-", xml:"-"`
}

type metadataGetSessionTokenOutput struct {
	SDKShapeTraits bool `type:"structure"`
}