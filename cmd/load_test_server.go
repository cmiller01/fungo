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
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// loadTestServerCmd represents the loadTestServer command
var loadTestServerCmd = &cobra.Command{
	Use:   "load_test_server",
	Short: "Running a server that returns ok (but sometimes is locked)",
	Run: func(cmd *cobra.Command, args []string) {
		var locked bool
		fn := func(w http.ResponseWriter, r *http.Request) {
			if locked {
				fmt.Println("yikes, locked")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// not sure exactly what i was looking to do with this one...
			locked = true
			go time.AfterFunc(time.Millisecond, func() { locked = false })
		}
		http.HandleFunc("/", fn)
		log.Printf("Listening on 5000")
		http.ListenAndServe(":5000", nil)
		return
	},
}

func init() {
	RootCmd.AddCommand(loadTestServerCmd)
}
