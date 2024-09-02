package service

import (
	"context"
	"encoding/json"
	"fmt"
	codewars "github.com/yasonofriychuk/service-codewars/internal/generate/proto/service"
	"io"
	"net/http"
	"net/url"
)

type Service struct {
	codewars.UnimplementedCodewarsServer
}

func New() *Service {
	return &Service{}
}

func (s Service) GetUserInfo(ctx context.Context, in *codewars.GetUserInfoIn) (*codewars.GetUserInfoOut, error) {
	requestUrl := fmt.Sprintf("https://www.codewars.com/api/v1/users/%s", url.PathEscape(in.GetUser()))

	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, fmt.Errorf("codewars request: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("codewars return status %s", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response body read: %w", err)
	}

	var out codewars.GetUserInfoOut
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, fmt.Errorf("response body unmarshal: %w", err)
	}

	return &out, nil
}
