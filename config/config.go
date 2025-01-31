/*
Copyright © 2022 Anil Chauhan <https://github.com/meanii>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package config

import (
	"log"
	"os"
	"path/filepath"
)

func GetUserDBPath() string {
	userDb := "/.user.json"
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("something went wrong while getting $HOME! Reason %v", err)
	}
	configPath := filepath.Join(dirname, ".meanii/sync/config")
	return configPath + userDb
}

func GetSymlinkPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("something went wrong while getting $HOME! Reason %v", err)
	}
	configPath := filepath.Join(dirname, ".meanii/sync/symlink")
	return configPath
}

func GetSyncDBPath() string {
	syncDb := "/.sync.json"
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("something went wrong while getting $HOME! Reason %v", err)
	}
	configPath := filepath.Join(dirname, ".meanii/sync")
	return configPath + syncDb
}

func GetWorkingPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("something went wrong while getting $HOME! Reason %v", err)
	}
	working := filepath.Join(dirname, ".meanii/sync")
	return working
}

func GetHistoryPath() string {
	historyDbPath := "/.history.json"
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("something went wrong while getting $HOME! Reason %v", err)
	}
	history := filepath.Join(dirname, ".meanii/sync")
	return history + historyDbPath
}
