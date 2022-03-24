//go:build !jsoniter
// +build !jsoniter

/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package json

import "encoding/json"

// RawMessage is exported by blog-app/pkg/json package.
type RawMessage = json.RawMessage

var (
	// Marshal is exported by cooool-blog-api/pkg/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by blog-app/pkg/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by blog-app/pkg/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by blog-app/pkg/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by blog-app/pkg/json package.
	NewEncoder = json.NewEncoder
)
