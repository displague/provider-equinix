// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cloudrouter "github.com/displague/provider-equinix/internal/controller/namespaced/fabric/cloudrouter"
	connection "github.com/displague/provider-equinix/internal/controller/namespaced/fabric/connection"
	network "github.com/displague/provider-equinix/internal/controller/namespaced/fabric/network"
	routingprotocol "github.com/displague/provider-equinix/internal/controller/namespaced/fabric/routingprotocol"
	serviceprofile "github.com/displague/provider-equinix/internal/controller/namespaced/fabric/serviceprofile"
	acltemplate "github.com/displague/provider-equinix/internal/controller/namespaced/network/acltemplate"
	bgp "github.com/displague/provider-equinix/internal/controller/namespaced/network/bgp"
	device "github.com/displague/provider-equinix/internal/controller/namespaced/network/device"
	devicelink "github.com/displague/provider-equinix/internal/controller/namespaced/network/devicelink"
	sshkey "github.com/displague/provider-equinix/internal/controller/namespaced/network/sshkey"
	sshuser "github.com/displague/provider-equinix/internal/controller/namespaced/network/sshuser"
	providerconfig "github.com/displague/provider-equinix/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cloudrouter.Setup,
		connection.Setup,
		network.Setup,
		routingprotocol.Setup,
		serviceprofile.Setup,
		acltemplate.Setup,
		bgp.Setup,
		device.Setup,
		devicelink.Setup,
		sshkey.Setup,
		sshuser.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cloudrouter.SetupGated,
		connection.SetupGated,
		network.SetupGated,
		routingprotocol.SetupGated,
		serviceprofile.SetupGated,
		acltemplate.SetupGated,
		bgp.SetupGated,
		device.SetupGated,
		devicelink.SetupGated,
		sshkey.SetupGated,
		sshuser.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
