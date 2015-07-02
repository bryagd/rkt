// Copyright 2015 The rkt Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/coreos/rkt/store"

	"github.com/coreos/rkt/Godeps/_workspace/src/github.com/appc/spec/discovery"
	"github.com/coreos/rkt/Godeps/_workspace/src/github.com/appc/spec/schema/types"
	"github.com/coreos/rkt/Godeps/_workspace/src/github.com/spf13/cobra"
)

var (
	cmdImage = &cobra.Command{
		Use:   "image [command]",
		Short: "Operate on image(s) in the local store",
	}
)

func init() {
	cmdRkt.AddCommand(cmdImage)
}

func getKeyFromAppOrHash(s *store.Store, input string) (string, error) {
	var key string
	if _, err := types.NewHash(input); err == nil {
		key, err = s.ResolveKey(input)
		if err != nil {
			return "", fmt.Errorf("cannot resolve key: %v", err)
		}
	} else {
		app, err := discovery.NewAppFromString(input)
		if err != nil {
			return "", fmt.Errorf("cannot parse the image name: %v", err)
		}
		labels, err := types.LabelsFromMap(app.Labels)
		if err != nil {
			return "", fmt.Errorf("invalid labels in the name: %v", err)
		}
		key, err = s.GetACI(app.Name, labels)
		if err != nil {
			return "", fmt.Errorf("cannot find image: %v", err)
		}
	}

	return key, nil
}
