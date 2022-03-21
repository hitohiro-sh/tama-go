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
)

type Api struct {
	apiMap map[string]func(data string)
	reader io.Reader
}

func ApiNew(reader io.Reader) *Api {
	return &Api{
		apiMap: map[string]func(arg string){},
		reader: reader,
	}
}

func (api *Api) On(id string, callback func(arg string)) {
	api.apiMap[id] = callback
}

func (api *Api) Run() {
	apicall := func(m, v string) {
		if f, ok := api.apiMap[m]; ok {
			f(v)
		}
	}

	for {
		var buf string
		fmt.Fscanln(api.reader, &buf)

		var inst []string
		err := json.Unmarshal([]byte(buf), &inst)
		if err != nil {
			log.Fatal(err)
		}
		apicall(inst[0], inst[1])
	}
}
