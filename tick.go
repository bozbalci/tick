// Copyright (c) 2017, Berk Ozbalci
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"os"

	"github.com/urfave/cli"

	"tick/track"
)

func main() {
	tickApp := cli.NewApp()
	tickApp.Name = "tick"
	tickApp.Usage = "One bit journal for the command-line"
	tickApp.Version = "0.0.1"

	tickApp.Action = func(c *cli.Context) error {
		name := c.Args().First()
		t := track.New(name)

		err := t.TickToday()

		return err
	}

	tickApp.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"cr"},
			Usage:   "create a new track",
			Action: func(c *cli.Context) error {
				name := c.Args().First()
				t := track.New(name)

				err := t.Create()
				if err == nil {
					fmt.Printf("Created track %q\n", name)
				}
				return err
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"del"},
			Usage:   "delete an existing track",
			Action: func(c *cli.Context) error {
				name := c.Args().First()
				t := track.New(name)

				err := t.Delete()
				if err == nil {
					fmt.Printf("Deleted track %q\n", name)
				}

				return err
			},
		},
	}

	tickApp.Run(os.Args)
}
