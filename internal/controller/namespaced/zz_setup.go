// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	sqssubscriber "github.com/philips-software/provider-hsdp/internal/controller/namespaced/dbs/sqssubscriber"
	subscription "github.com/philips-software/provider-hsdp/internal/controller/namespaced/dbs/subscription"
	application "github.com/philips-software/provider-hsdp/internal/controller/namespaced/iam/application"
	client "github.com/philips-software/provider-hsdp/internal/controller/namespaced/iam/client"
	emailtemplate "github.com/philips-software/provider-hsdp/internal/controller/namespaced/iam/emailtemplate"
	group "github.com/philips-software/provider-hsdp/internal/controller/namespaced/iam/group"
	organization "github.com/philips-software/provider-hsdp/internal/controller/namespaced/iam/organization"
	passwordpolicy "github.com/philips-software/provider-hsdp/internal/controller/namespaced/iam/passwordpolicy"
	proposition "github.com/philips-software/provider-hsdp/internal/controller/namespaced/iam/proposition"
	role "github.com/philips-software/provider-hsdp/internal/controller/namespaced/iam/role"
	rolesharingpolicy "github.com/philips-software/provider-hsdp/internal/controller/namespaced/iam/rolesharingpolicy"
	service "github.com/philips-software/provider-hsdp/internal/controller/namespaced/iam/service"
	user "github.com/philips-software/provider-hsdp/internal/controller/namespaced/iam/user"
	datatype "github.com/philips-software/provider-hsdp/internal/controller/namespaced/mdm/datatype"
	propositionmdm "github.com/philips-software/provider-hsdp/internal/controller/namespaced/mdm/proposition"
	providerconfig "github.com/philips-software/provider-hsdp/internal/controller/namespaced/providerconfig"
	key "github.com/philips-software/provider-hsdp/internal/controller/namespaced/tenant/key"
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
