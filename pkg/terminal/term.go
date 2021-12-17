package terminal

import (
	"errors"
	mTerm "github.com/moby/term"
	"io"
)

// GetSize 获取控制台尺寸
func GetSize(w io.Writer) (width int, height int, err error) {

	fd, isTerm := mTerm.GetFdInfo(w)

	if !isTerm {
		return 0, 0, errors.New("given writer is not terminal")
	}

	size, err := mTerm.GetWinsize(fd)
	if err != nil {
		return 0, 0, err
	}

	return int(size.Width), int(size.Height), nil
}
