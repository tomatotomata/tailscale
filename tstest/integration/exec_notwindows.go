// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows

package integration

import (
	"errors"
	"os"
	"os/exec"
)

// Non-Windows stubs for the peer process helpers; never called, since their
// callers are guarded by runtime.GOOS == "windows".

func setNewProcessGroup(cmd *exec.Cmd) {}

func gracefulStop(process *os.Process) error {
	return errors.New("CTRL_BREAK is only supported on Windows")
}
