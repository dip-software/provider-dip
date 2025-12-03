//go:build tools
// +build tools

package tools

// This file tracks dependencies for tools used in development and build processes.
// The blank imports ensure that tool dependencies are recorded in go.mod and not
// removed by 'go mod tidy'.

import (
	_ "golang.org/x/tools/cmd/goimports"
)
