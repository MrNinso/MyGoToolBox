package file

import (
	"bufio"
	"io"
	"os"
)

type File struct {
	*os.File
}

type OnLineRead func(number int, line []byte) (MoveOn bool)

type OnLineWrite func(number int) (data []byte, MoveOn bool)

type SetupReaderProxy func(this File) (io.Reader, error)

type SetupWriteProxy func(this io.Writer) (io.Writer, error)

func Open(path string) (File, error) {
	f, err := os.Open(path)

	return File{f}, err
}

func (f File) WriteLoop(writeHandle OnLineWrite, writerProxy SetupWriteProxy) error {
	var writer io.Writer

	if writerProxy != nil {
		w, err := writerProxy(f)
		if err != nil {
			return err
		}
		writer = w
	} else {
		writer = f
	}

	i := 0
	for {
		d, b := writeHandle(i)

		if _, err := writer.Write(d); err != nil {
			return err
		}

		if !b {
			break
		}
		i++
	}

	return f.Sync()
}

func (f File) ReadAllLines(handle OnLineRead, readproxy SetupReaderProxy) error {
	return f.ReadAllChuck(handle, []byte("\n")[0], readproxy)
}

func (f File) ReadAllChuck(handle OnLineRead, endChuck byte, readerProxy SetupReaderProxy) error {
	var reader *bufio.Reader

	if readerProxy != nil {
		r, err := readerProxy(f)
		if err != nil {
			return err
		}
		reader = bufio.NewReader(r)
	} else {
		reader = bufio.NewReader(f)
	}

	chuckNumber := 0
	for {
		b, err := reader.ReadBytes(endChuck)

		if len(b) > 0 {
			if !handle(chuckNumber, b[:len(b)-1]) {
				return nil
			}
		}

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		chuckNumber++
	}
}

func (f File) ResetSeekToStart() (int64, error) {
	return f.Seek(0, 0)
}

func (f File) ResetSeekToEnd() (int64, error) {
	return f.Seek(0, 2)
}
