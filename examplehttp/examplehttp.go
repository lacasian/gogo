package examplehttp

import (
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
)

type Config struct {
	TargetURL string
}

type WorkerLib struct {
	config Config
}

func NewWorkerLib(config Config) *WorkerLib {
	return &WorkerLib{
		config,
	}
}

func (w *WorkerLib) Run() error {
	req := gorequest.New()
	resp, body, errs := req.Get(w.config.TargetURL).End()
	if errs != nil {
		return errors.Wrap(errs[0], "could not call target url")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("target returned response code != 200")
	}

	fmt.Println(body)

	return nil
}