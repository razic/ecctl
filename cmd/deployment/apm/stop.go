// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cmdapm

import (
	sdkcmdutil "github.com/elastic/cloud-sdk-go/pkg/util/cmdutil"
	"github.com/spf13/cobra"

	cmdutil "github.com/elastic/ecctl/cmd/util"
	"github.com/elastic/ecctl/pkg/deployment/apm"
	"github.com/elastic/ecctl/pkg/ecctl"
	"github.com/elastic/ecctl/pkg/util"
)

var stopApmCmd = &cobra.Command{
	Use:     "stop <apm deployment id>",
	Short:   "Stops an APM deployment",
	PreRunE: sdkcmdutil.MinimumNArgsAndUUID(1),
	Aliases: []string{"shutdown"},
	RunE: func(cmd *cobra.Command, args []string) error {
		track, _ := cmd.Flags().GetBool("track")
		hide, _ := cmd.Flags().GetBool("hide")
		return apm.Shutdown(apm.ShutdownParams{
			API:  ecctl.Get().API,
			ID:   args[0],
			Hide: hide,
			TrackParams: util.TrackParams{
				Track:  track,
				Output: ecctl.Get().Config.OutputDevice,
			},
		})
	},
}

func init() {
	stopApmCmd.Flags().Bool("hide", false, "Hides the deployment from the user")
	stopApmCmd.Flags().Bool("track", false, cmdutil.TrackFlagMessage)
}
