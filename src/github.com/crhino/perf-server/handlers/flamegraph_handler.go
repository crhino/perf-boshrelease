package handlers

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func NewFlameGraphHandler(cmdPath string) *FlameGraphHandler {
	return &FlameGraphHandler{
		cmdPath: cmdPath,
	}
}

type FlameGraphHandler struct {
	cmdPath string
}

func (handler *FlameGraphHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	reader, err := handler.runFlameGraphShellCommands()
	log.Println("Finished running flamegraph commands")

	if err != nil {
		log.Println(err.Error())
		resp.WriteHeader(500)
		return
	}

	_, err = io.Copy(resp, reader)
	if err != nil {
		log.Println(err.Error())
	}

	reader.Close()
}

func (handler *FlameGraphHandler) runFlameGraphShellCommands() (io.ReadCloser, error) {
	path, err := ioutil.TempDir("", "perf")
	if err != nil {
		return nil, err
	}

	err = os.Chdir(path)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(handler.cmdPath, path, "graph.svg")
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return os.Open("graph.svg")
}
