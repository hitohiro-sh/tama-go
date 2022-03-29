// MIT License
//
// Copyright (c) 2022 hitohiro-sh
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package tama

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
)

type ApiClient struct {
	child  *exec.Cmd
	writer io.Writer
}

func ApiClientNew(child *exec.Cmd) *ApiClient {
	return &ApiClient{
		child: child,
	}
}

func (client *ApiClient) Start() {
	writer, err := client.child.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	client.writer = writer

	client.child.Start()
}

func (client *ApiClient) Wait() {
	client.child.Wait()
}

func (client *ApiClient) Send(id, data string) {
	writer := client.writer
	obj, err := json.Marshal([]string{id, data})
	if err == nil {
		fmt.Fprint(writer, fmt.Sprintf("%s\n", obj))
	}
}

func (client *ApiClient) StdoutPipe() (io.ReadCloser, error) {
	return client.child.StdoutPipe()
}
