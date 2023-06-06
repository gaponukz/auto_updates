package versions

import "os"

type FileVersionExporter struct {
	filename string
}

func NewFileVersionExporter(filename string) IVersionExporter {
	return FileVersionExporter{filename: filename}
}

func (exporter FileVersionExporter) Load() (Version, error) {
	content, err := os.ReadFile(exporter.filename)

	if err != nil {
		return Version{}, err
	}

	return NewVersion(string(content))
}
