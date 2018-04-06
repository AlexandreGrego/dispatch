///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package identitymanager

import (
	"fmt"
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"

	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/vmware/dispatch/pkg/entity-store"
	"github.com/vmware/dispatch/pkg/identity-manager/gen/models"
	serviceAccountOperations "github.com/vmware/dispatch/pkg/identity-manager/gen/restapi/operations/serviceaccount"
	"github.com/vmware/dispatch/pkg/trace"
)

func serviceAccountModelToEntity(m *models.ServiceAccount) *ServiceAccount {
	defer trace.Tracef("name '%s'", *m.Name)()

	e := ServiceAccount{
		BaseEntity: entitystore.BaseEntity{
			OrganizationID: IdentityManagerFlags.OrgID,
			Name:           *m.Name,
		},
	}
	e.PublicKey = *m.PublicKey
	// We don't allow users to change the algorithm for now.
	e.JWTAlgorithm = "RS256"
	// TODO: set the domain from user
	e.Domain = IdentityManagerFlags.ServiceAccountDomain
	return &e
}

func serviceAccountEntityToModel(e *ServiceAccount) *models.ServiceAccount {
	defer trace.Tracef("name '%s'", e.Name)()
	m := models.ServiceAccount{
		ID:           strfmt.UUID(e.ID),
		Name:         swag.String(e.Name),
		Status:       models.Status(e.Status),
		CreatedTime:  e.CreatedTime.Unix(),
		ModifiedTime: e.ModifiedTime.Unix(),
	}
	m.PublicKey = &e.PublicKey
	return &m
}

func (h *Handlers) getServiceAccounts(params serviceAccountOperations.GetServiceAccountsParams, principal interface{}) middleware.Responder {

	defer trace.Trace("")()
	var serviceAccounts []*ServiceAccount

	opts := entitystore.Options{
		Filter: entitystore.FilterExists(),
	}
	err := h.store.List(IdentityManagerFlags.OrgID, opts, &serviceAccounts)
	if err != nil {
		log.Errorf("store error when listing service accounts: %+v", err)
		return serviceAccountOperations.NewGetServiceAccountsInternalServerError().WithPayload(
			&models.Error{
				Code:    http.StatusInternalServerError,
				Message: swag.String("internal server error when getting service accounts"),
			})
	}
	var serviceAccountModels []*models.ServiceAccount
	for _, serviceAccount := range serviceAccounts {
		serviceAccountModels = append(serviceAccountModels, serviceAccountEntityToModel(serviceAccount))
	}
	return serviceAccountOperations.NewGetServiceAccountsOK().WithPayload(serviceAccountModels)
}

func (h *Handlers) getServiceAccount(params serviceAccountOperations.GetServiceAccountParams, principal interface{}) middleware.Responder {

	defer trace.Tracef("get service account name '%s'", params.ServiceAccountName)()
	var serviceAccount ServiceAccount

	opts := entitystore.Options{
		Filter: entitystore.FilterExists(),
	}

	name := params.ServiceAccountName
	if err := h.store.Get(IdentityManagerFlags.OrgID, name, opts, &serviceAccount); err != nil {
		log.Errorf("store error when getting service account '%s': %+v", name, err)
		return serviceAccountOperations.NewGetServiceAccountNotFound().WithPayload(
			&models.Error{
				Code:    http.StatusNotFound,
				Message: swag.String("service account not found"),
			})
	}

	serviceAccountModel := serviceAccountEntityToModel(&serviceAccount)

	return serviceAccountOperations.NewGetServiceAccountOK().WithPayload(serviceAccountModel)
}

func (h *Handlers) addServiceAccount(params serviceAccountOperations.AddServiceAccountParams, principal interface{}) middleware.Responder {
	defer trace.Trace("")()
	serviceAccountRequest := params.Body
	e := serviceAccountModelToEntity(serviceAccountRequest)

	e.Status = entitystore.StatusREADY

	if err := validateServiceAccountEntity(e); err != nil {
		return serviceAccountOperations.NewAddServiceAccountBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: swag.String(fmt.Sprintf("error validating service account: %s", err)),
		})
	}

	if _, err := h.store.Add(e); err != nil {
		if entitystore.IsUniqueViolation(err) {
			return serviceAccountOperations.NewAddServiceAccountConflict().WithPayload(&models.Error{
				Code:    http.StatusConflict,
				Message: swag.String("error creating service account: non-unique name"),
			})
		}
		log.Errorf("store error when adding a new service account %s: %+v", e.Name, err)
		return serviceAccountOperations.NewAddServiceAccountInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: swag.String("internal server error when storing new service account"),
		})
	}

	return serviceAccountOperations.NewAddServiceAccountCreated().WithPayload(serviceAccountEntityToModel(e))
}

func (h *Handlers) deleteServiceAccount(params serviceAccountOperations.DeleteServiceAccountParams, principal interface{}) middleware.Responder {
	defer trace.Tracef("name '%s'", params.ServiceAccountName)()
	name := params.ServiceAccountName

	opts := entitystore.Options{
		Filter: entitystore.FilterExists(),
	}

	var e ServiceAccount
	if err := h.store.Get(IdentityManagerFlags.OrgID, name, opts, &e); err != nil {
		log.Errorf("store error when getting service account: %+v", err)
		return serviceAccountOperations.NewDeleteServiceAccountNotFound().WithPayload(
			&models.Error{
				Code:    http.StatusNotFound,
				Message: swag.String("service account not found"),
			})
	}

	if e.Status == entitystore.StatusDELETING {
		log.Warnf("Attempting to delete service account  %s which already is in DELETING state: %+v", e.Name)
		return serviceAccountOperations.NewDeleteServiceAccountBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: swag.String(fmt.Sprintf("Unable to delete service account %s: service account is already being deleted", e.Name)),
		})
	}

	e.Status = entitystore.StatusDELETING
	if err := h.store.Delete(e.OrganizationID, e.Name, &e); err != nil {
		log.Errorf("store error when deleting a service account %s: %+v", e.Name, err)
		return serviceAccountOperations.NewDeleteServiceAccountInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: swag.String("internal server error when deleting a service account"),
		})
	}

	return serviceAccountOperations.NewDeleteServiceAccountOK().WithPayload(serviceAccountEntityToModel(&e))
}

func (h *Handlers) updateServiceAccount(params serviceAccountOperations.UpdateServiceAccountParams, principal interface{}) middleware.Responder {
	defer trace.Tracef("updated serviceAccount '%s'", params.ServiceAccountName)()

	opts := entitystore.Options{
		Filter: entitystore.FilterExists(),
	}

	e := ServiceAccount{}
	if err := h.store.Get(IdentityManagerFlags.OrgID, params.ServiceAccountName, opts, &e); err != nil {
		log.Errorf("store error when getting service account: %+v", err)
		return serviceAccountOperations.NewUpdateServiceAccountNotFound().WithPayload(
			&models.Error{
				Code:    http.StatusNotFound,
				Message: swag.String("service account not found"),
			})
	}

	updateEntity := serviceAccountModelToEntity(params.Body)
	updateEntity.CreatedTime = e.CreatedTime
	updateEntity.ID = e.ID
	updateEntity.Status = entitystore.StatusREADY

	if err := validateServiceAccountEntity(updateEntity); err != nil {
		return serviceAccountOperations.NewUpdateServiceAccountBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: swag.String(fmt.Sprintf("error validating service account: %s", err)),
		})
	}

	if _, err := h.store.Update(e.Revision, updateEntity); err != nil {
		log.Errorf("store error when updating a service account %s: %+v", e.Name, err)
		return serviceAccountOperations.NewUpdateServiceAccountInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: swag.String("internal server error when updating a service account"),
		})
	}

	return serviceAccountOperations.NewUpdateServiceAccountOK().WithPayload(serviceAccountEntityToModel(updateEntity))
}

func validateServiceAccountEntity(e *ServiceAccount) error {
	// Validate public key provided by user
	pubKeyPEM, err := base64.StdEncoding.DecodeString(e.PublicKey)
	if err != nil {
		log.Debugf("Error validating service account %s: error %s", e.Name, err)
		return errors.New("public key is not base64 encoded")
	}
	_, err = jwt.ParseRSAPublicKeyFromPEM(pubKeyPEM)
	if err != nil {
		log.Debugf("Error validating service account %s: error %s", e.Name, err)
		return errors.New("invalid public key or public key not in PEM format")
	}
	return nil
}