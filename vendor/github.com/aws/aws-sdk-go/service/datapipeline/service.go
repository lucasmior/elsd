// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package datapipeline

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// AWS Data Pipeline configures and manages a data-driven workflow called a
// pipeline. AWS Data Pipeline handles the details of scheduling and ensuring
// that data dependencies are met so that your application can focus on processing
// the data.
//
// AWS Data Pipeline provides a JAR implementation of a task runner called AWS
// Data Pipeline Task Runner. AWS Data Pipeline Task Runner provides logic for
// common data management scenarios, such as performing database queries and
// running data analysis using Amazon Elastic MapReduce (Amazon EMR). You can
// use AWS Data Pipeline Task Runner as your task runner, or you can write your
// own task runner to provide custom data management.
//
// AWS Data Pipeline implements two main sets of functionality. Use the first
// set to create a pipeline and define data sources, schedules, dependencies,
// and the transforms to be performed on the data. Use the second set in your
// task runner application to receive the next task ready for processing. The
// logic for performing the task, such as querying the data, running data analysis,
// or converting the data from one format to another, is contained within the
// task runner. The task runner performs the task assigned to it by the web
// service, reporting progress to the web service as it does so. When the task
// is done, the task runner reports the final success or failure of the task
// to the web service.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type DataPipeline struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "datapipeline"

// New creates a new instance of the DataPipeline client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a DataPipeline client from just a session.
//     svc := datapipeline.New(mySession)
//
//     // Create a DataPipeline client with additional configuration
//     svc := datapipeline.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *DataPipeline {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *DataPipeline {
	svc := &DataPipeline{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2012-10-29",
				JSONVersion:   "1.1",
				TargetPrefix:  "DataPipeline",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a DataPipeline operation and runs any
// custom request initialization.
func (c *DataPipeline) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
