package golevel7

import (
	"bufio"
	"io"

	"github.com/mehrdad-khojastefar/golevel7/commons"
)

type MessageScanner struct {
	r       io.Reader
	b       *bufio.Scanner
	thisMsg *Message
	err     error
}

// NewMessageScanner returns a new scanner that returns
// hl7 messages from an io.Reader
func NewMessageScanner(r io.Reader) *MessageScanner {
	ms := &MessageScanner{
		r: r,
		b: commons.NewBufScanner(r),
	}
	return ms
}

func (ms *MessageScanner) Scan() (gotOne bool) {
	var err error
	if scan := ms.b.Scan(); scan {
		if ms.err = ms.b.Err(); ms.err != nil || len(ms.b.Bytes()) < 5 {
			if ms.b.Bytes() != nil && !(len(ms.b.Bytes()) < 5) {
				gotOne = true
			}
		} else {
			gotOne = true
		}
		if gotOne {
			ms.thisMsg, err = NewMessage(ms.b.Bytes())
			if err != nil {
				return false
			}
		} else {
			ms.thisMsg = nil
		}
	}
	if !gotOne {
		ms.b = nil
	}
	return gotOne
}

func (ms *MessageScanner) Message() *Message {
	return ms.thisMsg
}

func (ms *MessageScanner) Err() error {
	return ms.err
}
