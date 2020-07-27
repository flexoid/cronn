package service

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

const prefixCommandMaxlen = 16
const prefixCutCommandSuffix = "..."

// LogPrefixer implements `io.writer` interface and adds prefix to each output line.
type LogPrefixer struct {
	writer io.Writer
	prefix []byte
}

// NewLogPrefixer initializes log prefixer.
func NewLogPrefixer(writer io.Writer, command string) *LogPrefixer {
	logPrefixer := &LogPrefixer{writer: writer}
	logPrefixer.prefix = logPrefixer.prefixForCommand(command)
	return logPrefixer
}

func (p *LogPrefixer) Write(data []byte) (int, error) {
	reader := bufio.NewReader(bytes.NewReader(data))

	var line []byte
	var err error
	var bytesWritten int = 0

	for {
		line, err = reader.ReadBytes('\n')

		// There can be data to write in `line` even if `io.EOF` error is returned.
		// Exit immediately only in case of unexpected error.
		if err != nil && err != io.EOF {
			return bytesWritten, err
		}

		if len(line) > 0 {
			_, writeErr := p.writer.Write(p.prefix)
			if writeErr != nil {
				return bytesWritten, writeErr
			}

			n, writeErr := p.writer.Write(line)
			bytesWritten += n
			if writeErr != nil {
				return bytesWritten, writeErr
			}
		}

		if err == io.EOF {
			break
		}
	}

	return bytesWritten, nil
}

func (p *LogPrefixer) prefixForCommand(command string) []byte {
	if len(command) > prefixCommandMaxlen {
		command = command[:prefixCommandMaxlen]
		command += prefixCutCommandSuffix
	}

	prefix := fmt.Sprintf("{%s} ", command)
	return []byte(prefix)
}
