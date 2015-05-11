package portable

import (
	"github.com/kpmy/ypk/assert"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Data interface {
	Get(name string) (data []byte, err error)
}

type httpData struct{}

func (d *httpData) Get(n string) (data []byte, err error) {
	resp, err := http.Get("../netlist/" + n)
	assert.For(err == nil && resp.StatusCode == http.StatusOK, 40)
	return ioutil.ReadAll(resp.Body)
}

var DataSource Data

type fileData struct{}

func (f *fileData) Get(fn string) (data []byte, err error) {
	path, _ := os.Getwd()
	input, err := os.Open(filepath.Join(path, "netlist", fn))
	assert.For(err == nil, 40, path, fn)
	return ioutil.ReadAll(input)
}

func init() {
	DataSource = &httpData{}
	//DataSource = &fileData{}
}
