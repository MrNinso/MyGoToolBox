package file

import (
	"compress/gzip"
	"io"
)

func NewGzipReaderProxy() SetupReaderProxy {
	return func(this File) (io.Reader, error) {
		return gzip.NewReader(this)
	}
}

func NewGzipWriteProxy() SetupWriteProxy {
	return func(this io.Writer) (io.Writer, error) {
		return gzip.NewWriter(this), nil
	}
}

func NewGzipWriteProxyLevel(level int) SetupWriteProxy {
	return func(this io.Writer) (io.Writer, error) {
		return gzip.NewWriterLevel(this, level)
	}
}
