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
	"fmt"
	"os"

	"github.com/urfave/cli"

	"tick/track"
)

func main() {
	tickApp := cli.NewApp()
	tickApp.Name = "tick"
	tickApp.Usage = "One bit journal for the command-line"
	tickApp.Version = "0.0.1"
	tickApp.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "create a new track",
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
			Name:  "delete",
			Usage: "delete an existing track",
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
		{
			Name:    "tick",
			Aliases: []string{"t"},
			Usage:   "ticks a track",
			Action: func(c *cli.Context) error {
				name := c.Args().First()
				t := track.New(name)

				err := t.TickToday()
				return err
			},
		},
		{
			Name:    "correlate",
			Aliases: []string{"corr"},
			Usage:   "correlate two tracks",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	tickApp.CommandNotFound = func(c *cli.Context, cmd string) {
		fmt.Fprintf(os.Stderr, "tick: '%s' is not a tick command. See 'tick help'.\n", cmd)

		os.Exit(1)
	}

	tickApp.Run(os.Args)
}
