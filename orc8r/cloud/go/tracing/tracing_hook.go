/*
Copyright 2022 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/* Copyright The OpenTelemetry Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tracing

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"magma/orc8r/cloud/go/storage"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	trace "go.opentelemetry.io/otel/trace"
)

const (
	TracerName = "db.postgres"
	SpanName   = "orc8r.service"
	DBSystem   = "postgresql"
)

type TracingHook struct {
	tracerProvider *tracesdk.TracerProvider
	tracer         trace.Tracer
}

// initialize new tracing hook
func NewTracingHook() *TracingHook {
	th := new(TracingHook)
	th.tracerProvider = otel.GetTracerProvider().(*tracesdk.TracerProvider)
	th.tracer = otel.GetTracerProvider().Tracer(TracerName)
	return th
}

// parse a value in text for the given key
func parseKeyValue(key string, text string) string {
	var pairs = strings.Split(text, " ")

	for _, entry := range pairs {
		var pair = strings.Split(entry, "=")

		if len(pair) == 2 && pair[0] == key {
			return pair[1]
		}
	}
	return ""
}

func (th *TracingHook) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {

	if trace.SpanFromContext(ctx).IsRecording() {
		fmt.Printf("Before:: Record is ongoing\n")
		return ctx, nil
	}
	ctx, span := th.tracer.Start(ctx, SpanName)
	_ = span
	return ctx, nil
}

func (th *TracingHook) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	span := trace.SpanFromContext(ctx)

	if !span.IsRecording() {
		err := errors.New("Span Error")
		th.OnError(ctx, err, query)
		return ctx, err
	}
	defer span.End()
	operation := strings.Split(query, " ")[0]
	span.SetName(operation)
	attrs := make([]attribute.KeyValue, 0, 10)
	attrs = append(attrs,
		attribute.String("internal.span.format", "sql"),
		attribute.String("db.system", DBSystem),
		attribute.String("db.statement", query),
		attribute.String("db.host", parseKeyValue("host", storage.GetDatabaseSource())),
		attribute.String("db.user", parseKeyValue("user", storage.GetDatabaseSource())),
		attribute.String("db.name", parseKeyValue("dbname", storage.GetDatabaseSource())),
		attribute.String("db.ssl", parseKeyValue("sslmode", storage.GetDatabaseSource())),
	)
	span.SetAttributes(attrs...)
	span.SetStatus(codes.Ok, "OK")
	return ctx, nil
}

func (th *TracingHook) OnError(ctx context.Context, err error, query string) error {
	span := trace.SpanFromContext(ctx)

	if span != nil {
		span.SetStatus(codes.Error, err.Error())
		span.End()
	}
	return err
}
