package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

func (s Server) HandleDeploy(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(err, "failed to read body"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	fmt.Fprintln(os.Stdout, "Received deploy request: ")
	fmt.Fprintln(os.Stdout, string(body))

	w.WriteHeader(http.StatusOK)
}
