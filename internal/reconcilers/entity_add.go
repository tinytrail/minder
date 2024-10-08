// Copyright 2024 Stacklok, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reconcilers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/stacklok/minder/internal/entities/properties"
	"github.com/stacklok/minder/internal/logger"
	"github.com/stacklok/minder/internal/reconcilers/messages"
)

// handleEntityAddEvent handles the entity add event.
// Although this method is meant to be generic and handle all types of entities,
// it currently only does so for repositories.
// Todo: Utilise for other entities when such are supported.
// nolint
func (r *Reconciler) handleEntityAddEvent(msg *message.Message) error {
	ctx := msg.Context()
	l := zerolog.Ctx(ctx).With().Logger()

	var event messages.MinderEvent
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return fmt.Errorf("error unmarshalling payload: %w", err)
	}

	// validate event
	validate := validator.New()
	if err := validate.Struct(&event); err != nil {
		// We don't return the event since there's no use
		// retrying it if it's invalid.
		l.Error().Err(err).Msg("error validating event")
		return nil
	}

	var repoOwner string
	var repoName string
	var ok bool
	if repoOwner, ok = event.Entity["repoOwner"].(string); !ok {
		return errors.New("invalid repo owner")
	}
	if repoName, ok = event.Entity["repoName"].(string); !ok {
		return errors.New("invalid repo name")
	}

	// TODO: This should be using the properties map coming from the event
	// as-is, but for now we are using the repoOwner and repoName to create
	// the properties map.
	fetchByProps, err := properties.NewProperties(map[string]interface{}{
		properties.PropertyName: fmt.Sprintf("%s/%s", repoOwner, repoName),
	})
	if err != nil {
		return fmt.Errorf("error creating properties: %w", err)
	}

	l = zerolog.Ctx(ctx).With().
		Str("provider_id", event.ProviderID.String()).
		Str("project_id", event.ProjectID.String()).
		Str("repo_name", repoName).
		Str("repo_owner", repoOwner).
		Logger()

	// Telemetry logging
	logger.BusinessRecord(ctx).ProviderID = event.ProviderID
	logger.BusinessRecord(ctx).Project = event.ProjectID

	dbProvider, err := r.store.GetProviderByID(ctx, event.ProviderID)
	if err != nil {
		return fmt.Errorf("error retrieving provider: %w", err)
	}

	pbRepo, err := r.repos.CreateRepository(ctx, &dbProvider, event.ProjectID, fetchByProps)
	if err != nil {
		return fmt.Errorf("error add repository from DB: %w", err)
	}

	if pbRepo.Id == nil {
		return fmt.Errorf("repository id is nil")
	}
	repoID, err := uuid.Parse(*pbRepo.Id)
	if err != nil {
		return fmt.Errorf("repository id is not a UUID: %w", err)
	}

	logger.BusinessRecord(ctx).Repository = repoID
	return nil
}
