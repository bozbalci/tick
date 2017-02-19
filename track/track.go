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

package track

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

const PREFIX = "."

type Track struct {
	Name string
	Path string
}

func (t *Track) buildPath(prefix string) {
	t.Path = path.Join(prefix, t.Name)
}

func (t Track) Exists() bool {
	if t.Path == "" {
		return false
	}

	_, err := os.Stat(t.Path)

	return !os.IsNotExist(err)
}

func (t Track) Create() error {
	if t.Name == "" {
		return errors.New("tick: no track name specified")
	}

	if t.Exists() {
		return fmt.Errorf("tick: track %q already exists", t.Name)
	}

	newFile, err := os.Create(t.Path)
	if err != nil {
		return fmt.Errorf("tick: %s", err.Error())
	}

	newFile.Close()

	return nil
}

func (t Track) Delete() error {
	if t.Name == "" {
		return errors.New("tick: no track name specified")
	}

	if !t.Exists() {
		return fmt.Errorf("tick: track %q does not exist", t.Name)
	}

	err := os.Remove(t.Path)
	if err != nil {
		return fmt.Errorf("tick: %s", err.Error())
	}

	return nil
}

func (t Track) Tick(date time.Time) error {
	if t.Name == "" {
		return errors.New("tick: no track name specified")
	}

	if !t.Exists() {
		return fmt.Errorf("tick: track %q does not exist", t.Name)
	}

	formattedDate := date.Format("2006-01-02")

	f, err := os.OpenFile(t.Path, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		return fmt.Errorf("tick: %s", err.Error())
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), formattedDate) {
			return fmt.Errorf("tick: track %q is already ticked on %s", t.Name, formattedDate)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("tick: %s", err.Error())
	}

	writer := bufio.NewWriter(f)
	fmt.Fprintln(writer, formattedDate)
	writer.Flush()

	return nil
}

func (t Track) TickToday() error {
	return t.Tick(time.Now())
}

func Correlate(t1, t2 Track) float32 {
	// TODO: Implement track.Correlate

	return 0
}

func New(name string) *Track {
	t := &Track{Name: name}
	t.buildPath(PREFIX)

	return t
}
