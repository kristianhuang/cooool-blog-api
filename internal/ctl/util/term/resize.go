/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package term

import "github.com/moby/term"

type TerminalSize struct {
	Width  uint16
	Height uint16
}

type TerminalSizeQueue interface {
	Next() *TerminalSize
}

func (t TTY) GetSize() *TerminalSize {
	outFd, isTerminal := term.GetFdInfo(t.Out)
	if !isTerminal {
		return nil
	}

	return GetSize(outFd)
}

func GetSize(fd uintptr) *TerminalSize {
	winsize, err := term.GetWinsize(fd)
	if err != nil {
		return nil
	}

	return &TerminalSize{
		Width:  winsize.Width,
		Height: winsize.Height,
	}
}
