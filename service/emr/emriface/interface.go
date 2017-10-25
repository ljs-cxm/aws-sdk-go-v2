// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package emriface provides an interface to enable mocking the Amazon Elastic MapReduce service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package emriface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
)

// EMRAPI provides an interface to enable mocking the
// emr.EMR service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Elastic MapReduce.
//    func myFunc(svc emriface.EMRAPI) bool {
//        // Make svc.AddInstanceFleet request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := emr.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockEMRClient struct {
//        emriface.EMRAPI
//    }
//    func (m *mockEMRClient) AddInstanceFleet(input *emr.AddInstanceFleetInput) (*emr.AddInstanceFleetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockEMRClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type EMRAPI interface {
	AddInstanceFleet(*emr.AddInstanceFleetInput) (*emr.AddInstanceFleetOutput, error)
	AddInstanceFleetWithContext(aws.Context, *emr.AddInstanceFleetInput, ...aws.Option) (*emr.AddInstanceFleetOutput, error)
	AddInstanceFleetRequest(*emr.AddInstanceFleetInput) (*aws.Request, *emr.AddInstanceFleetOutput)

	AddInstanceGroups(*emr.AddInstanceGroupsInput) (*emr.AddInstanceGroupsOutput, error)
	AddInstanceGroupsWithContext(aws.Context, *emr.AddInstanceGroupsInput, ...aws.Option) (*emr.AddInstanceGroupsOutput, error)
	AddInstanceGroupsRequest(*emr.AddInstanceGroupsInput) (*aws.Request, *emr.AddInstanceGroupsOutput)

	AddJobFlowSteps(*emr.AddJobFlowStepsInput) (*emr.AddJobFlowStepsOutput, error)
	AddJobFlowStepsWithContext(aws.Context, *emr.AddJobFlowStepsInput, ...aws.Option) (*emr.AddJobFlowStepsOutput, error)
	AddJobFlowStepsRequest(*emr.AddJobFlowStepsInput) (*aws.Request, *emr.AddJobFlowStepsOutput)

	AddTags(*emr.AddTagsInput) (*emr.AddTagsOutput, error)
	AddTagsWithContext(aws.Context, *emr.AddTagsInput, ...aws.Option) (*emr.AddTagsOutput, error)
	AddTagsRequest(*emr.AddTagsInput) (*aws.Request, *emr.AddTagsOutput)

	CancelSteps(*emr.CancelStepsInput) (*emr.CancelStepsOutput, error)
	CancelStepsWithContext(aws.Context, *emr.CancelStepsInput, ...aws.Option) (*emr.CancelStepsOutput, error)
	CancelStepsRequest(*emr.CancelStepsInput) (*aws.Request, *emr.CancelStepsOutput)

	CreateSecurityConfiguration(*emr.CreateSecurityConfigurationInput) (*emr.CreateSecurityConfigurationOutput, error)
	CreateSecurityConfigurationWithContext(aws.Context, *emr.CreateSecurityConfigurationInput, ...aws.Option) (*emr.CreateSecurityConfigurationOutput, error)
	CreateSecurityConfigurationRequest(*emr.CreateSecurityConfigurationInput) (*aws.Request, *emr.CreateSecurityConfigurationOutput)

	DeleteSecurityConfiguration(*emr.DeleteSecurityConfigurationInput) (*emr.DeleteSecurityConfigurationOutput, error)
	DeleteSecurityConfigurationWithContext(aws.Context, *emr.DeleteSecurityConfigurationInput, ...aws.Option) (*emr.DeleteSecurityConfigurationOutput, error)
	DeleteSecurityConfigurationRequest(*emr.DeleteSecurityConfigurationInput) (*aws.Request, *emr.DeleteSecurityConfigurationOutput)

	DescribeCluster(*emr.DescribeClusterInput) (*emr.DescribeClusterOutput, error)
	DescribeClusterWithContext(aws.Context, *emr.DescribeClusterInput, ...aws.Option) (*emr.DescribeClusterOutput, error)
	DescribeClusterRequest(*emr.DescribeClusterInput) (*aws.Request, *emr.DescribeClusterOutput)

	DescribeJobFlows(*emr.DescribeJobFlowsInput) (*emr.DescribeJobFlowsOutput, error)
	DescribeJobFlowsWithContext(aws.Context, *emr.DescribeJobFlowsInput, ...aws.Option) (*emr.DescribeJobFlowsOutput, error)
	DescribeJobFlowsRequest(*emr.DescribeJobFlowsInput) (*aws.Request, *emr.DescribeJobFlowsOutput)

	DescribeSecurityConfiguration(*emr.DescribeSecurityConfigurationInput) (*emr.DescribeSecurityConfigurationOutput, error)
	DescribeSecurityConfigurationWithContext(aws.Context, *emr.DescribeSecurityConfigurationInput, ...aws.Option) (*emr.DescribeSecurityConfigurationOutput, error)
	DescribeSecurityConfigurationRequest(*emr.DescribeSecurityConfigurationInput) (*aws.Request, *emr.DescribeSecurityConfigurationOutput)

	DescribeStep(*emr.DescribeStepInput) (*emr.DescribeStepOutput, error)
	DescribeStepWithContext(aws.Context, *emr.DescribeStepInput, ...aws.Option) (*emr.DescribeStepOutput, error)
	DescribeStepRequest(*emr.DescribeStepInput) (*aws.Request, *emr.DescribeStepOutput)

	ListBootstrapActions(*emr.ListBootstrapActionsInput) (*emr.ListBootstrapActionsOutput, error)
	ListBootstrapActionsWithContext(aws.Context, *emr.ListBootstrapActionsInput, ...aws.Option) (*emr.ListBootstrapActionsOutput, error)
	ListBootstrapActionsRequest(*emr.ListBootstrapActionsInput) (*aws.Request, *emr.ListBootstrapActionsOutput)

	ListBootstrapActionsPages(*emr.ListBootstrapActionsInput, func(*emr.ListBootstrapActionsOutput, bool) bool) error
	ListBootstrapActionsPagesWithContext(aws.Context, *emr.ListBootstrapActionsInput, func(*emr.ListBootstrapActionsOutput, bool) bool, ...aws.Option) error

	ListClusters(*emr.ListClustersInput) (*emr.ListClustersOutput, error)
	ListClustersWithContext(aws.Context, *emr.ListClustersInput, ...aws.Option) (*emr.ListClustersOutput, error)
	ListClustersRequest(*emr.ListClustersInput) (*aws.Request, *emr.ListClustersOutput)

	ListClustersPages(*emr.ListClustersInput, func(*emr.ListClustersOutput, bool) bool) error
	ListClustersPagesWithContext(aws.Context, *emr.ListClustersInput, func(*emr.ListClustersOutput, bool) bool, ...aws.Option) error

	ListInstanceFleets(*emr.ListInstanceFleetsInput) (*emr.ListInstanceFleetsOutput, error)
	ListInstanceFleetsWithContext(aws.Context, *emr.ListInstanceFleetsInput, ...aws.Option) (*emr.ListInstanceFleetsOutput, error)
	ListInstanceFleetsRequest(*emr.ListInstanceFleetsInput) (*aws.Request, *emr.ListInstanceFleetsOutput)

	ListInstanceFleetsPages(*emr.ListInstanceFleetsInput, func(*emr.ListInstanceFleetsOutput, bool) bool) error
	ListInstanceFleetsPagesWithContext(aws.Context, *emr.ListInstanceFleetsInput, func(*emr.ListInstanceFleetsOutput, bool) bool, ...aws.Option) error

	ListInstanceGroups(*emr.ListInstanceGroupsInput) (*emr.ListInstanceGroupsOutput, error)
	ListInstanceGroupsWithContext(aws.Context, *emr.ListInstanceGroupsInput, ...aws.Option) (*emr.ListInstanceGroupsOutput, error)
	ListInstanceGroupsRequest(*emr.ListInstanceGroupsInput) (*aws.Request, *emr.ListInstanceGroupsOutput)

	ListInstanceGroupsPages(*emr.ListInstanceGroupsInput, func(*emr.ListInstanceGroupsOutput, bool) bool) error
	ListInstanceGroupsPagesWithContext(aws.Context, *emr.ListInstanceGroupsInput, func(*emr.ListInstanceGroupsOutput, bool) bool, ...aws.Option) error

	ListInstances(*emr.ListInstancesInput) (*emr.ListInstancesOutput, error)
	ListInstancesWithContext(aws.Context, *emr.ListInstancesInput, ...aws.Option) (*emr.ListInstancesOutput, error)
	ListInstancesRequest(*emr.ListInstancesInput) (*aws.Request, *emr.ListInstancesOutput)

	ListInstancesPages(*emr.ListInstancesInput, func(*emr.ListInstancesOutput, bool) bool) error
	ListInstancesPagesWithContext(aws.Context, *emr.ListInstancesInput, func(*emr.ListInstancesOutput, bool) bool, ...aws.Option) error

	ListSecurityConfigurations(*emr.ListSecurityConfigurationsInput) (*emr.ListSecurityConfigurationsOutput, error)
	ListSecurityConfigurationsWithContext(aws.Context, *emr.ListSecurityConfigurationsInput, ...aws.Option) (*emr.ListSecurityConfigurationsOutput, error)
	ListSecurityConfigurationsRequest(*emr.ListSecurityConfigurationsInput) (*aws.Request, *emr.ListSecurityConfigurationsOutput)

	ListSteps(*emr.ListStepsInput) (*emr.ListStepsOutput, error)
	ListStepsWithContext(aws.Context, *emr.ListStepsInput, ...aws.Option) (*emr.ListStepsOutput, error)
	ListStepsRequest(*emr.ListStepsInput) (*aws.Request, *emr.ListStepsOutput)

	ListStepsPages(*emr.ListStepsInput, func(*emr.ListStepsOutput, bool) bool) error
	ListStepsPagesWithContext(aws.Context, *emr.ListStepsInput, func(*emr.ListStepsOutput, bool) bool, ...aws.Option) error

	ModifyInstanceFleet(*emr.ModifyInstanceFleetInput) (*emr.ModifyInstanceFleetOutput, error)
	ModifyInstanceFleetWithContext(aws.Context, *emr.ModifyInstanceFleetInput, ...aws.Option) (*emr.ModifyInstanceFleetOutput, error)
	ModifyInstanceFleetRequest(*emr.ModifyInstanceFleetInput) (*aws.Request, *emr.ModifyInstanceFleetOutput)

	ModifyInstanceGroups(*emr.ModifyInstanceGroupsInput) (*emr.ModifyInstanceGroupsOutput, error)
	ModifyInstanceGroupsWithContext(aws.Context, *emr.ModifyInstanceGroupsInput, ...aws.Option) (*emr.ModifyInstanceGroupsOutput, error)
	ModifyInstanceGroupsRequest(*emr.ModifyInstanceGroupsInput) (*aws.Request, *emr.ModifyInstanceGroupsOutput)

	PutAutoScalingPolicy(*emr.PutAutoScalingPolicyInput) (*emr.PutAutoScalingPolicyOutput, error)
	PutAutoScalingPolicyWithContext(aws.Context, *emr.PutAutoScalingPolicyInput, ...aws.Option) (*emr.PutAutoScalingPolicyOutput, error)
	PutAutoScalingPolicyRequest(*emr.PutAutoScalingPolicyInput) (*aws.Request, *emr.PutAutoScalingPolicyOutput)

	RemoveAutoScalingPolicy(*emr.RemoveAutoScalingPolicyInput) (*emr.RemoveAutoScalingPolicyOutput, error)
	RemoveAutoScalingPolicyWithContext(aws.Context, *emr.RemoveAutoScalingPolicyInput, ...aws.Option) (*emr.RemoveAutoScalingPolicyOutput, error)
	RemoveAutoScalingPolicyRequest(*emr.RemoveAutoScalingPolicyInput) (*aws.Request, *emr.RemoveAutoScalingPolicyOutput)

	RemoveTags(*emr.RemoveTagsInput) (*emr.RemoveTagsOutput, error)
	RemoveTagsWithContext(aws.Context, *emr.RemoveTagsInput, ...aws.Option) (*emr.RemoveTagsOutput, error)
	RemoveTagsRequest(*emr.RemoveTagsInput) (*aws.Request, *emr.RemoveTagsOutput)

	RunJobFlow(*emr.RunJobFlowInput) (*emr.RunJobFlowOutput, error)
	RunJobFlowWithContext(aws.Context, *emr.RunJobFlowInput, ...aws.Option) (*emr.RunJobFlowOutput, error)
	RunJobFlowRequest(*emr.RunJobFlowInput) (*aws.Request, *emr.RunJobFlowOutput)

	SetTerminationProtection(*emr.SetTerminationProtectionInput) (*emr.SetTerminationProtectionOutput, error)
	SetTerminationProtectionWithContext(aws.Context, *emr.SetTerminationProtectionInput, ...aws.Option) (*emr.SetTerminationProtectionOutput, error)
	SetTerminationProtectionRequest(*emr.SetTerminationProtectionInput) (*aws.Request, *emr.SetTerminationProtectionOutput)

	SetVisibleToAllUsers(*emr.SetVisibleToAllUsersInput) (*emr.SetVisibleToAllUsersOutput, error)
	SetVisibleToAllUsersWithContext(aws.Context, *emr.SetVisibleToAllUsersInput, ...aws.Option) (*emr.SetVisibleToAllUsersOutput, error)
	SetVisibleToAllUsersRequest(*emr.SetVisibleToAllUsersInput) (*aws.Request, *emr.SetVisibleToAllUsersOutput)

	TerminateJobFlows(*emr.TerminateJobFlowsInput) (*emr.TerminateJobFlowsOutput, error)
	TerminateJobFlowsWithContext(aws.Context, *emr.TerminateJobFlowsInput, ...aws.Option) (*emr.TerminateJobFlowsOutput, error)
	TerminateJobFlowsRequest(*emr.TerminateJobFlowsInput) (*aws.Request, *emr.TerminateJobFlowsOutput)

	WaitUntilClusterRunning(*emr.DescribeClusterInput) error
	WaitUntilClusterRunningWithContext(aws.Context, *emr.DescribeClusterInput, ...aws.WaiterOption) error

	WaitUntilClusterTerminated(*emr.DescribeClusterInput) error
	WaitUntilClusterTerminatedWithContext(aws.Context, *emr.DescribeClusterInput, ...aws.WaiterOption) error

	WaitUntilStepComplete(*emr.DescribeStepInput) error
	WaitUntilStepCompleteWithContext(aws.Context, *emr.DescribeStepInput, ...aws.WaiterOption) error
}

var _ EMRAPI = (*emr.EMR)(nil)