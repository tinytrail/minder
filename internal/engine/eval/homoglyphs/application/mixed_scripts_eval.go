// Copyright 2023 Stacklok, Inc.
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

package application

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	evalerrors "github.com/mindersec/minder/internal/engine/errors"
	"github.com/mindersec/minder/internal/engine/eval/homoglyphs/communication"
	"github.com/mindersec/minder/internal/engine/eval/homoglyphs/domain"
	"github.com/mindersec/minder/internal/engine/eval/templates"
	engif "github.com/mindersec/minder/internal/engine/interfaces"
	eoptions "github.com/mindersec/minder/internal/engine/options"
	provifv1 "github.com/mindersec/minder/pkg/providers/v1"
)

// MixedScriptsEvaluator is the evaluator for the mixed scripts rule type
type MixedScriptsEvaluator struct {
	processor     domain.HomoglyphProcessor
	reviewHandler *communication.GhReviewPrHandler
}

// NewMixedScriptEvaluator creates a new mixed scripts evaluator
func NewMixedScriptEvaluator(
	ctx context.Context,
	ghClient provifv1.GitHub,
	opts ...eoptions.Option,
) (*MixedScriptsEvaluator, error) {
	msProcessor, err := domain.NewMixedScriptsProcessor(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not create mixed scripts processor: %w", err)
	}

	evaluator := &MixedScriptsEvaluator{
		processor:     msProcessor,
		reviewHandler: communication.NewGhReviewPrHandler(ghClient),
	}

	for _, opt := range opts {
		if err := opt(evaluator); err != nil {
			return nil, err
		}
	}

	return evaluator, nil
}

// Eval evaluates the mixed scripts rule type
func (mse *MixedScriptsEvaluator) Eval(
	ctx context.Context,
	_ map[string]any,
	_ protoreflect.ProtoMessage,
	res *engif.Result,
) error {
	violations, err := evaluateHomoglyphs(ctx, mse.processor, res, mse.reviewHandler)
	if err != nil {
		return err
	}

	if len(violations) > 0 {
		return evalerrors.NewDetailedErrEvaluationFailed(
			templates.MixedScriptsTemplate,
			map[string]any{"violations": violations},
			"found mixed scripts violations",
		)
	}

	return nil
}
