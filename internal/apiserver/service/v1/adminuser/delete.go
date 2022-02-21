/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"context"

	metav1 "blog-api/pkg/meta/v1"
)

func (a *adminUserService) Delete(ctx context.Context, account string, opts metav1.DeleteOptions) error {

	if err := a.store.AdminUser().Delete(ctx, account, opts); err != nil {
		return err
	}

	return nil
}
