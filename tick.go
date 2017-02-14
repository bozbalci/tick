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

// XXX This program is currently a prototype.

package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/urfave/cli"
)

const prefix = "."

type Track struct {
    name string
    path string
}

func buildTrackPath(prefix, trackName string) string {
	return path.Join(prefix, trackName + ".csv")
}

func tickFormatTime(date time.Time) string {
    // This is actually the idiomatic way to do it. Unbelievable.

    return date.Format("2006-01-02")
}

func trackExists(trackName string) bool {
    trackPath := buildTrackPath(prefix, trackName)
	_, err := os.Stat(trackPath)

	return !os.IsNotExist(err)
}

func createTrack(trackName string) error {
	if trackName == "" {
		return errors.New("tick: no track name specified")
	}

	trackPath := buildTrackPath(prefix, trackName)

	if trackExists(trackPath) {
		return errors.New(fmt.Sprintf("tick: track %q already exists", trackName))
	}

	newFile, err := os.Create(trackPath)
	if err != nil {
		return errors.New("tick: " + err.Error())
	}

	fmt.Printf("Created track %q\n", trackName)
	newFile.Close()
	return nil
}

func deleteTrack(trackName string) error {
	if trackName == "" {
		return errors.New("tick: no track name specified")
	}

	trackPath := buildTrackPath(prefix, trackName)

	if !trackExists(trackName) {
		return errors.New(fmt.Sprintf("tick: track %q does not exist", trackName))
	}

	err := os.Remove(trackPath)

	if err != nil {
		return errors.New("tick: " + err.Error())
	}

	fmt.Printf("Deleted track %q\n", trackName)
	return nil
}

func tickTrack(trackName string, date time.Time) error {
	// trackPath := buildTrackPath(prefix, trackName)

	if !trackExists(trackName) {
        return errors.New(fmt.Sprintf("tick: track %q does not exist", trackName))
	}

    fmt.Printf("ticked %q in %s\n", trackName, tickFormatTime(date))

	return nil
}

func tickToday(trackName string) error {
    return tickTrack(trackName, time.Now())
}

func main() {
	tickApp := cli.NewApp()
	tickApp.Name = "tick"
	tickApp.Usage = "One bit journal for the command-line"
	tickApp.Version = "0.0.1"

    // The default behavior of tick is to add a tick to the track with
    // time.Now() as the Tick time.
	tickApp.Action = func(c *cli.Context) error {
		err := tickToday(c.Args().First())

        return err
	}

	tickApp.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "create a new track",
			Action: func(c *cli.Context) error {
				err := createTrack(c.Args().First())

				return err
			},
		},
		{
			Name:  "delete",
			Usage: "delete an existing track",
			Action: func(c *cli.Context) error {
				err := deleteTrack(c.Args().First())

				return err
			},
		},
	}

	tickApp.Run(os.Args)
}
