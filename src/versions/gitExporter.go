package versions

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

type GitVersionExporter struct {
	versionFileUrl string
}

func NewGitVersionExporter(url string) IVersionExporter {
	return GitVersionExporter{versionFileUrl: url}
}

func (exporter GitVersionExporter) Load() (Version, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, exporter.versionFileUrl, nil)
	if err != nil {
		return Version{}, err
	}

	req = req.WithContext(ctx)
	res, err := client.Do(req)

	if err != nil {
		return Version{}, err
	}
	defer res.Body.Close()

	var bodyBuffer bytes.Buffer

	_, err = bodyBuffer.ReadFrom(res.Body)
	if err != nil {
		return Version{}, err
	}

	return NewVersion(bodyBuffer.String())
}
