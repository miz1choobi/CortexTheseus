// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the resource record sets in a specified hosted zone.
// ListResourceRecordSets returns up to 300 resource record sets at a time in ASCII
// order, beginning at a position specified by the name and type elements. Sort
// order ListResourceRecordSets sorts results first by DNS name with the labels
// reversed, for example: com.example.www. Note the trailing dot, which can change
// the sort order when the record name contains characters that appear before .
// (decimal 46) in the ASCII table. These characters include the following: ! " #
// $ % & ' ( ) * + , - When multiple records have the same DNS name,
// ListResourceRecordSets sorts results by the record type. Specifying where to
// start listing records You can use the name and type elements to specify the
// resource record set that the list begins with: If you do not specify Name or
// Type The results begin with the first resource record set that the hosted zone
// contains. If you specify Name but not Type The results begin with the first
// resource record set in the list whose name is greater than or equal to Name . If
// you specify Type but not Name Amazon Route 53 returns the InvalidInput error.
// If you specify both Name and Type The results begin with the first resource
// record set in the list whose name is greater than or equal to Name , and whose
// type is greater than or equal to Type . Resource record sets that are PENDING
// This action returns the most current version of the records. This includes
// records that are PENDING , and that are not yet available on all Route 53 DNS
// servers. Changing resource record sets To ensure that you get an accurate
// listing of the resource record sets for a hosted zone at a point in time, do not
// submit a ChangeResourceRecordSets request while you're paging through the
// results of a ListResourceRecordSets request. If you do, some pages may display
// results without the latest changes while other pages display results with the
// latest changes. Displaying the next page of results If a ListResourceRecordSets
// command returns more than one page of results, the value of IsTruncated is true
// . To display the next page of results, get the values of NextRecordName ,
// NextRecordType , and NextRecordIdentifier (if any) from the response. Then
// submit another ListResourceRecordSets request, and specify those values for
// StartRecordName , StartRecordType , and StartRecordIdentifier .
func (c *Client) ListResourceRecordSets(ctx context.Context, params *ListResourceRecordSetsInput, optFns ...func(*Options)) (*ListResourceRecordSetsOutput, error) {
	if params == nil {
		params = &ListResourceRecordSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceRecordSets", params, optFns, c.addOperationListResourceRecordSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceRecordSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for the resource record sets that are associated with a specified
// hosted zone.
type ListResourceRecordSetsInput struct {

	// The ID of the hosted zone that contains the resource record sets that you want
	// to list.
	//
	// This member is required.
	HostedZoneId *string

	// (Optional) The maximum number of resource records sets to include in the
	// response body for this request. If the response includes more than maxitems
	// resource record sets, the value of the IsTruncated element in the response is
	// true , and the values of the NextRecordName and NextRecordType elements in the
	// response identify the first resource record set in the next group of maxitems
	// resource record sets.
	MaxItems *int32

	// Resource record sets that have a routing policy other than simple: If results
	// were truncated for a given DNS name and type, specify the value of
	// NextRecordIdentifier from the previous response to get the next resource record
	// set that has the current DNS name and type.
	StartRecordIdentifier *string

	// The first name in the lexicographic ordering of resource record sets that you
	// want to list. If the specified record name doesn't exist, the results begin with
	// the first resource record set that has a name greater than the value of name .
	StartRecordName *string

	// The type of resource record set to begin the record listing from. Valid values
	// for basic resource record sets: A | AAAA | CAA | CNAME | MX | NAPTR | NS | PTR
	// | SOA | SPF | SRV | TXT Values for weighted, latency, geolocation, and failover
	// resource record sets: A | AAAA | CAA | CNAME | MX | NAPTR | PTR | SPF | SRV |
	// TXT Values for alias resource record sets:
	//   - API Gateway custom regional API or edge-optimized API: A
	//   - CloudFront distribution: A or AAAA
	//   - Elastic Beanstalk environment that has a regionalized subdomain: A
	//   - Elastic Load Balancing load balancer: A | AAAA
	//   - S3 bucket: A
	//   - VPC interface VPC endpoint: A
	//   - Another resource record set in this hosted zone: The type of the resource
	//   record set that the alias references.
	// Constraint: Specifying type without specifying name returns an InvalidInput
	// error.
	StartRecordType types.RRType

	noSmithyDocumentSerde
}

// A complex type that contains list information for the resource record set.
type ListResourceRecordSetsOutput struct {

	// A flag that indicates whether more resource record sets remain to be listed. If
	// your results were truncated, you can make a follow-up pagination request by
	// using the NextRecordName element.
	//
	// This member is required.
	IsTruncated bool

	// The maximum number of records you requested.
	//
	// This member is required.
	MaxItems *int32

	// Information about multiple resource record sets.
	//
	// This member is required.
	ResourceRecordSets []types.ResourceRecordSet

	// Resource record sets that have a routing policy other than simple: If results
	// were truncated for a given DNS name and type, the value of SetIdentifier for
	// the next resource record set that has the current DNS name and type. For
	// information about routing policies, see Choosing a Routing Policy (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html)
	// in the Amazon Route 53 Developer Guide.
	NextRecordIdentifier *string

	// If the results were truncated, the name of the next record in the list. This
	// element is present only if IsTruncated is true.
	NextRecordName *string

	// If the results were truncated, the type of the next record in the list. This
	// element is present only if IsTruncated is true.
	NextRecordType types.RRType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourceRecordSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListResourceRecordSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListResourceRecordSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResourceRecordSets"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpListResourceRecordSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceRecordSets(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
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

func newServiceMetadataMiddleware_opListResourceRecordSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResourceRecordSets",
	}
}
