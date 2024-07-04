package outputs

import (
	pulumicommonsloadbalancerservice "github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/kubernetes/loadbalancer/service"
	solrcontextstate "github.com/plantoncloud/solr-kubernetes-pulumi-blueprint/pkg/solr/contextstate"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	ResourceId                    string
	ResourceName                  string
	EnvironmentName               string
	EndpointDomainName            string
	NamespaceName                 string
	ExternalLoadBalancerIpAddress string
	InternalLoadBalancerIpAddress string
	InternalHostname              string
	ExternalHostname              string
	KubeServiceName               string
	KubeLocalEndpoint             string
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxConfig = ctx.Value(solrcontextstate.Key).(solrcontextstate.ContextState)
	var externalLoadBalancerIpAddress = ""
	var internalLoadBalancerIpAddress = ""

	if ctxConfig.Status.AddedResources.LoadBalancerExternalService != nil {
		externalLoadBalancerIpAddress = pulumicommonsloadbalancerservice.GetIpAddress(ctxConfig.Status.AddedResources.LoadBalancerExternalService)
	}

	if ctxConfig.Status.AddedResources.LoadBalancerInternalService != nil {
		internalLoadBalancerIpAddress = pulumicommonsloadbalancerservice.GetIpAddress(ctxConfig.Status.AddedResources.LoadBalancerExternalService)
	}

	return &input{
		ResourceId:                    ctxConfig.Spec.ResourceId,
		ResourceName:                  ctxConfig.Spec.ResourceName,
		EnvironmentName:               ctxConfig.Spec.EnvironmentInfo.EnvironmentName,
		EndpointDomainName:            ctxConfig.Spec.EndpointDomainName,
		NamespaceName:                 ctxConfig.Spec.NamespaceName,
		ExternalLoadBalancerIpAddress: externalLoadBalancerIpAddress,
		InternalLoadBalancerIpAddress: internalLoadBalancerIpAddress,
		InternalHostname:              ctxConfig.Spec.InternalHostname,
		ExternalHostname:              ctxConfig.Spec.ExternalHostname,
		KubeServiceName:               ctxConfig.Spec.KubeServiceName,
		KubeLocalEndpoint:             ctxConfig.Spec.KubeLocalEndpoint,
	}
}
