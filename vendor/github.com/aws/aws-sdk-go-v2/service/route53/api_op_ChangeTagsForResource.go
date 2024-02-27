// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds, edits, or deletes tags for a health check or a hosted zone. For
// information about using tags for cost allocation, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the Billing and Cost Management User Guide.
func (c *Client) ChangeTagsForResource(ctx context.Context, params *ChangeTagsForResourceInput, optFns ...func(*Options)) (*ChangeTagsForResourceOutput, error) {
	if params == nil {
		params = &ChangeTagsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChangeTagsForResource", params, optFns, c.addOperationChangeTagsForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChangeTagsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the tags that you want to add,
// edit, or delete.
type ChangeTagsForResourceInput struct {

	// The ID of the resource for which you want to add, change, or delete tags.
	//
	// This member is required.
	ResourceId *string

	// The type of the resource.
	//   - The resource type for health checks is healthcheck .
	//   - The resource type for hosted zones is hostedzone .
	//
	// This member is required.
	ResourceType types.TagResourceType

	// A complex type that contains a list of the tags that you want to add to the
	// specified health check or hosted zone and/or the tags that you want to edit
	// Value for. You can add a maximum of 10 tags to a health check or a hosted zone.
	AddTags []types.Tag

	// A complex type that contains a list of the tags that you want to delete from
	// the specified health check or hosted zone. You can specify up to 10 keys.
	RemoveTagKeys []string

	noSmithyDocumentSerde
}

// Empty response for the request.
type ChangeTagsForResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationChangeTagsForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpChangeTagsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpChangeTagsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ChangeTagsForResource"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpChangeTagsForResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opChangeTagsForResource(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opChangeTagsForResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ChangeTagsForResource",
	}
}
