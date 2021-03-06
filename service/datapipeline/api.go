package datapipeline

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"github.com/awslabs/aws-sdk-go/aws"
)

// ActivatePipelineRequest generates a request for the ActivatePipeline operation.
func (c *DataPipeline) ActivatePipelineRequest(input *ActivatePipelineInput) (req *aws.Request, output *ActivatePipelineOutput) {
	if opActivatePipeline == nil {
		opActivatePipeline = &aws.Operation{
			Name:       "ActivatePipeline",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opActivatePipeline, input, output)
	output = &ActivatePipelineOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) ActivatePipeline(input *ActivatePipelineInput) (output *ActivatePipelineOutput, err error) {
	req, out := c.ActivatePipelineRequest(input)
	output = out
	err = req.Send()
	return
}

var opActivatePipeline *aws.Operation

// CreatePipelineRequest generates a request for the CreatePipeline operation.
func (c *DataPipeline) CreatePipelineRequest(input *CreatePipelineInput) (req *aws.Request, output *CreatePipelineOutput) {
	if opCreatePipeline == nil {
		opCreatePipeline = &aws.Operation{
			Name:       "CreatePipeline",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreatePipeline, input, output)
	output = &CreatePipelineOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) CreatePipeline(input *CreatePipelineInput) (output *CreatePipelineOutput, err error) {
	req, out := c.CreatePipelineRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreatePipeline *aws.Operation

// DeletePipelineRequest generates a request for the DeletePipeline operation.
func (c *DataPipeline) DeletePipelineRequest(input *DeletePipelineInput) (req *aws.Request, output *DeletePipelineOutput) {
	if opDeletePipeline == nil {
		opDeletePipeline = &aws.Operation{
			Name:       "DeletePipeline",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeletePipeline, input, output)
	output = &DeletePipelineOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) DeletePipeline(input *DeletePipelineInput) (output *DeletePipelineOutput, err error) {
	req, out := c.DeletePipelineRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeletePipeline *aws.Operation

// DescribeObjectsRequest generates a request for the DescribeObjects operation.
func (c *DataPipeline) DescribeObjectsRequest(input *DescribeObjectsInput) (req *aws.Request, output *DescribeObjectsOutput) {
	if opDescribeObjects == nil {
		opDescribeObjects = &aws.Operation{
			Name:       "DescribeObjects",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeObjects, input, output)
	output = &DescribeObjectsOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) DescribeObjects(input *DescribeObjectsInput) (output *DescribeObjectsOutput, err error) {
	req, out := c.DescribeObjectsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeObjects *aws.Operation

// DescribePipelinesRequest generates a request for the DescribePipelines operation.
func (c *DataPipeline) DescribePipelinesRequest(input *DescribePipelinesInput) (req *aws.Request, output *DescribePipelinesOutput) {
	if opDescribePipelines == nil {
		opDescribePipelines = &aws.Operation{
			Name:       "DescribePipelines",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribePipelines, input, output)
	output = &DescribePipelinesOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) DescribePipelines(input *DescribePipelinesInput) (output *DescribePipelinesOutput, err error) {
	req, out := c.DescribePipelinesRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribePipelines *aws.Operation

// EvaluateExpressionRequest generates a request for the EvaluateExpression operation.
func (c *DataPipeline) EvaluateExpressionRequest(input *EvaluateExpressionInput) (req *aws.Request, output *EvaluateExpressionOutput) {
	if opEvaluateExpression == nil {
		opEvaluateExpression = &aws.Operation{
			Name:       "EvaluateExpression",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opEvaluateExpression, input, output)
	output = &EvaluateExpressionOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) EvaluateExpression(input *EvaluateExpressionInput) (output *EvaluateExpressionOutput, err error) {
	req, out := c.EvaluateExpressionRequest(input)
	output = out
	err = req.Send()
	return
}

var opEvaluateExpression *aws.Operation

// GetPipelineDefinitionRequest generates a request for the GetPipelineDefinition operation.
func (c *DataPipeline) GetPipelineDefinitionRequest(input *GetPipelineDefinitionInput) (req *aws.Request, output *GetPipelineDefinitionOutput) {
	if opGetPipelineDefinition == nil {
		opGetPipelineDefinition = &aws.Operation{
			Name:       "GetPipelineDefinition",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetPipelineDefinition, input, output)
	output = &GetPipelineDefinitionOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) GetPipelineDefinition(input *GetPipelineDefinitionInput) (output *GetPipelineDefinitionOutput, err error) {
	req, out := c.GetPipelineDefinitionRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetPipelineDefinition *aws.Operation

// ListPipelinesRequest generates a request for the ListPipelines operation.
func (c *DataPipeline) ListPipelinesRequest(input *ListPipelinesInput) (req *aws.Request, output *ListPipelinesOutput) {
	if opListPipelines == nil {
		opListPipelines = &aws.Operation{
			Name:       "ListPipelines",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListPipelines, input, output)
	output = &ListPipelinesOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) ListPipelines(input *ListPipelinesInput) (output *ListPipelinesOutput, err error) {
	req, out := c.ListPipelinesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListPipelines *aws.Operation

// PollForTaskRequest generates a request for the PollForTask operation.
func (c *DataPipeline) PollForTaskRequest(input *PollForTaskInput) (req *aws.Request, output *PollForTaskOutput) {
	if opPollForTask == nil {
		opPollForTask = &aws.Operation{
			Name:       "PollForTask",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPollForTask, input, output)
	output = &PollForTaskOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) PollForTask(input *PollForTaskInput) (output *PollForTaskOutput, err error) {
	req, out := c.PollForTaskRequest(input)
	output = out
	err = req.Send()
	return
}

var opPollForTask *aws.Operation

// PutPipelineDefinitionRequest generates a request for the PutPipelineDefinition operation.
func (c *DataPipeline) PutPipelineDefinitionRequest(input *PutPipelineDefinitionInput) (req *aws.Request, output *PutPipelineDefinitionOutput) {
	if opPutPipelineDefinition == nil {
		opPutPipelineDefinition = &aws.Operation{
			Name:       "PutPipelineDefinition",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPutPipelineDefinition, input, output)
	output = &PutPipelineDefinitionOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) PutPipelineDefinition(input *PutPipelineDefinitionInput) (output *PutPipelineDefinitionOutput, err error) {
	req, out := c.PutPipelineDefinitionRequest(input)
	output = out
	err = req.Send()
	return
}

var opPutPipelineDefinition *aws.Operation

// QueryObjectsRequest generates a request for the QueryObjects operation.
func (c *DataPipeline) QueryObjectsRequest(input *QueryObjectsInput) (req *aws.Request, output *QueryObjectsOutput) {
	if opQueryObjects == nil {
		opQueryObjects = &aws.Operation{
			Name:       "QueryObjects",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opQueryObjects, input, output)
	output = &QueryObjectsOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) QueryObjects(input *QueryObjectsInput) (output *QueryObjectsOutput, err error) {
	req, out := c.QueryObjectsRequest(input)
	output = out
	err = req.Send()
	return
}

var opQueryObjects *aws.Operation

// ReportTaskProgressRequest generates a request for the ReportTaskProgress operation.
func (c *DataPipeline) ReportTaskProgressRequest(input *ReportTaskProgressInput) (req *aws.Request, output *ReportTaskProgressOutput) {
	if opReportTaskProgress == nil {
		opReportTaskProgress = &aws.Operation{
			Name:       "ReportTaskProgress",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opReportTaskProgress, input, output)
	output = &ReportTaskProgressOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) ReportTaskProgress(input *ReportTaskProgressInput) (output *ReportTaskProgressOutput, err error) {
	req, out := c.ReportTaskProgressRequest(input)
	output = out
	err = req.Send()
	return
}

var opReportTaskProgress *aws.Operation

// ReportTaskRunnerHeartbeatRequest generates a request for the ReportTaskRunnerHeartbeat operation.
func (c *DataPipeline) ReportTaskRunnerHeartbeatRequest(input *ReportTaskRunnerHeartbeatInput) (req *aws.Request, output *ReportTaskRunnerHeartbeatOutput) {
	if opReportTaskRunnerHeartbeat == nil {
		opReportTaskRunnerHeartbeat = &aws.Operation{
			Name:       "ReportTaskRunnerHeartbeat",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opReportTaskRunnerHeartbeat, input, output)
	output = &ReportTaskRunnerHeartbeatOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) ReportTaskRunnerHeartbeat(input *ReportTaskRunnerHeartbeatInput) (output *ReportTaskRunnerHeartbeatOutput, err error) {
	req, out := c.ReportTaskRunnerHeartbeatRequest(input)
	output = out
	err = req.Send()
	return
}

var opReportTaskRunnerHeartbeat *aws.Operation

// SetStatusRequest generates a request for the SetStatus operation.
func (c *DataPipeline) SetStatusRequest(input *SetStatusInput) (req *aws.Request, output *SetStatusOutput) {
	if opSetStatus == nil {
		opSetStatus = &aws.Operation{
			Name:       "SetStatus",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetStatus, input, output)
	output = &SetStatusOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) SetStatus(input *SetStatusInput) (output *SetStatusOutput, err error) {
	req, out := c.SetStatusRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetStatus *aws.Operation

// SetTaskStatusRequest generates a request for the SetTaskStatus operation.
func (c *DataPipeline) SetTaskStatusRequest(input *SetTaskStatusInput) (req *aws.Request, output *SetTaskStatusOutput) {
	if opSetTaskStatus == nil {
		opSetTaskStatus = &aws.Operation{
			Name:       "SetTaskStatus",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetTaskStatus, input, output)
	output = &SetTaskStatusOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) SetTaskStatus(input *SetTaskStatusInput) (output *SetTaskStatusOutput, err error) {
	req, out := c.SetTaskStatusRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetTaskStatus *aws.Operation

// ValidatePipelineDefinitionRequest generates a request for the ValidatePipelineDefinition operation.
func (c *DataPipeline) ValidatePipelineDefinitionRequest(input *ValidatePipelineDefinitionInput) (req *aws.Request, output *ValidatePipelineDefinitionOutput) {
	if opValidatePipelineDefinition == nil {
		opValidatePipelineDefinition = &aws.Operation{
			Name:       "ValidatePipelineDefinition",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opValidatePipelineDefinition, input, output)
	output = &ValidatePipelineDefinitionOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) ValidatePipelineDefinition(input *ValidatePipelineDefinitionInput) (output *ValidatePipelineDefinitionOutput, err error) {
	req, out := c.ValidatePipelineDefinitionRequest(input)
	output = out
	err = req.Send()
	return
}

var opValidatePipelineDefinition *aws.Operation

type ActivatePipelineInput struct {
	ParameterValues []*ParameterValue `locationName:"parameterValues" type:"list"`
	PipelineID      *string           `locationName:"pipelineId" type:"string" required:"true"`

	metadataActivatePipelineInput `json:"-", xml:"-"`
}

type metadataActivatePipelineInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ActivatePipelineOutput struct {
	metadataActivatePipelineOutput `json:"-", xml:"-"`
}

type metadataActivatePipelineOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreatePipelineInput struct {
	Description *string `locationName:"description" type:"string"`
	Name        *string `locationName:"name" type:"string" required:"true"`
	UniqueID    *string `locationName:"uniqueId" type:"string" required:"true"`

	metadataCreatePipelineInput `json:"-", xml:"-"`
}

type metadataCreatePipelineInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreatePipelineOutput struct {
	PipelineID *string `locationName:"pipelineId" type:"string" required:"true"`

	metadataCreatePipelineOutput `json:"-", xml:"-"`
}

type metadataCreatePipelineOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeletePipelineInput struct {
	PipelineID *string `locationName:"pipelineId" type:"string" required:"true"`

	metadataDeletePipelineInput `json:"-", xml:"-"`
}

type metadataDeletePipelineInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeletePipelineOutput struct {
	metadataDeletePipelineOutput `json:"-", xml:"-"`
}

type metadataDeletePipelineOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeObjectsInput struct {
	EvaluateExpressions *bool     `locationName:"evaluateExpressions" type:"boolean"`
	Marker              *string   `locationName:"marker" type:"string"`
	ObjectIDs           []*string `locationName:"objectIds" type:"list" required:"true"`
	PipelineID          *string   `locationName:"pipelineId" type:"string" required:"true"`

	metadataDescribeObjectsInput `json:"-", xml:"-"`
}

type metadataDescribeObjectsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeObjectsOutput struct {
	HasMoreResults  *bool             `locationName:"hasMoreResults" type:"boolean"`
	Marker          *string           `locationName:"marker" type:"string"`
	PipelineObjects []*PipelineObject `locationName:"pipelineObjects" type:"list" required:"true"`

	metadataDescribeObjectsOutput `json:"-", xml:"-"`
}

type metadataDescribeObjectsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribePipelinesInput struct {
	PipelineIDs []*string `locationName:"pipelineIds" type:"list" required:"true"`

	metadataDescribePipelinesInput `json:"-", xml:"-"`
}

type metadataDescribePipelinesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribePipelinesOutput struct {
	PipelineDescriptionList []*PipelineDescription `locationName:"pipelineDescriptionList" type:"list" required:"true"`

	metadataDescribePipelinesOutput `json:"-", xml:"-"`
}

type metadataDescribePipelinesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type EvaluateExpressionInput struct {
	Expression *string `locationName:"expression" type:"string" required:"true"`
	ObjectID   *string `locationName:"objectId" type:"string" required:"true"`
	PipelineID *string `locationName:"pipelineId" type:"string" required:"true"`

	metadataEvaluateExpressionInput `json:"-", xml:"-"`
}

type metadataEvaluateExpressionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type EvaluateExpressionOutput struct {
	EvaluatedExpression *string `locationName:"evaluatedExpression" type:"string" required:"true"`

	metadataEvaluateExpressionOutput `json:"-", xml:"-"`
}

type metadataEvaluateExpressionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Field struct {
	Key         *string `locationName:"key" type:"string" required:"true"`
	RefValue    *string `locationName:"refValue" type:"string"`
	StringValue *string `locationName:"stringValue" type:"string"`

	metadataField `json:"-", xml:"-"`
}

type metadataField struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetPipelineDefinitionInput struct {
	PipelineID *string `locationName:"pipelineId" type:"string" required:"true"`
	Version    *string `locationName:"version" type:"string"`

	metadataGetPipelineDefinitionInput `json:"-", xml:"-"`
}

type metadataGetPipelineDefinitionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetPipelineDefinitionOutput struct {
	ParameterObjects []*ParameterObject `locationName:"parameterObjects" type:"list"`
	ParameterValues  []*ParameterValue  `locationName:"parameterValues" type:"list"`
	PipelineObjects  []*PipelineObject  `locationName:"pipelineObjects" type:"list"`

	metadataGetPipelineDefinitionOutput `json:"-", xml:"-"`
}

type metadataGetPipelineDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type InstanceIdentity struct {
	Document  *string `locationName:"document" type:"string"`
	Signature *string `locationName:"signature" type:"string"`

	metadataInstanceIdentity `json:"-", xml:"-"`
}

type metadataInstanceIdentity struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListPipelinesInput struct {
	Marker *string `locationName:"marker" type:"string"`

	metadataListPipelinesInput `json:"-", xml:"-"`
}

type metadataListPipelinesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListPipelinesOutput struct {
	HasMoreResults *bool             `locationName:"hasMoreResults" type:"boolean"`
	Marker         *string           `locationName:"marker" type:"string"`
	PipelineIDList []*PipelineIDName `locationName:"pipelineIdList" type:"list" required:"true"`

	metadataListPipelinesOutput `json:"-", xml:"-"`
}

type metadataListPipelinesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Operator struct {
	Type   *string   `locationName:"type" type:"string"`
	Values []*string `locationName:"values" type:"list"`

	metadataOperator `json:"-", xml:"-"`
}

type metadataOperator struct {
	SDKShapeTraits bool `type:"structure"`
}

type ParameterAttribute struct {
	Key         *string `locationName:"key" type:"string" required:"true"`
	StringValue *string `locationName:"stringValue" type:"string" required:"true"`

	metadataParameterAttribute `json:"-", xml:"-"`
}

type metadataParameterAttribute struct {
	SDKShapeTraits bool `type:"structure"`
}

type ParameterObject struct {
	Attributes []*ParameterAttribute `locationName:"attributes" type:"list" required:"true"`
	ID         *string               `locationName:"id" type:"string" required:"true"`

	metadataParameterObject `json:"-", xml:"-"`
}

type metadataParameterObject struct {
	SDKShapeTraits bool `type:"structure"`
}

type ParameterValue struct {
	ID          *string `locationName:"id" type:"string" required:"true"`
	StringValue *string `locationName:"stringValue" type:"string" required:"true"`

	metadataParameterValue `json:"-", xml:"-"`
}

type metadataParameterValue struct {
	SDKShapeTraits bool `type:"structure"`
}

type PipelineDescription struct {
	Description *string  `locationName:"description" type:"string"`
	Fields      []*Field `locationName:"fields" type:"list" required:"true"`
	Name        *string  `locationName:"name" type:"string" required:"true"`
	PipelineID  *string  `locationName:"pipelineId" type:"string" required:"true"`

	metadataPipelineDescription `json:"-", xml:"-"`
}

type metadataPipelineDescription struct {
	SDKShapeTraits bool `type:"structure"`
}

type PipelineIDName struct {
	ID   *string `locationName:"id" type:"string"`
	Name *string `locationName:"name" type:"string"`

	metadataPipelineIDName `json:"-", xml:"-"`
}

type metadataPipelineIDName struct {
	SDKShapeTraits bool `type:"structure"`
}

type PipelineObject struct {
	Fields []*Field `locationName:"fields" type:"list" required:"true"`
	ID     *string  `locationName:"id" type:"string" required:"true"`
	Name   *string  `locationName:"name" type:"string" required:"true"`

	metadataPipelineObject `json:"-", xml:"-"`
}

type metadataPipelineObject struct {
	SDKShapeTraits bool `type:"structure"`
}

type PollForTaskInput struct {
	Hostname         *string           `locationName:"hostname" type:"string"`
	InstanceIdentity *InstanceIdentity `locationName:"instanceIdentity" type:"structure"`
	WorkerGroup      *string           `locationName:"workerGroup" type:"string" required:"true"`

	metadataPollForTaskInput `json:"-", xml:"-"`
}

type metadataPollForTaskInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PollForTaskOutput struct {
	TaskObject *TaskObject `locationName:"taskObject" type:"structure"`

	metadataPollForTaskOutput `json:"-", xml:"-"`
}

type metadataPollForTaskOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutPipelineDefinitionInput struct {
	ParameterObjects []*ParameterObject `locationName:"parameterObjects" type:"list"`
	ParameterValues  []*ParameterValue  `locationName:"parameterValues" type:"list"`
	PipelineID       *string            `locationName:"pipelineId" type:"string" required:"true"`
	PipelineObjects  []*PipelineObject  `locationName:"pipelineObjects" type:"list" required:"true"`

	metadataPutPipelineDefinitionInput `json:"-", xml:"-"`
}

type metadataPutPipelineDefinitionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutPipelineDefinitionOutput struct {
	Errored            *bool                `locationName:"errored" type:"boolean" required:"true"`
	ValidationErrors   []*ValidationError   `locationName:"validationErrors" type:"list"`
	ValidationWarnings []*ValidationWarning `locationName:"validationWarnings" type:"list"`

	metadataPutPipelineDefinitionOutput `json:"-", xml:"-"`
}

type metadataPutPipelineDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Query struct {
	Selectors []*Selector `locationName:"selectors" type:"list"`

	metadataQuery `json:"-", xml:"-"`
}

type metadataQuery struct {
	SDKShapeTraits bool `type:"structure"`
}

type QueryObjectsInput struct {
	Limit      *int64  `locationName:"limit" type:"integer"`
	Marker     *string `locationName:"marker" type:"string"`
	PipelineID *string `locationName:"pipelineId" type:"string" required:"true"`
	Query      *Query  `locationName:"query" type:"structure"`
	Sphere     *string `locationName:"sphere" type:"string" required:"true"`

	metadataQueryObjectsInput `json:"-", xml:"-"`
}

type metadataQueryObjectsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type QueryObjectsOutput struct {
	HasMoreResults *bool     `locationName:"hasMoreResults" type:"boolean"`
	IDs            []*string `locationName:"ids" type:"list"`
	Marker         *string   `locationName:"marker" type:"string"`

	metadataQueryObjectsOutput `json:"-", xml:"-"`
}

type metadataQueryObjectsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ReportTaskProgressInput struct {
	Fields []*Field `locationName:"fields" type:"list"`
	TaskID *string  `locationName:"taskId" type:"string" required:"true"`

	metadataReportTaskProgressInput `json:"-", xml:"-"`
}

type metadataReportTaskProgressInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ReportTaskProgressOutput struct {
	Canceled *bool `locationName:"canceled" type:"boolean" required:"true"`

	metadataReportTaskProgressOutput `json:"-", xml:"-"`
}

type metadataReportTaskProgressOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ReportTaskRunnerHeartbeatInput struct {
	Hostname     *string `locationName:"hostname" type:"string"`
	TaskRunnerID *string `locationName:"taskrunnerId" type:"string" required:"true"`
	WorkerGroup  *string `locationName:"workerGroup" type:"string"`

	metadataReportTaskRunnerHeartbeatInput `json:"-", xml:"-"`
}

type metadataReportTaskRunnerHeartbeatInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ReportTaskRunnerHeartbeatOutput struct {
	Terminate *bool `locationName:"terminate" type:"boolean" required:"true"`

	metadataReportTaskRunnerHeartbeatOutput `json:"-", xml:"-"`
}

type metadataReportTaskRunnerHeartbeatOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Selector struct {
	FieldName *string   `locationName:"fieldName" type:"string"`
	Operator  *Operator `locationName:"operator" type:"structure"`

	metadataSelector `json:"-", xml:"-"`
}

type metadataSelector struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetStatusInput struct {
	ObjectIDs  []*string `locationName:"objectIds" type:"list" required:"true"`
	PipelineID *string   `locationName:"pipelineId" type:"string" required:"true"`
	Status     *string   `locationName:"status" type:"string" required:"true"`

	metadataSetStatusInput `json:"-", xml:"-"`
}

type metadataSetStatusInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetStatusOutput struct {
	metadataSetStatusOutput `json:"-", xml:"-"`
}

type metadataSetStatusOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetTaskStatusInput struct {
	ErrorID         *string `locationName:"errorId" type:"string"`
	ErrorMessage    *string `locationName:"errorMessage" type:"string"`
	ErrorStackTrace *string `locationName:"errorStackTrace" type:"string"`
	TaskID          *string `locationName:"taskId" type:"string" required:"true"`
	TaskStatus      *string `locationName:"taskStatus" type:"string" required:"true"`

	metadataSetTaskStatusInput `json:"-", xml:"-"`
}

type metadataSetTaskStatusInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetTaskStatusOutput struct {
	metadataSetTaskStatusOutput `json:"-", xml:"-"`
}

type metadataSetTaskStatusOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type TaskObject struct {
	AttemptID  *string                     `locationName:"attemptId" type:"string"`
	Objects    *map[string]*PipelineObject `locationName:"objects" type:"map"`
	PipelineID *string                     `locationName:"pipelineId" type:"string"`
	TaskID     *string                     `locationName:"taskId" type:"string"`

	metadataTaskObject `json:"-", xml:"-"`
}

type metadataTaskObject struct {
	SDKShapeTraits bool `type:"structure"`
}

type ValidatePipelineDefinitionInput struct {
	ParameterObjects []*ParameterObject `locationName:"parameterObjects" type:"list"`
	ParameterValues  []*ParameterValue  `locationName:"parameterValues" type:"list"`
	PipelineID       *string            `locationName:"pipelineId" type:"string" required:"true"`
	PipelineObjects  []*PipelineObject  `locationName:"pipelineObjects" type:"list" required:"true"`

	metadataValidatePipelineDefinitionInput `json:"-", xml:"-"`
}

type metadataValidatePipelineDefinitionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ValidatePipelineDefinitionOutput struct {
	Errored            *bool                `locationName:"errored" type:"boolean" required:"true"`
	ValidationErrors   []*ValidationError   `locationName:"validationErrors" type:"list"`
	ValidationWarnings []*ValidationWarning `locationName:"validationWarnings" type:"list"`

	metadataValidatePipelineDefinitionOutput `json:"-", xml:"-"`
}

type metadataValidatePipelineDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ValidationError struct {
	Errors []*string `locationName:"errors" type:"list"`
	ID     *string   `locationName:"id" type:"string"`

	metadataValidationError `json:"-", xml:"-"`
}

type metadataValidationError struct {
	SDKShapeTraits bool `type:"structure"`
}

type ValidationWarning struct {
	ID       *string   `locationName:"id" type:"string"`
	Warnings []*string `locationName:"warnings" type:"list"`

	metadataValidationWarning `json:"-", xml:"-"`
}

type metadataValidationWarning struct {
	SDKShapeTraits bool `type:"structure"`
}