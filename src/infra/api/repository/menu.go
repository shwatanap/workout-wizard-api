package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/shwatanap/workout-wizard-api/src/domain/entity"
	"github.com/shwatanap/workout-wizard-api/src/domain/repository"
	"github.com/shwatanap/workout-wizard-api/src/infra/api/request"
	"github.com/shwatanap/workout-wizard-api/src/infra/api/response"
	"github.com/shwatanap/workout-wizard-api/src/utils"
)

type menuExternalRepository struct {
}

func NewMenuExternalRepository() repository.IMenuExternalRepository {
	return &menuExternalRepository{}
}

func (mer *menuExternalRepository) Create(ctx context.Context, userID int, target, comment string) (*entity.Menu, error) {
	messages, err := createMessages()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(request.NewChatGPTRequest(os.Getenv("MODEL"), messages))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, os.Getenv("COMPLETION_ENDPOINT"), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("CHATGPT_API_KEY"))
	body, err := sendHttpRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	// client := &http.Client{
	// 	// リソース節約のためにタイムアウトを設定する
	// 	Timeout: 20 * time.Second,
	// }

	// resp, err := client.Do(req)
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	return nil, errors.New("bad status: " + resp.Status)
	// }

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }

	var res response.Response
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	me := utils.ParseMenu(res.Choices[0].Message.Content)

	return me, nil
}

func createMessages() ([]*request.RequestMessage, error) {
	return []*request.RequestMessage{
		request.NewRequestMessage("user", "Hello"),
	}, nil
}

func sendHttpRequest(ctx context.Context, req *http.Request) ([]byte, error) {
	client := &http.Client{
		// リソース節約のためにタイムアウトを設定する
		Timeout: 20 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad status: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
