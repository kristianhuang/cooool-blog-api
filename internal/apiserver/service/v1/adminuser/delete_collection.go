/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"context"

	"blog-api/internal/pkg/code"
	"blog-api/pkg/errors"
	metav1 "blog-api/pkg/meta/v1"
)

func (a *adminUserService) DeleteCollection(ctx context.Context, accounts []string, opts metav1.DeleteOptions) error {
	if err := a.store.AdminUser().DeleteCollection(ctx, accounts, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}
