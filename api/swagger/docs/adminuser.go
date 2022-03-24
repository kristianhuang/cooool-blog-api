/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package docs

import metav1 "cooool-blog-api/pkg/meta/v1"

//	swagger:route GET /admin_users adminUserList
//
//	Get admin user list.
//
//  获取后台用户列表
//
//     Security:
//       api_key:
//     Responses: // 定义状态码关联的响应
//       default: errResponse
//       200: adminUserListResponse

// swagger:parameters adminUserList
type adminUserListRequestParamsWrapper struct {
	// in:query // 表示该参数所在的位置
	metav1.ListOptions
}
