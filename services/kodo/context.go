// Code generated by go generate via internal/cmd/context; DO NOT EDIT.
package kodo

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"

	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/types"
	"github.com/Xuanwo/storage/types/metadata"
	ps "github.com/Xuanwo/storage/types/pairs"
)

var _ storage.Storager

// CreateWithContext adds context support for Create.
func (s *Service) CreateWithContext(ctx context.Context, name string, pairs ...*types.Pair) (storage.Storager, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/kodo.Servicer.Create")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Create(name, pairs...)
}

// DeleteWithContext adds context support for Delete.
func (s *Service) DeleteWithContext(ctx context.Context, name string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/kodo.Servicer.Delete")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Delete(name, pairs...)
}

// GetWithContext adds context support for Get.
func (s *Service) GetWithContext(ctx context.Context, name string, pairs ...*types.Pair) (storage.Storager, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/kodo.Servicer.Get")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Get(name, pairs...)
}

// ListWithContext adds context support for List.
func (s *Service) ListWithContext(ctx context.Context, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/kodo.Servicer.List")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.List(pairs...)
}

// DeleteWithContext adds context support for Delete.
func (s *Storage) DeleteWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/kodo.Storager.Delete")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Delete(path, pairs...)
}

// InitWithContext adds context support for Init.
func (s *Storage) InitWithContext(ctx context.Context, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/kodo.Storager.Init")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Init(pairs...)
}

// ListWithContext adds context support for List.
func (s *Storage) ListWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/kodo.Storager.List")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.List(path, pairs...)
}

// MetadataWithContext adds context support for Metadata.
func (s *Storage) MetadataWithContext(ctx context.Context, pairs ...*types.Pair) (m metadata.StorageMeta, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/kodo.Storager.Metadata")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Metadata(pairs...)
}

// ReadWithContext adds context support for Read.
func (s *Storage) ReadWithContext(ctx context.Context, path string, pairs ...*types.Pair) (r io.ReadCloser, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/kodo.Storager.Read")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Read(path, pairs...)
}

// StatWithContext adds context support for Stat.
func (s *Storage) StatWithContext(ctx context.Context, path string, pairs ...*types.Pair) (o *types.Object, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/kodo.Storager.Stat")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Stat(path, pairs...)
}

// WriteWithContext adds context support for Write.
func (s *Storage) WriteWithContext(ctx context.Context, path string, r io.Reader, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/kodo.Storager.Write")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Write(path, r, pairs...)
}
