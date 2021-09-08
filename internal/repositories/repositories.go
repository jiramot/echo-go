package repository

import (
    "encoding/json"
    "github.com/matiasvarela/errors"
    "io/ioutil"
    "jiramot/echo-go/internal/core/domain"
    "jiramot/echo-go/internal/pkg"
    "jiramot/echo-go/pkg/apperrors"
)

type repository struct {

}

func NewEchoRepository() *repository {
    return &repository{}
}

func (repo *repository) EchoMessage(message string) (domain.Echo, error){
    response, err := restclient.Get("https://postman-echo.com/get?message=" + message)
    if err != nil {
        return domain.Echo{}, errors.New(apperrors.Internal, err, "message", "cause message")
    }
    bytes, err := ioutil.ReadAll(response.Body)
    defer response.Body.Close()
    var result EchoResponse
    if err := json.Unmarshal(bytes, &result); err != nil {
        return domain.Echo{}, errors.New(apperrors.Internal, err, "message", "cause message")
    }
    return domain.Echo{
        Message: result.Args.Message,
    }, nil
}

type EchoResponse struct {
    Args    Args    `json:"args"`
}
type Args struct {
    Message string  `json:"message"`
}