// Code generated by "internal/generate/listpages/main.go -ListOps=ListTrafficPolicies -Paginator=TrafficPolicyIdMarker list_traffic_policies_pages_gen.go"; DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

func listTrafficPoliciesPages(conn *route53.Route53, input *route53.ListTrafficPoliciesInput, fn func(*route53.ListTrafficPoliciesOutput, bool) bool) error {
	return listTrafficPoliciesPagesWithContext(context.Background(), conn, input, fn)
}

func listTrafficPoliciesPagesWithContext(ctx context.Context, conn *route53.Route53, input *route53.ListTrafficPoliciesInput, fn func(*route53.ListTrafficPoliciesOutput, bool) bool) error {
	for {
		output, err := conn.ListTrafficPoliciesWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.TrafficPolicyIdMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.TrafficPolicyIdMarker = output.TrafficPolicyIdMarker
	}
	return nil
}
