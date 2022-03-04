/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package v1

import (
	"context"

	"blog-api/internal/apiserver/store"
	"blog-api/internal/pkg/code"
	"blog-api/internal/pkg/model"
	"blog-api/pkg/errors"
	metav1 "blog-api/pkg/meta/v1"
)

type SecretService interface {
	Create(ctx context.Context, secret *model.Secret, opts metav1.CreateOptions) error
	Update(ctx context.Context, secret *model.Secret, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username, secretID string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, secretIDs []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username, secretID string, opts metav1.GetOptions) (*model.Secret, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*model.SecretList, error)
}

type secretService struct {
	store store.Factory
}

func (s *secretService) Create(ctx context.Context, secret *model.Secret, opts metav1.CreateOptions) error {
	if err := s.store.Secrets().Create(ctx, secret, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (s *secretService) Update(ctx context.Context, secret *model.Secret, opts metav1.UpdateOptions) error {
	if err := s.store.Secrets().Update(ctx, secret, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (s *secretService) Delete(ctx context.Context, username, secretID string, opts metav1.DeleteOptions) error {
	if err := s.store.Secrets().Delete(ctx, username, secretID, opts); err != nil {
		return err
	}

	return nil
}

func (s *secretService) DeleteCollection(ctx context.Context, username string, secretIDs []string, opts metav1.DeleteOptions) error {
	if err := s.store.Secrets().DeleteCollection(ctx, username, secretIDs, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (s *secretService) Get(ctx context.Context, username, secretID string, opts metav1.GetOptions) (*model.Secret, error) {
	secret, err := s.store.Secrets().Get(ctx, username, secretID, opts)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (s *secretService) List(ctx context.Context, username string, opts metav1.ListOptions) (*model.SecretList, error) {
	secrets, err := s.store.Secrets().List(ctx, username, opts)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return secrets, nil
}

func newSecrets(srv *service) *secretService {
	return &secretService{store: srv.store}
}
