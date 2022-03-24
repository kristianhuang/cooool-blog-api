/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"context"

	pb "cooool-blog-api/internal/pkg/proto/apiserver/v1"
	"cooool-blog-api/pkg/errors"
	log "cooool-blog-api/pkg/rollinglog"

	"github.com/AlekSi/pointer"
	"github.com/avast/retry-go"
)

type secrets struct {
	cli pb.CacheClient
}

func newSecrets(ds *datastore) *secrets {
	return &secrets{ds.cli}
}

// List returns all the authorization secrets.
func (s *secrets) List() (map[string]*pb.SecretInfo, error) {
	secrets := make(map[string]*pb.SecretInfo)

	log.Info("Loading secrets")

	req := &pb.ListSecretsRequest{
		Offset: pointer.ToInt64(0),
		Limit:  pointer.ToInt64(-1),
	}

	var resp *pb.ListSecretsResponse
	err := retry.Do(
		func() error {
			var listErr error
			resp, listErr = s.cli.ListSecrets(context.Background(), req)
			if listErr != nil {
				return listErr
			}

			return nil
		}, retry.Attempts(3),
	)
	if err != nil {
		return nil, errors.Wrap(err, "list secrets failed")
	}

	log.Infof("Secrets found (%d total):", len(resp.Items))

	for _, v := range resp.Items {
		log.Infof(" - %s:%s", v.Username, v.SecretId)
		secrets[v.SecretId] = v
	}

	return secrets, nil
}
