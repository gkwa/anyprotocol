package main

import (
	"context"

	"dagger/anyprotocol/internal/dagger"
)

type Anyprotocol struct{}

func (m *Anyprotocol) HelloModule() *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", "hello from anyprotocol"})
}

func (m *Anyprotocol) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

func (m *Anyprotocol) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}
