/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package store

var client Factory

type Factory interface {
	AdminUser() AdminUserStore
	Policies() PolicyStore
	PolicyAudits() PolicyAuditStore
	Secrets() SecretStore
	Close() error
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
