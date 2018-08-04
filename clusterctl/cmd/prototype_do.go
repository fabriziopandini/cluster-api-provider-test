/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

type DoOptions struct {
}

var co = &DoOptions{}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "prototype clusterctl extension for a sub-command",
	Long:  `prototype clusterctl extension for a sub-command`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := RunDo(co); err != nil {
			glog.Exit(err)
		}
	},
}

func RunDo(co *DoOptions) error {
	return nil
}

func init() {
	prototypeCmd.AddCommand(doCmd)
}
