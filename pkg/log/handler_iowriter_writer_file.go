package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/justtrackio/gosoline/pkg/cfg"
)

func init() {
	AddHandlerIoWriterFactory("file", ioWriterFileFactory)
}

type ioWriterFileSettings struct {
	Path string `cfg:"path" default:"logs.log"`
}

func ioWriterFileFactory(config cfg.Config, configKey string) (io.Writer, error) {
	settings := &ioWriterFileSettings{}
	config.UnmarshalKey(configKey, settings)

	return NewIoWriterFile(settings.Path)
}

func NewIoWriterFile(path string) (io.Writer, error) {
	var err error
	var file *os.File

	ex, _ := os.Executable()
	logfilePath := filepath.Join(filepath.Dir(ex), path)

	if file, err = os.OpenFile(logfilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o600); err != nil {
		return nil, fmt.Errorf("can not open file %s to write logs to: %w", path, err)
	}

	return file, nil
}
