/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package v1

import (
	"context"

	"cooool-blog-api/internal/apiserver/store"
	"cooool-blog-api/internal/pkg/code"
	v1 "cooool-blog-api/internal/pkg/model"
	"cooool-blog-api/pkg/errors"
	metav1 "cooool-blog-api/pkg/meta/v1"
)

// PolicySrv defines functions used to handle policy request.
type PolicyService interface {
	Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error
	Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, name string, opts metav1.GetOptions) (*v1.Policy, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error)
}

type policyService struct {
	store store.Factory
}

var _ PolicyService = (*policyService)(nil)

func newPolicies(s store.Factory) *policyService {
	return &policyService{store: s}
}

func (p *policyService) Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error {
	if err := p.store.Policies().Create(ctx, policy, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (p *policyService) Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error {
	if err := p.store.Policies().Update(ctx, policy, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (p *policyService) Delete(ctx context.Context, username, name string, opts metav1.DeleteOptions) error {
	if err := p.store.Policies().Delete(ctx, username, name, opts); err != nil {
		return err
	}

	return nil
}

func (p *policyService) DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error {
	if err := p.store.Policies().DeleteCollection(ctx, username, names, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (p *policyService) Get(ctx context.Context, username, name string, opts metav1.GetOptions) (*v1.Policy, error) {
	policy, err := p.store.Policies().Get(ctx, username, name, opts)
	if err != nil {
		return nil, err
	}

	return policy, nil
}

func (p *policyService) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error) {
	policies, err := p.store.Policies().List(ctx, username, opts)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return policies, nil
}
