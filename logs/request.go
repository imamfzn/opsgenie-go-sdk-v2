package logs

import (
	"github.com/pkg/errors"
	"net/url"
	"strconv"
)

type ListLogFilesRequest struct {
	Marker string
	Limit  int
	params string
}

func (r ListLogFilesRequest) Validate() error {
	if len(r.Marker) == 0 {
		return errors.New("marker cannot be empty")
	}

	return nil
}

func (r ListLogFilesRequest) Endpoint() string {
	return "/v2/logs/list/" + r.Marker + r.setParams(r)
}

func (r ListLogFilesRequest) Method() string {
	return "GET"
}

func (r ListLogFilesRequest) setParams(request ListLogFilesRequest) string {
	params := url.Values{}

	if request.Limit >= 0 {
		params.Add("limit", strconv.Itoa(request.Limit))
	}

	if len(params) > 0 {
		request.params = "?" + params.Encode()
	} else {
		request.params = ""
	}

	return request.params
}

type GenerateLogFileDownloadLinkRequest struct {
	FileName string
}

func (r GenerateLogFileDownloadLinkRequest) Validate() error {
	if len(r.FileName) == 0 {
		return errors.New("fileName cannot be empty")
	}

	return nil
}

func (r GenerateLogFileDownloadLinkRequest) Endpoint() string {
	return "/v2/logs/download/" + r.FileName
}

func (r GenerateLogFileDownloadLinkRequest) Method() string {
	return "GET"
}