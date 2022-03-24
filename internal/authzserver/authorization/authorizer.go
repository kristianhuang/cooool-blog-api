/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package authorization

import (
	"cooool-blog-api/internal/pkg/authz"

	"github.com/ory/ladon"
)

// Authorizer implement the authorize interface that use local repository to
// authorize the subject access review.
type Authorizer struct {
	warden ladon.Warden
}

// NewAuthorizer creates a local repository authorizer and returns it.
func NewAuthorizer(authorizationClient AuthorizationInterface) *Authorizer {
	return &Authorizer{
		warden: &ladon.Ladon{
			Manager:     NewPolicyManager(authorizationClient),
			AuditLogger: NewAuditLogger(authorizationClient),
		},
	}
}

// Authorize to determine the subject access.
func (a *Authorizer) Authorize(req *ladon.Request) *authz.Response {
	if err := a.warden.IsAllowed(req); err != nil {
		return &authz.Response{
			Denied: true,
			Reason: err.Error(),
		}
	}

	return &authz.Response{
		Allowed: true,
	}
}
