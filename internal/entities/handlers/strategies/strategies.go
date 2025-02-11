// Copyright 2024 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package strategies contains the message creation strategies for entities and messages
package strategies

import (
	"context"

	watermill "github.com/ThreeDotsLabs/watermill/message"

	"github.com/mindersec/minder/internal/entities/handlers/message"
	"github.com/mindersec/minder/internal/entities/models"
)

// MessageCreateStrategy is the interface for creating messages
type MessageCreateStrategy interface {
	CreateMessage(
		ctx context.Context, ewp *models.EntityWithProperties,
	) (*watermill.Message, error)
	GetName() string
}

// GetEntityStrategy is the interface for getting entities
type GetEntityStrategy interface {
	GetEntity(
		ctx context.Context, entMsg *message.HandleEntityAndDoMessage,
	) (*models.EntityWithProperties, error)
	GetName() string
}
