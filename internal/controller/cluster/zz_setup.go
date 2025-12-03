// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	sqssubscriber "github.com/philips-software/provider-hsdp/internal/controller/cluster/dbs/sqssubscriber"
	subscription "github.com/philips-software/provider-hsdp/internal/controller/cluster/dbs/subscription"
	application "github.com/philips-software/provider-hsdp/internal/controller/cluster/iam/application"
	client "github.com/philips-software/provider-hsdp/internal/controller/cluster/iam/client"
	emailtemplate "github.com/philips-software/provider-hsdp/internal/controller/cluster/iam/emailtemplate"
	group "github.com/philips-software/provider-hsdp/internal/controller/cluster/iam/group"
	organization "github.com/philips-software/provider-hsdp/internal/controller/cluster/iam/organization"
	passwordpolicy "github.com/philips-software/provider-hsdp/internal/controller/cluster/iam/passwordpolicy"
	proposition "github.com/philips-software/provider-hsdp/internal/controller/cluster/iam/proposition"
	role "github.com/philips-software/provider-hsdp/internal/controller/cluster/iam/role"
	rolesharingpolicy "github.com/philips-software/provider-hsdp/internal/controller/cluster/iam/rolesharingpolicy"
	service "github.com/philips-software/provider-hsdp/internal/controller/cluster/iam/service"
	user "github.com/philips-software/provider-hsdp/internal/controller/cluster/iam/user"
	datatype "github.com/philips-software/provider-hsdp/internal/controller/cluster/mdm/datatype"
	propositionmdm "github.com/philips-software/provider-hsdp/internal/controller/cluster/mdm/proposition"
	providerconfig "github.com/philips-software/provider-hsdp/internal/controller/cluster/providerconfig"
	key "github.com/philips-software/provider-hsdp/internal/controller/cluster/tenant/key"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		sqssubscriber.Setup,
		subscription.Setup,
		application.Setup,
		client.Setup,
		emailtemplate.Setup,
		group.Setup,
		organization.Setup,
		passwordpolicy.Setup,
		proposition.Setup,
		role.Setup,
		rolesharingpolicy.Setup,
		service.Setup,
		user.Setup,
		datatype.Setup,
		propositionmdm.Setup,
		providerconfig.Setup,
		key.Setup,
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
		sqssubscriber.SetupGated,
		subscription.SetupGated,
		application.SetupGated,
		client.SetupGated,
		emailtemplate.SetupGated,
		group.SetupGated,
		organization.SetupGated,
		passwordpolicy.SetupGated,
		proposition.SetupGated,
		role.SetupGated,
		rolesharingpolicy.SetupGated,
		service.SetupGated,
		user.SetupGated,
		datatype.SetupGated,
		propositionmdm.SetupGated,
		providerconfig.SetupGated,
		key.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
