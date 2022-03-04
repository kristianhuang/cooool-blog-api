/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package scheme

type ObjectKind interface {
	SetGroupVersionKind()
	GroupVersionKind()
}

var EmptyObjectKind = emptyObjectKind{}

type emptyObjectKind struct{}

// SetGroupVersionKind implements the ObjectKind interface.
func (emptyObjectKind) SetGroupVersionKind(gvk GroupVersionKind) {}

// GroupVersionKind implements the ObjectKind interface.
func (emptyObjectKind) GroupVersionKind() GroupVersionKind { return GroupVersionKind{} }
