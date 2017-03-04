// Copyright 2017 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// These test cases use the logger to verify behaviour in the example functions.
package example

import (
	"bytes"
	"log"
	"testing"
)

var (
	config  = map[string]string{}
	event   = map[string]string{}
	session = map[string]string{
		"user_id": "a",
		"handle":  "b",
	}
	user = map[string]string{}
)

func TestEvent(t *testing.T) {
	t.Parallel()
	buffer := new(bytes.Buffer)
	logger := log.New(buffer, "", 0)

	Event(logger, session, event)
	expected := "Event for 'a' with content: map[]\n"
	actual := buffer.String()
	if expected != actual {
		t.Errorf("got '%s', want '%s'", actual, expected)
	}
}

func TestLogin(t *testing.T) {
	t.Parallel()
	buffer := new(bytes.Buffer)
	logger := log.New(buffer, "", 0)

	Login(logger, session, user)
	expected := "User 'a' logged in with properties: map[]\n"
	actual := buffer.String()
	if expected != actual {
		t.Errorf("got '%s', want '%s'", actual, expected)
	}
}

func TestRawReceipt(t *testing.T) {
	t.Parallel()
	buffer := new(bytes.Buffer)
	logger := log.New(buffer, "", 0)

	receipt := map[string]string{}
	RawReceipt(logger, session, receipt)
	expected := "Purchase by 'a' with receipt: map[]\n"
	actual := buffer.String()
	if expected != actual {
		t.Errorf("got '%s', want '%s'", actual, expected)
	}
}

func TestRegister(t *testing.T) {
	t.Parallel()
	buffer := new(bytes.Buffer)
	logger := log.New(buffer, "", 0)

	Register(logger, session, user)
	expected := "New user 'a' registered with properties: map[]\n"
	actual := buffer.String()
	if expected != actual {
		t.Errorf("got '%s', want '%s'", actual, expected)
	}
}

func TestSessionStart(t *testing.T) {
	t.Parallel()
	buffer := new(bytes.Buffer)
	logger := log.New(buffer, "", 0)

	startedAt := int64(1488642908246)
	SessionStart(logger, session, startedAt, event)
	expected := "User 'a' connected at '1488642908246', properties: map[]\n"
	actual := buffer.String()
	if expected != actual {
		t.Errorf("got '%s', want '%s'", actual, expected)
	}
}

func TestSessionStop(t *testing.T) {
	t.Parallel()
	buffer := new(bytes.Buffer)
	logger := log.New(buffer, "", 0)

	startedAt := int64(1488642908246)
	SessionStop(logger, session, startedAt, event)
	expected := "User 'a' disconnected at '1488642908246', properties: map[]\n"
	actual := buffer.String()
	if expected != actual {
		t.Errorf("got '%s', want '%s'", actual, expected)
	}
}

func TestStart(t *testing.T) {
	t.Parallel()
	buffer := new(bytes.Buffer)

	result, err := Start(log.New(buffer, "", 0), config)
	expected := "Example plugin '+' loaded: map[]"
	if err != nil || result != expected {
		t.Errorf("got '%s', want '%s'", result, expected)
	}
}

func TestStop(t *testing.T) {
	t.Parallel()
	buffer := new(bytes.Buffer)
	logger := log.New(buffer, "", 0)

	Stop(logger, config)
	expected := "Example plugin stopped: map[]\n"
	actual := buffer.String()
	if expected != actual {
		t.Errorf("got '%s', want '%s'", actual, expected)
	}
}

func TestUnload(t *testing.T) {
	t.Parallel()
	buffer := new(bytes.Buffer)
	logger := log.New(buffer, "", 0)

	Unload(logger, config)
	expected := "Example plugin unloaded: map[]\n"
	actual := buffer.String()
	if expected != actual {
		t.Errorf("got '%s', want '%s'", actual, expected)
	}
}

func TestUserUpdate(t *testing.T) {
	t.Parallel()
	buffer := new(bytes.Buffer)
	logger := log.New(buffer, "", 0)

	UserUpdate(logger, session, user)
	expected := "User 'a' updated, properties: map[]\n"
	actual := buffer.String()
	if expected != actual {
		t.Errorf("got '%s', want '%s'", actual, expected)
	}
}

func TestValidReceipt(t *testing.T) {
	t.Parallel()
	buffer := new(bytes.Buffer)
	logger := log.New(buffer, "", 0)

	ValidReceipt(logger, session, user)
	expected := "Valid purchase by 'a' with receipt: map[]\n"
	actual := buffer.String()
	if expected != actual {
		t.Errorf("got '%s', want '%s'", actual, expected)
	}
}
