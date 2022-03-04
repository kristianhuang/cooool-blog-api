/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package v1

import (
	"context"
	"regexp"
	"sync"

	"blog-api/internal/apiserver/store"
	"blog-api/internal/pkg/code"
	"blog-api/internal/pkg/model"
	"blog-api/pkg/errors"
	metav1 "blog-api/pkg/meta/v1"
	log "blog-api/pkg/rollinglog"
)

type AdminUserService interface {
	Create(ctx context.Context, au *model.AdminUser, options metav1.CreateOptions) error
	Update(ctx context.Context, adminUser *model.AdminUser, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*model.AdminUser, error)
	List(ctx context.Context, opts metav1.ListOptions) (*model.AdminUserList, error)
	ListWithBadPerformance(ctx context.Context, opts metav1.ListOptions) (*model.AdminUserList, error)
	ChangePassword(ctx context.Context, user *model.AdminUser) error
}

type adminUserService struct {
	store store.Factory
}

var _ AdminUserService = (*adminUserService)(nil)

func newAdminUserService(s store.Factory) *adminUserService {
	return &adminUserService{store: s}
}

func (a *adminUserService) Create(ctx context.Context, au *model.AdminUser, options metav1.CreateOptions) error {
	if err := a.store.AdminUser().Create(ctx, au, options); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errors.WithCode(code.ErrUserAlreadyExist, err.Error())
		}

		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'nickname'", err.Error()); match {
			return errors.WithCode(code.ErrUserAlreadyExist, err.Error())
		}

		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (a *adminUserService) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {

	if err := a.store.AdminUser().Delete(ctx, username, opts); err != nil {
		return err
	}

	return nil
}

func (a *adminUserService) DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	if err := a.store.AdminUser().DeleteCollection(ctx, usernames, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (a *adminUserService) Get(ctx context.Context, username string, opts metav1.GetOptions) (*model.AdminUser, error) {
	au, err := a.store.AdminUser().Get(ctx, username, opts)
	if err != nil {
		return nil, err
	}

	return au, nil
}

func (a *adminUserService) List(ctx context.Context, opts metav1.ListOptions) (*model.AdminUserList, error) {

	aus, err := a.store.AdminUser().List(ctx, opts)

	if err != nil {
		log.L(ctx).Errorf("list admin users from storage failed: %s", err.Error())

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	var m sync.Map

	for _, item := range aus.Items {
		wg.Add(1)

		go func(au *model.AdminUser) {
			defer wg.Done()

			policies, err := a.store.Policies().List(ctx, au.Name, metav1.ListOptions{})
			if err != nil {
				errChan <- errors.WithCode(code.ErrDatabase, err.Error())

				return
			}

			m.Store(au.ID, &model.AdminUser{
				ObjectMeta: metav1.ObjectMeta{
					ID:              au.ID,
					InstanceID:      au.InstanceID,
					Name:            au.Name,
					CreatedAt:       au.CreatedAt,
					UpdatedAt:       au.UpdatedAt,
					Extend:          au.Extend,
					CreatedAtFormat: au.CreatedAtFormat,
					UpdatedAtFormat: au.UpdatedAtFormat,
				},
				NickName:    au.NickName,
				Email:       au.Email,
				Mobile:      au.Mobile,
				TotalPolicy: policies.TotalCount,
			})
		}(item)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}

	infos := make([]*model.AdminUser, 0, len(aus.Items))
	for _, item := range aus.Items {
		item, _ := m.Load(item.ID)
		infos = append(infos, item.(*model.AdminUser))
	}

	return &model.AdminUserList{ListMeta: aus.ListMeta, Items: infos}, nil
}

func (a *adminUserService) ListWithBadPerformance(ctx context.Context, opts metav1.ListOptions) (*model.AdminUserList, error) {
	aus, err := a.store.AdminUser().List(ctx, opts)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	infos := make([]*model.AdminUser, 0)

	for _, item := range aus.Items {
		policies, err := a.store.Policies().List(ctx, item.Name, metav1.ListOptions{})
		if err != nil {
			return nil, errors.WithCode(code.ErrDatabase, err.Error())
		}

		infos = append(infos, &model.AdminUser{
			ObjectMeta: metav1.ObjectMeta{
				ID:              item.ID,
				Name:            item.Name,
				CreatedAt:       item.CreatedAt,
				UpdatedAt:       item.UpdatedAt,
				CreatedAtFormat: item.CreatedAtFormat,
				UpdatedAtFormat: item.UpdatedAtFormat,
			},
			NickName:    item.NickName,
			Email:       item.Email,
			Mobile:      item.Mobile,
			TotalPolicy: policies.TotalCount,
		})
	}

	return &model.AdminUserList{ListMeta: aus.ListMeta, Items: aus.Items}, nil
}

func (a *adminUserService) Update(ctx context.Context, adminUser *model.AdminUser, opts metav1.UpdateOptions) error {
	if err := a.store.AdminUser().Update(ctx, adminUser, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (a *adminUserService) ChangePassword(ctx context.Context, adminUser *model.AdminUser) error {
	if err := a.store.AdminUser().Update(ctx, adminUser, metav1.UpdateOptions{}); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())

	}

	return nil
}
