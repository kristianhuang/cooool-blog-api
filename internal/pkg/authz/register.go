/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package authz

import "cooool-blog-api/pkg/scheme"

const GroupName = "blog.authz"

// SchemeGroupVersion is group version used to register these objects.
var SchemeGroupVersion = scheme.GroupVersion{Group: GroupName, Version: "v1"}

// Resource takes an unqualified resource and returns a Group qualified GroupResource.
func Resource(resource string) scheme.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
