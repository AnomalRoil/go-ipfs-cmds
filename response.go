package cmds

import (
	"fmt"

	"gx/ipfs/Qmf7G7FikwUsm48Jm4Yw4VBGNZuyRaAMzpWDJcW8V71uV2/go-ipfs-cmdkit"
)

var (
	ErrRcvdError = fmt.Errorf("received command error")
)

// Response is the result of a command request. Response is returned to the client.
type Response interface {
	Request() Request

	Error() *cmdsutil.Error
	Length() uint64

	// Next returns the next emitted value.
	// The returned error can be a network or decoding error.
	// The error can also be ErrRcvdError if an error has been emitted.
	// In this case the emitted error can be accessed using the Error() method.
	Next() (interface{}, error)
}

type Head struct {
	Len uint64
	Err *cmdsutil.Error
}

func (h Head) Length() uint64 {
	return h.Len
}

func (h Head) Error() *cmdsutil.Error {
	return h.Err
}
