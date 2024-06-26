// Copyright 2023 Stacklok, Inc
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
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"

	minderlogger "github.com/stacklok/minder/internal/logger"
	"github.com/stacklok/minder/internal/reconcilers/messages"
)

//nolint:exhaustive
func (r *Reconciler) handleEntityDeleteEvent(msg *message.Message) error {
	ctx := msg.Context()
	l := zerolog.Ctx(ctx).With().Logger()

	var event messages.RepoEvent
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

	l = zerolog.Ctx(ctx).With().
		Str("provider_id", event.ProviderID.String()).
		Str("project_id", event.ProjectID.String()).
		Str("repo_id", event.RepoID.String()).
		Logger()

	// Telemetry logging
	minderlogger.BusinessRecord(ctx).ProviderID = event.ProviderID
	minderlogger.BusinessRecord(ctx).Project = event.ProjectID

	l.Info().Msg("handling entity delete event")
	// Remove the entry in the DB. There's no need to clean any webhook we created for this repository, as GitHub
	// will automatically remove them when the repository is deleted.
	if err := r.repos.DeleteByID(ctx, event.RepoID, event.ProjectID); err != nil {
		return fmt.Errorf("error deleting repository from DB: %w", err)
	}

	minderlogger.BusinessRecord(ctx).Repository = event.RepoID
	return nil
}
