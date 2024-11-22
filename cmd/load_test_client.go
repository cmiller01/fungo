// Copyright © 2017 Christian Miller
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	rng "github.com/leesper/go_rng"
	"github.com/spf13/cobra"
)

// load_test_clientCmd represents the load_test_client command
var load_test_clientCmd = &cobra.Command{
	Use:   "load_test_client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		r = rng.NewPoissonGenerator(1)
		Loop("Constant", constant)
		Loop("Uniform", uniform)
		Loop("Poisson", poisson)
	},
}

const delay = 1 * time.Millisecond
const N = 1000

var r *rng.PoissonGenerator

type delayFunc func() float32

func constant() float32 {
	return 1
}
func uniform() float32 {
	return rand.Float32()
}
func poisson() float32 {
	return float32(r.Poisson(1))
}

func Loop(name string, fn delayFunc) {
	fmt.Println(name)
	for i := 0; i < N; i++ {
		v := fn() * float32(delay)
		time.Sleep(time.Duration(v))
		makeRequest()
	}
	fmt.Println()
}

func makeRequest() {
	url := "http://localhost:5000/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
}

func init() {
	RootCmd.AddCommand(load_test_clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// load_test_clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// load_test_clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
