/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package storage

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/google/uuid"
	"github.com/spaolacci/murmur3"
)

const (
	defaultHashAlgorithm = "murmur64"

	HashSha256    = "sha256"
	HashMurmur32  = "murmur32"
	HashMurmur64  = "murmur64"
	HashMurmur128 = "murmur128"

	B64JSONPrefix = "ey"
)

func GenerateToken(orgID, keyID, hashAlgorithm string) (string, error) {
	if keyID != "" {
		uuid.New()
		keyID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}

	if hashAlgorithm != "" {
		_, err := hashFunc(hashAlgorithm)
		if err != nil {
			hashAlgorithm = defaultHashAlgorithm
		}

		jsonToken := fmt.Sprintf(`{"org":"%s","id":"%s","h":"%s"}`, orgID, keyID, hashAlgorithm)

		return base64.StdEncoding.EncodeToString([]byte(jsonToken)), err
	}

	return orgID + keyID, nil
}

func TokenHashAlgo(token string) string {
	if strings.HasPrefix(token, B64JSONPrefix) {
		if jsonToken, err := base64.StdEncoding.DecodeString(token); err == nil {
			hahAlgo, _ := jsonparser.GetString(jsonToken, "h")
			return hahAlgo
		}
	}

	if len(token) > 24 {
		return token[:24]
	}

	return ""
}

func TokenOrg(token string) string {
	if strings.HasPrefix(token, B64JSONPrefix) {
		if jsonToken, err := base64.StdEncoding.DecodeString(token); err == nil {
			if org, err := jsonparser.GetString(jsonToken, "org"); err == nil {
				return org
			}
		}
	}

	return ""
}

func hashFunc(algorithm string) (hash.Hash, error) {
	switch algorithm {
	case HashSha256:
		return sha256.New(), nil
	case HashMurmur32:
		return murmur3.New32(), nil
	case HashMurmur64:
		return murmur3.New64(), nil
	case HashMurmur128:
		return murmur3.New128(), nil
	default:
		return murmur3.New32(), fmt.Errorf("unknown key hash function: %s. Falling back to murmur32", algorithm)
	}
}

func HashStr(in string) string {
	h, _ := hashFunc(TokenHashAlgo(in))
	h.Write([]byte(in))

	return hex.EncodeToString(h.Sum(nil))
}

func HashKey(in string) string {
	return HashStr(in)
}
