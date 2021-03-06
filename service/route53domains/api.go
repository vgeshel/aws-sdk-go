package route53domains

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// CheckDomainAvailabilityRequest generates a request for the CheckDomainAvailability operation.
func (c *Route53Domains) CheckDomainAvailabilityRequest(input *CheckDomainAvailabilityInput) (req *aws.Request, output *CheckDomainAvailabilityOutput) {
	if opCheckDomainAvailability == nil {
		opCheckDomainAvailability = &aws.Operation{
			Name:       "CheckDomainAvailability",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCheckDomainAvailability, input, output)
	output = &CheckDomainAvailabilityOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) CheckDomainAvailability(input *CheckDomainAvailabilityInput) (output *CheckDomainAvailabilityOutput, err error) {
	req, out := c.CheckDomainAvailabilityRequest(input)
	output = out
	err = req.Send()
	return
}

var opCheckDomainAvailability *aws.Operation

// DeleteTagsForDomainRequest generates a request for the DeleteTagsForDomain operation.
func (c *Route53Domains) DeleteTagsForDomainRequest(input *DeleteTagsForDomainInput) (req *aws.Request, output *DeleteTagsForDomainOutput) {
	if opDeleteTagsForDomain == nil {
		opDeleteTagsForDomain = &aws.Operation{
			Name:       "DeleteTagsForDomain",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteTagsForDomain, input, output)
	output = &DeleteTagsForDomainOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) DeleteTagsForDomain(input *DeleteTagsForDomainInput) (output *DeleteTagsForDomainOutput, err error) {
	req, out := c.DeleteTagsForDomainRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteTagsForDomain *aws.Operation

// DisableDomainAutoRenewRequest generates a request for the DisableDomainAutoRenew operation.
func (c *Route53Domains) DisableDomainAutoRenewRequest(input *DisableDomainAutoRenewInput) (req *aws.Request, output *DisableDomainAutoRenewOutput) {
	if opDisableDomainAutoRenew == nil {
		opDisableDomainAutoRenew = &aws.Operation{
			Name:       "DisableDomainAutoRenew",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDisableDomainAutoRenew, input, output)
	output = &DisableDomainAutoRenewOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) DisableDomainAutoRenew(input *DisableDomainAutoRenewInput) (output *DisableDomainAutoRenewOutput, err error) {
	req, out := c.DisableDomainAutoRenewRequest(input)
	output = out
	err = req.Send()
	return
}

var opDisableDomainAutoRenew *aws.Operation

// DisableDomainTransferLockRequest generates a request for the DisableDomainTransferLock operation.
func (c *Route53Domains) DisableDomainTransferLockRequest(input *DisableDomainTransferLockInput) (req *aws.Request, output *DisableDomainTransferLockOutput) {
	if opDisableDomainTransferLock == nil {
		opDisableDomainTransferLock = &aws.Operation{
			Name:       "DisableDomainTransferLock",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDisableDomainTransferLock, input, output)
	output = &DisableDomainTransferLockOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) DisableDomainTransferLock(input *DisableDomainTransferLockInput) (output *DisableDomainTransferLockOutput, err error) {
	req, out := c.DisableDomainTransferLockRequest(input)
	output = out
	err = req.Send()
	return
}

var opDisableDomainTransferLock *aws.Operation

// EnableDomainAutoRenewRequest generates a request for the EnableDomainAutoRenew operation.
func (c *Route53Domains) EnableDomainAutoRenewRequest(input *EnableDomainAutoRenewInput) (req *aws.Request, output *EnableDomainAutoRenewOutput) {
	if opEnableDomainAutoRenew == nil {
		opEnableDomainAutoRenew = &aws.Operation{
			Name:       "EnableDomainAutoRenew",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opEnableDomainAutoRenew, input, output)
	output = &EnableDomainAutoRenewOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) EnableDomainAutoRenew(input *EnableDomainAutoRenewInput) (output *EnableDomainAutoRenewOutput, err error) {
	req, out := c.EnableDomainAutoRenewRequest(input)
	output = out
	err = req.Send()
	return
}

var opEnableDomainAutoRenew *aws.Operation

// EnableDomainTransferLockRequest generates a request for the EnableDomainTransferLock operation.
func (c *Route53Domains) EnableDomainTransferLockRequest(input *EnableDomainTransferLockInput) (req *aws.Request, output *EnableDomainTransferLockOutput) {
	if opEnableDomainTransferLock == nil {
		opEnableDomainTransferLock = &aws.Operation{
			Name:       "EnableDomainTransferLock",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opEnableDomainTransferLock, input, output)
	output = &EnableDomainTransferLockOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) EnableDomainTransferLock(input *EnableDomainTransferLockInput) (output *EnableDomainTransferLockOutput, err error) {
	req, out := c.EnableDomainTransferLockRequest(input)
	output = out
	err = req.Send()
	return
}

var opEnableDomainTransferLock *aws.Operation

// GetDomainDetailRequest generates a request for the GetDomainDetail operation.
func (c *Route53Domains) GetDomainDetailRequest(input *GetDomainDetailInput) (req *aws.Request, output *GetDomainDetailOutput) {
	if opGetDomainDetail == nil {
		opGetDomainDetail = &aws.Operation{
			Name:       "GetDomainDetail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetDomainDetail, input, output)
	output = &GetDomainDetailOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) GetDomainDetail(input *GetDomainDetailInput) (output *GetDomainDetailOutput, err error) {
	req, out := c.GetDomainDetailRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetDomainDetail *aws.Operation

// GetOperationDetailRequest generates a request for the GetOperationDetail operation.
func (c *Route53Domains) GetOperationDetailRequest(input *GetOperationDetailInput) (req *aws.Request, output *GetOperationDetailOutput) {
	if opGetOperationDetail == nil {
		opGetOperationDetail = &aws.Operation{
			Name:       "GetOperationDetail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetOperationDetail, input, output)
	output = &GetOperationDetailOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) GetOperationDetail(input *GetOperationDetailInput) (output *GetOperationDetailOutput, err error) {
	req, out := c.GetOperationDetailRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetOperationDetail *aws.Operation

// ListDomainsRequest generates a request for the ListDomains operation.
func (c *Route53Domains) ListDomainsRequest(input *ListDomainsInput) (req *aws.Request, output *ListDomainsOutput) {
	if opListDomains == nil {
		opListDomains = &aws.Operation{
			Name:       "ListDomains",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListDomains, input, output)
	output = &ListDomainsOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) ListDomains(input *ListDomainsInput) (output *ListDomainsOutput, err error) {
	req, out := c.ListDomainsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListDomains *aws.Operation

// ListOperationsRequest generates a request for the ListOperations operation.
func (c *Route53Domains) ListOperationsRequest(input *ListOperationsInput) (req *aws.Request, output *ListOperationsOutput) {
	if opListOperations == nil {
		opListOperations = &aws.Operation{
			Name:       "ListOperations",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListOperations, input, output)
	output = &ListOperationsOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) ListOperations(input *ListOperationsInput) (output *ListOperationsOutput, err error) {
	req, out := c.ListOperationsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListOperations *aws.Operation

// ListTagsForDomainRequest generates a request for the ListTagsForDomain operation.
func (c *Route53Domains) ListTagsForDomainRequest(input *ListTagsForDomainInput) (req *aws.Request, output *ListTagsForDomainOutput) {
	if opListTagsForDomain == nil {
		opListTagsForDomain = &aws.Operation{
			Name:       "ListTagsForDomain",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListTagsForDomain, input, output)
	output = &ListTagsForDomainOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) ListTagsForDomain(input *ListTagsForDomainInput) (output *ListTagsForDomainOutput, err error) {
	req, out := c.ListTagsForDomainRequest(input)
	output = out
	err = req.Send()
	return
}

var opListTagsForDomain *aws.Operation

// RegisterDomainRequest generates a request for the RegisterDomain operation.
func (c *Route53Domains) RegisterDomainRequest(input *RegisterDomainInput) (req *aws.Request, output *RegisterDomainOutput) {
	if opRegisterDomain == nil {
		opRegisterDomain = &aws.Operation{
			Name:       "RegisterDomain",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opRegisterDomain, input, output)
	output = &RegisterDomainOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) RegisterDomain(input *RegisterDomainInput) (output *RegisterDomainOutput, err error) {
	req, out := c.RegisterDomainRequest(input)
	output = out
	err = req.Send()
	return
}

var opRegisterDomain *aws.Operation

// RetrieveDomainAuthCodeRequest generates a request for the RetrieveDomainAuthCode operation.
func (c *Route53Domains) RetrieveDomainAuthCodeRequest(input *RetrieveDomainAuthCodeInput) (req *aws.Request, output *RetrieveDomainAuthCodeOutput) {
	if opRetrieveDomainAuthCode == nil {
		opRetrieveDomainAuthCode = &aws.Operation{
			Name:       "RetrieveDomainAuthCode",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opRetrieveDomainAuthCode, input, output)
	output = &RetrieveDomainAuthCodeOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) RetrieveDomainAuthCode(input *RetrieveDomainAuthCodeInput) (output *RetrieveDomainAuthCodeOutput, err error) {
	req, out := c.RetrieveDomainAuthCodeRequest(input)
	output = out
	err = req.Send()
	return
}

var opRetrieveDomainAuthCode *aws.Operation

// TransferDomainRequest generates a request for the TransferDomain operation.
func (c *Route53Domains) TransferDomainRequest(input *TransferDomainInput) (req *aws.Request, output *TransferDomainOutput) {
	if opTransferDomain == nil {
		opTransferDomain = &aws.Operation{
			Name:       "TransferDomain",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opTransferDomain, input, output)
	output = &TransferDomainOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) TransferDomain(input *TransferDomainInput) (output *TransferDomainOutput, err error) {
	req, out := c.TransferDomainRequest(input)
	output = out
	err = req.Send()
	return
}

var opTransferDomain *aws.Operation

// UpdateDomainContactRequest generates a request for the UpdateDomainContact operation.
func (c *Route53Domains) UpdateDomainContactRequest(input *UpdateDomainContactInput) (req *aws.Request, output *UpdateDomainContactOutput) {
	if opUpdateDomainContact == nil {
		opUpdateDomainContact = &aws.Operation{
			Name:       "UpdateDomainContact",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateDomainContact, input, output)
	output = &UpdateDomainContactOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) UpdateDomainContact(input *UpdateDomainContactInput) (output *UpdateDomainContactOutput, err error) {
	req, out := c.UpdateDomainContactRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateDomainContact *aws.Operation

// UpdateDomainContactPrivacyRequest generates a request for the UpdateDomainContactPrivacy operation.
func (c *Route53Domains) UpdateDomainContactPrivacyRequest(input *UpdateDomainContactPrivacyInput) (req *aws.Request, output *UpdateDomainContactPrivacyOutput) {
	if opUpdateDomainContactPrivacy == nil {
		opUpdateDomainContactPrivacy = &aws.Operation{
			Name:       "UpdateDomainContactPrivacy",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateDomainContactPrivacy, input, output)
	output = &UpdateDomainContactPrivacyOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) UpdateDomainContactPrivacy(input *UpdateDomainContactPrivacyInput) (output *UpdateDomainContactPrivacyOutput, err error) {
	req, out := c.UpdateDomainContactPrivacyRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateDomainContactPrivacy *aws.Operation

// UpdateDomainNameserversRequest generates a request for the UpdateDomainNameservers operation.
func (c *Route53Domains) UpdateDomainNameserversRequest(input *UpdateDomainNameserversInput) (req *aws.Request, output *UpdateDomainNameserversOutput) {
	if opUpdateDomainNameservers == nil {
		opUpdateDomainNameservers = &aws.Operation{
			Name:       "UpdateDomainNameservers",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateDomainNameservers, input, output)
	output = &UpdateDomainNameserversOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) UpdateDomainNameservers(input *UpdateDomainNameserversInput) (output *UpdateDomainNameserversOutput, err error) {
	req, out := c.UpdateDomainNameserversRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateDomainNameservers *aws.Operation

// UpdateTagsForDomainRequest generates a request for the UpdateTagsForDomain operation.
func (c *Route53Domains) UpdateTagsForDomainRequest(input *UpdateTagsForDomainInput) (req *aws.Request, output *UpdateTagsForDomainOutput) {
	if opUpdateTagsForDomain == nil {
		opUpdateTagsForDomain = &aws.Operation{
			Name:       "UpdateTagsForDomain",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateTagsForDomain, input, output)
	output = &UpdateTagsForDomainOutput{}
	req.Data = output
	return
}

func (c *Route53Domains) UpdateTagsForDomain(input *UpdateTagsForDomainInput) (output *UpdateTagsForDomainOutput, err error) {
	req, out := c.UpdateTagsForDomainRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateTagsForDomain *aws.Operation

type CheckDomainAvailabilityInput struct {
	DomainName  *string `type:"string" required:"true"`
	IDNLangCode *string `locationName:"IdnLangCode" type:"string"`

	metadataCheckDomainAvailabilityInput `json:"-", xml:"-"`
}

type metadataCheckDomainAvailabilityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CheckDomainAvailabilityOutput struct {
	Availability *string `type:"string" required:"true"`

	metadataCheckDomainAvailabilityOutput `json:"-", xml:"-"`
}

type metadataCheckDomainAvailabilityOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ContactDetail struct {
	AddressLine1     *string       `type:"string"`
	AddressLine2     *string       `type:"string"`
	City             *string       `type:"string"`
	ContactType      *string       `type:"string"`
	CountryCode      *string       `type:"string"`
	Email            *string       `type:"string"`
	ExtraParams      []*ExtraParam `type:"list"`
	Fax              *string       `type:"string"`
	FirstName        *string       `type:"string"`
	LastName         *string       `type:"string"`
	OrganizationName *string       `type:"string"`
	PhoneNumber      *string       `type:"string"`
	State            *string       `type:"string"`
	ZipCode          *string       `type:"string"`

	metadataContactDetail `json:"-", xml:"-"`
}

type metadataContactDetail struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteTagsForDomainInput struct {
	DomainName   *string   `type:"string" required:"true"`
	TagsToDelete []*string `type:"list" required:"true"`

	metadataDeleteTagsForDomainInput `json:"-", xml:"-"`
}

type metadataDeleteTagsForDomainInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteTagsForDomainOutput struct {
	metadataDeleteTagsForDomainOutput `json:"-", xml:"-"`
}

type metadataDeleteTagsForDomainOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DisableDomainAutoRenewInput struct {
	DomainName *string `type:"string" required:"true"`

	metadataDisableDomainAutoRenewInput `json:"-", xml:"-"`
}

type metadataDisableDomainAutoRenewInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DisableDomainAutoRenewOutput struct {
	metadataDisableDomainAutoRenewOutput `json:"-", xml:"-"`
}

type metadataDisableDomainAutoRenewOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DisableDomainTransferLockInput struct {
	DomainName *string `type:"string" required:"true"`

	metadataDisableDomainTransferLockInput `json:"-", xml:"-"`
}

type metadataDisableDomainTransferLockInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DisableDomainTransferLockOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" required:"true"`

	metadataDisableDomainTransferLockOutput `json:"-", xml:"-"`
}

type metadataDisableDomainTransferLockOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DomainSummary struct {
	AutoRenew    *bool      `type:"boolean"`
	DomainName   *string    `type:"string" required:"true"`
	Expiry       *time.Time `type:"timestamp" timestampFormat:"unix"`
	TransferLock *bool      `type:"boolean"`

	metadataDomainSummary `json:"-", xml:"-"`
}

type metadataDomainSummary struct {
	SDKShapeTraits bool `type:"structure"`
}

type EnableDomainAutoRenewInput struct {
	DomainName *string `type:"string" required:"true"`

	metadataEnableDomainAutoRenewInput `json:"-", xml:"-"`
}

type metadataEnableDomainAutoRenewInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type EnableDomainAutoRenewOutput struct {
	metadataEnableDomainAutoRenewOutput `json:"-", xml:"-"`
}

type metadataEnableDomainAutoRenewOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type EnableDomainTransferLockInput struct {
	DomainName *string `type:"string" required:"true"`

	metadataEnableDomainTransferLockInput `json:"-", xml:"-"`
}

type metadataEnableDomainTransferLockInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type EnableDomainTransferLockOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" required:"true"`

	metadataEnableDomainTransferLockOutput `json:"-", xml:"-"`
}

type metadataEnableDomainTransferLockOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ExtraParam struct {
	Name  *string `type:"string" required:"true"`
	Value *string `type:"string" required:"true"`

	metadataExtraParam `json:"-", xml:"-"`
}

type metadataExtraParam struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetDomainDetailInput struct {
	DomainName *string `type:"string" required:"true"`

	metadataGetDomainDetailInput `json:"-", xml:"-"`
}

type metadataGetDomainDetailInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetDomainDetailOutput struct {
	AbuseContactEmail *string        `type:"string"`
	AbuseContactPhone *string        `type:"string"`
	AdminContact      *ContactDetail `type:"structure" required:"true"`
	AdminPrivacy      *bool          `type:"boolean"`
	AutoRenew         *bool          `type:"boolean"`
	CreationDate      *time.Time     `type:"timestamp" timestampFormat:"unix"`
	DNSSec            *string        `locationName:"DnsSec" type:"string"`
	DomainName        *string        `type:"string" required:"true"`
	ExpirationDate    *time.Time     `type:"timestamp" timestampFormat:"unix"`
	Nameservers       []*Nameserver  `type:"list" required:"true"`
	RegistrantContact *ContactDetail `type:"structure" required:"true"`
	RegistrantPrivacy *bool          `type:"boolean"`
	RegistrarName     *string        `type:"string"`
	RegistrarURL      *string        `locationName:"RegistrarUrl" type:"string"`
	RegistryDomainID  *string        `locationName:"RegistryDomainId" type:"string"`
	Reseller          *string        `type:"string"`
	StatusList        []*string      `type:"list"`
	TechContact       *ContactDetail `type:"structure" required:"true"`
	TechPrivacy       *bool          `type:"boolean"`
	UpdatedDate       *time.Time     `type:"timestamp" timestampFormat:"unix"`
	WhoIsServer       *string        `type:"string"`

	metadataGetDomainDetailOutput `json:"-", xml:"-"`
}

type metadataGetDomainDetailOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetOperationDetailInput struct {
	OperationID *string `locationName:"OperationId" type:"string" required:"true"`

	metadataGetOperationDetailInput `json:"-", xml:"-"`
}

type metadataGetOperationDetailInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetOperationDetailOutput struct {
	DomainName    *string    `type:"string"`
	Message       *string    `type:"string"`
	OperationID   *string    `locationName:"OperationId" type:"string"`
	Status        *string    `type:"string"`
	SubmittedDate *time.Time `type:"timestamp" timestampFormat:"unix"`
	Type          *string    `type:"string"`

	metadataGetOperationDetailOutput `json:"-", xml:"-"`
}

type metadataGetOperationDetailOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListDomainsInput struct {
	Marker   *string `type:"string"`
	MaxItems *int64  `type:"integer"`

	metadataListDomainsInput `json:"-", xml:"-"`
}

type metadataListDomainsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListDomainsOutput struct {
	Domains        []*DomainSummary `type:"list" required:"true"`
	NextPageMarker *string          `type:"string"`

	metadataListDomainsOutput `json:"-", xml:"-"`
}

type metadataListDomainsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListOperationsInput struct {
	Marker   *string `type:"string"`
	MaxItems *int64  `type:"integer"`

	metadataListOperationsInput `json:"-", xml:"-"`
}

type metadataListOperationsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListOperationsOutput struct {
	NextPageMarker *string             `type:"string"`
	Operations     []*OperationSummary `type:"list" required:"true"`

	metadataListOperationsOutput `json:"-", xml:"-"`
}

type metadataListOperationsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListTagsForDomainInput struct {
	DomainName *string `type:"string" required:"true"`

	metadataListTagsForDomainInput `json:"-", xml:"-"`
}

type metadataListTagsForDomainInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListTagsForDomainOutput struct {
	TagList []*Tag `type:"list" required:"true"`

	metadataListTagsForDomainOutput `json:"-", xml:"-"`
}

type metadataListTagsForDomainOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Nameserver struct {
	GlueIPs []*string `locationName:"GlueIps" type:"list"`
	Name    *string   `type:"string" required:"true"`

	metadataNameserver `json:"-", xml:"-"`
}

type metadataNameserver struct {
	SDKShapeTraits bool `type:"structure"`
}

type OperationSummary struct {
	OperationID   *string    `locationName:"OperationId" type:"string" required:"true"`
	Status        *string    `type:"string" required:"true"`
	SubmittedDate *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`
	Type          *string    `type:"string" required:"true"`

	metadataOperationSummary `json:"-", xml:"-"`
}

type metadataOperationSummary struct {
	SDKShapeTraits bool `type:"structure"`
}

type RegisterDomainInput struct {
	AdminContact                    *ContactDetail `type:"structure" required:"true"`
	AutoRenew                       *bool          `type:"boolean"`
	DomainName                      *string        `type:"string" required:"true"`
	DurationInYears                 *int64         `type:"integer" required:"true"`
	IDNLangCode                     *string        `locationName:"IdnLangCode" type:"string"`
	PrivacyProtectAdminContact      *bool          `type:"boolean"`
	PrivacyProtectRegistrantContact *bool          `type:"boolean"`
	PrivacyProtectTechContact       *bool          `type:"boolean"`
	RegistrantContact               *ContactDetail `type:"structure" required:"true"`
	TechContact                     *ContactDetail `type:"structure" required:"true"`

	metadataRegisterDomainInput `json:"-", xml:"-"`
}

type metadataRegisterDomainInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RegisterDomainOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" required:"true"`

	metadataRegisterDomainOutput `json:"-", xml:"-"`
}

type metadataRegisterDomainOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RetrieveDomainAuthCodeInput struct {
	DomainName *string `type:"string" required:"true"`

	metadataRetrieveDomainAuthCodeInput `json:"-", xml:"-"`
}

type metadataRetrieveDomainAuthCodeInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RetrieveDomainAuthCodeOutput struct {
	AuthCode *string `type:"string" required:"true"`

	metadataRetrieveDomainAuthCodeOutput `json:"-", xml:"-"`
}

type metadataRetrieveDomainAuthCodeOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Tag struct {
	Key   *string `type:"string"`
	Value *string `type:"string"`

	metadataTag `json:"-", xml:"-"`
}

type metadataTag struct {
	SDKShapeTraits bool `type:"structure"`
}

type TransferDomainInput struct {
	AdminContact                    *ContactDetail `type:"structure" required:"true"`
	AuthCode                        *string        `type:"string"`
	AutoRenew                       *bool          `type:"boolean"`
	DomainName                      *string        `type:"string" required:"true"`
	DurationInYears                 *int64         `type:"integer" required:"true"`
	IDNLangCode                     *string        `locationName:"IdnLangCode" type:"string"`
	Nameservers                     []*Nameserver  `type:"list"`
	PrivacyProtectAdminContact      *bool          `type:"boolean"`
	PrivacyProtectRegistrantContact *bool          `type:"boolean"`
	PrivacyProtectTechContact       *bool          `type:"boolean"`
	RegistrantContact               *ContactDetail `type:"structure" required:"true"`
	TechContact                     *ContactDetail `type:"structure" required:"true"`

	metadataTransferDomainInput `json:"-", xml:"-"`
}

type metadataTransferDomainInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type TransferDomainOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" required:"true"`

	metadataTransferDomainOutput `json:"-", xml:"-"`
}

type metadataTransferDomainOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateDomainContactInput struct {
	AdminContact      *ContactDetail `type:"structure"`
	DomainName        *string        `type:"string" required:"true"`
	RegistrantContact *ContactDetail `type:"structure"`
	TechContact       *ContactDetail `type:"structure"`

	metadataUpdateDomainContactInput `json:"-", xml:"-"`
}

type metadataUpdateDomainContactInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateDomainContactOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" required:"true"`

	metadataUpdateDomainContactOutput `json:"-", xml:"-"`
}

type metadataUpdateDomainContactOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateDomainContactPrivacyInput struct {
	AdminPrivacy      *bool   `type:"boolean"`
	DomainName        *string `type:"string" required:"true"`
	RegistrantPrivacy *bool   `type:"boolean"`
	TechPrivacy       *bool   `type:"boolean"`

	metadataUpdateDomainContactPrivacyInput `json:"-", xml:"-"`
}

type metadataUpdateDomainContactPrivacyInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateDomainContactPrivacyOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" required:"true"`

	metadataUpdateDomainContactPrivacyOutput `json:"-", xml:"-"`
}

type metadataUpdateDomainContactPrivacyOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateDomainNameserversInput struct {
	DomainName  *string       `type:"string" required:"true"`
	FIAuthKey   *string       `type:"string"`
	Nameservers []*Nameserver `type:"list" required:"true"`

	metadataUpdateDomainNameserversInput `json:"-", xml:"-"`
}

type metadataUpdateDomainNameserversInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateDomainNameserversOutput struct {
	OperationID *string `locationName:"OperationId" type:"string" required:"true"`

	metadataUpdateDomainNameserversOutput `json:"-", xml:"-"`
}

type metadataUpdateDomainNameserversOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateTagsForDomainInput struct {
	DomainName   *string `type:"string" required:"true"`
	TagsToUpdate []*Tag  `type:"list"`

	metadataUpdateTagsForDomainInput `json:"-", xml:"-"`
}

type metadataUpdateTagsForDomainInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateTagsForDomainOutput struct {
	metadataUpdateTagsForDomainOutput `json:"-", xml:"-"`
}

type metadataUpdateTagsForDomainOutput struct {
	SDKShapeTraits bool `type:"structure"`
}