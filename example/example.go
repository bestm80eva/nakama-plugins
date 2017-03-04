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

// Package example shows simple usage on various available functions in a
// Nakama plugin.
package example

import (
	"fmt"
	"log"
)

// These fields are set at compile.
var (
	version  string
	commitID string
)

// Event is called when a custom event is received by the server. The event will
// be deserialized from JSON and passed into this function. Events are rate
// limited by the server and may be dropped to ensure QoS.
func Event(logger *log.Logger, session map[string]string, event map[string]string) {
	logger.Printf("Event for '%s' with content: %+v", session["user_id"], event)
}

// Login is called when a user logs into the server.
func Login(logger *log.Logger, session map[string]string, user map[string]string) {
	logger.Printf("User '%s' logged in with properties: %+v", session["user_id"], user)
}

// RawReceipt is called when an IAP purchase has been made by a user. This
// function is called before the receipt is validated with the IAP provider or
// checked by the server for replay attacks, etc.
func RawReceipt(logger *log.Logger, session map[string]string, receipt map[string]string) {
	logger.Printf("Purchase by '%s' with receipt: %+v", session["user_id"], receipt)
}

// Register is called when a user creates an account with the server. If the
// plugin is activated after some users have already registered this function
// will not be retroactively called for those users.
func Register(logger *log.Logger, session map[string]string, user map[string]string) {
	logger.Printf("New user '%s' registered with properties: %+v", session["user_id"], user)
}

// SessionStart is called when a game client connects. The user who connected is
// passed into this function with additional event details. This function may be
// called multiple times for the same user as they can be connected via multiple
// game devices.
func SessionStart(logger *log.Logger, session map[string]string, startedAt int64, event map[string]string) {
	logger.Printf("User '%s' connected at '%d', properties: %+v", session["user_id"], startedAt, event)
}

// SessionStop is called when a game client disconnects. The user who disconnected
// is passed into this function with additional event details. This function may
// be called multiple times for the same user as they can be connected via
// multiple game devices.
func SessionStop(logger *log.Logger, session map[string]string, stoppedAt int64, event map[string]string) {
	logger.Printf("User '%s' disconnected at '%d', properties: %+v", session["user_id"], stoppedAt, event)
}

// Start is called when a plugin is loaded by the server. If load fails the
// plugin will never be loaded.
func Start(logger *log.Logger, config map[string]string) (message string, err error) {
	return fmt.Sprintf("Example plugin '%s+%s' loaded: %+v", version, commitID, config), nil
}

// Stop is called when a plugin must shutdown due to graceful termination of
// the server. This can be used to cleanup active resources.
func Stop(logger *log.Logger, config map[string]string) {
	logger.Printf("Example plugin stopped: %+v", config)
}

// Unload is called when the plugin encountered an error and must be forcibly
// unloaded by the server. This cannot be undone.
func Unload(logger *log.Logger, config map[string]string) {
	logger.Printf("Example plugin unloaded: %+v", config)
}

// UserUpdate is called when a user has updated their profile. A complete copy
// of the user is sent to this function.
func UserUpdate(logger *log.Logger, session map[string]string, user map[string]string) {
	logger.Printf("User '%s' updated, properties: %+v", session["user_id"], user)
}

// ValidReceipt is called when an IAP has been made by a user. This function is
// called after the receipt is validated with the IAP provider and has been
// checked by the server for replay attacks, etc.
func ValidReceipt(logger *log.Logger, session map[string]string, receipt map[string]string) {
	logger.Printf("Valid purchase by '%s' with receipt: %+v", session["user_id"], receipt)
}
