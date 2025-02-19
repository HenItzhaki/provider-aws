/*
Copyright 2023 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vpc

import (
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/flowlog"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/internetgateway"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/natgateway"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/route"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/routetable"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/securitygroup"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/securitygrouprule"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/subnet"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/transitgateway"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/transitgatewayroute"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/transitgatewayroutetable"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/transitgatewayvpcattachment"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/vpc"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/vpccidrblock"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/vpcendpoint"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/vpcendpointserviceconfiguration"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/vpc/vpcpeeringconnection"
	"github.com/crossplane-contrib/provider-aws/pkg/utils/setup"
)

// Setup vpc controllers.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	return setup.SetupControllers(
		mgr, o,
		flowlog.SetupFlowLog,
		internetgateway.SetupInternetGateway,
		natgateway.SetupNatGateway,
		route.SetupRoute,
		routetable.SetupRouteTable,
		securitygroup.SetupSecurityGroup,
		securitygrouprule.SetupSecurityGroupRule,
		subnet.SetupSubnet,
		transitgateway.SetupTransitGateway,
		transitgatewayroute.SetupTransitGatewayRoute,
		transitgatewayroutetable.SetupTransitGatewayRouteTable,
		transitgatewayvpcattachment.SetupTransitGatewayVPCAttachment,
		vpc.SetupVPC,
		vpccidrblock.SetupVPCCIDRBlock,
		vpcendpoint.SetupVPCEndpoint,
		vpcendpointserviceconfiguration.SetupVPCEndpointServiceConfiguration,
		vpcpeeringconnection.SetupVPCPeeringConnection,
	)
}
