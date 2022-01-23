/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package dir

import (
	"os"
	"path/filepath"
	"runtime"
)

func HomeDir() string {
	if runtime.GOOS != "windows" {
		return os.Getenv("HOME")
	}
	home := os.Getenv("HOME")
	homeDriveHomePath := ""
	if homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"); len(homeDrive) > 0 && len(homePath) > 0 {
		homeDriveHomePath = homeDrive + homePath
	}
	userProfile := os.Getenv("USERPROFILE")

	// Return first of %HOME%, %HOMEDRIVE%/%HOMEPATH%, %USERPROFILE% that contains a `.apimachinery\config` file.
	// %HOMEDRIVE%/%HOMEPATH% is preferred over %USERPROFILE% for backwards-compatibility.
	for _, p := range []string{home, homeDriveHomePath, userProfile} {
		if len(p) == 0 {
			continue
		}
		if _, err := os.Stat(filepath.Join(p, ".apimachinery", "config")); err != nil {
			continue
		}
		return p
	}

	firstSetPath := ""
	firstExistingPath := ""

	// Prefer %USERPROFILE% over %HOMEDRIVE%/%HOMEPATH% for compatibility with other auth-writing tools
	for _, p := range []string{home, userProfile, homeDriveHomePath} {
		if len(p) == 0 {
			continue
		}
		if len(firstSetPath) == 0 {
			// remember the first path that is set
			firstSetPath = p
		}
		info, err := os.Stat(p)
		if err != nil {
			continue
		}
		if len(firstExistingPath) == 0 {
			// remember the first path that exists
			firstExistingPath = p
		}
		if info.IsDir() && info.Mode().Perm()&(1<<(uint(7))) != 0 {
			// return first path that is writeable
			return p
		}
	}

	// If none are writeable, return first location that exists
	if len(firstExistingPath) > 0 {
		return firstExistingPath
	}

	// If none exist, return first location that is set
	if len(firstSetPath) > 0 {
		return firstSetPath
	}

	// We've got nothing
	return ""
}
