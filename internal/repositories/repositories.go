package repository

import (
    "encoding/json"
    "fmt"
    "github.com/matiasvarela/errors"
    "io/ioutil"
    "jiramot/echo-go/internal/core/domain"
    "jiramot/echo-go/internal/pkg"
    "jiramot/echo-go/pkg/apperrors"
)

type echoRepository struct {
}

func NewEchoRepository() *echoRepository {
    return &echoRepository{}
}

func (repo *echoRepository) EchoMessage(message string) (domain.Echo, error) {
    url := fmt.Sprintf("https://postman-echo.com/get?message=%s", message)
    response, err := restclient.Get(url)
    if err != nil {
        return domain.Echo{}, errors.New(apperrors.Internal, err, "Error", "Endpoint error")
    }
    bytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return domain.Echo{}, errors.New(apperrors.Internal, err, "Error", "Read stream error")
    }
    defer response.Body.Close()
    var result PostmanEchoResponse
    if err := json.Unmarshal(bytes, &result); err != nil {
        return domain.Echo{}, errors.New(apperrors.Internal, err, "message", "Parse json error")
    }
    return domain.Echo{
        Message: result.Args.Message,
    }, nil
}

type PostmanEchoResponse struct {
    Args Args `json:"args"`
}
type Args struct {
    Message string `json:"message"`
}
