// Code generated by go generate internal/cmd/pairs; DO NOT EDIT.
package pairs

import (
	"context"

	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/pkg/credential"
	"github.com/Xuanwo/storage/pkg/endpoint"
	"github.com/Xuanwo/storage/pkg/segment"
	"github.com/Xuanwo/storage/types"
)

// All available pairs.
const (
	Checksum     = "checksum"
	Context      = "context"
	Credential   = "credential"
	DirFunc      = "dir_func"
	Endpoint     = "endpoint"
	Expire       = "expire"
	FileFunc     = "file_func"
	Location     = "location"
	Name         = "name"
	Offset       = "offset"
	PartSize     = "part_size"
	Project      = "project"
	SegmentFunc  = "segment_func"
	Size         = "size"
	StorageClass = "storage_class"
	StoragerFunc = "storager_func"
	Type         = "type"
	WorkDir      = "work_dir"
)

// WithChecksum will apply checksum value to Options
func WithChecksum(v string) *types.Pair {
	return &types.Pair{
		Key:   Checksum,
		Value: v,
	}
}

// WithContext will apply context value to Options
func WithContext(v context.Context) *types.Pair {
	return &types.Pair{
		Key:   Context,
		Value: v,
	}
}

// WithCredential will apply credential value to Options
func WithCredential(v *credential.Provider) *types.Pair {
	return &types.Pair{
		Key:   Credential,
		Value: v,
	}
}

// WithDirFunc will apply dir_func value to Options
func WithDirFunc(v types.ObjectFunc) *types.Pair {
	return &types.Pair{
		Key:   DirFunc,
		Value: v,
	}
}

// WithEndpoint will apply endpoint value to Options
func WithEndpoint(v endpoint.Provider) *types.Pair {
	return &types.Pair{
		Key:   Endpoint,
		Value: v,
	}
}

// WithExpire will apply expire value to Options
func WithExpire(v int) *types.Pair {
	return &types.Pair{
		Key:   Expire,
		Value: v,
	}
}

// WithFileFunc will apply file_func value to Options
func WithFileFunc(v types.ObjectFunc) *types.Pair {
	return &types.Pair{
		Key:   FileFunc,
		Value: v,
	}
}

// WithLocation will apply location value to Options
func WithLocation(v string) *types.Pair {
	return &types.Pair{
		Key:   Location,
		Value: v,
	}
}

// WithName will apply name value to Options
func WithName(v string) *types.Pair {
	return &types.Pair{
		Key:   Name,
		Value: v,
	}
}

// WithOffset will apply offset value to Options
func WithOffset(v int64) *types.Pair {
	return &types.Pair{
		Key:   Offset,
		Value: v,
	}
}

// WithPartSize will apply part_size value to Options
func WithPartSize(v int64) *types.Pair {
	return &types.Pair{
		Key:   PartSize,
		Value: v,
	}
}

// WithProject will apply project value to Options
func WithProject(v string) *types.Pair {
	return &types.Pair{
		Key:   Project,
		Value: v,
	}
}

// WithSegmentFunc will apply segment_func value to Options
func WithSegmentFunc(v segment.Func) *types.Pair {
	return &types.Pair{
		Key:   SegmentFunc,
		Value: v,
	}
}

// WithSize will apply size value to Options
func WithSize(v int64) *types.Pair {
	return &types.Pair{
		Key:   Size,
		Value: v,
	}
}

// WithStorageClass will apply storage_class value to Options
func WithStorageClass(v string) *types.Pair {
	return &types.Pair{
		Key:   StorageClass,
		Value: v,
	}
}

// WithStoragerFunc will apply storager_func value to Options
func WithStoragerFunc(v storage.StoragerFunc) *types.Pair {
	return &types.Pair{
		Key:   StoragerFunc,
		Value: v,
	}
}

// WithType will apply type value to Options
func WithType(v string) *types.Pair {
	return &types.Pair{
		Key:   Type,
		Value: v,
	}
}

// WithWorkDir will apply work_dir value to Options
func WithWorkDir(v string) *types.Pair {
	return &types.Pair{
		Key:   WorkDir,
		Value: v,
	}
}
