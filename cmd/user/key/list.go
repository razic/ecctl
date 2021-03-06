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

package cmduserkey

import (
	"github.com/spf13/cobra"

	"github.com/elastic/ecctl/pkg/ecctl"
	userauthadmin "github.com/elastic/ecctl/pkg/user/auth/admin"
)

var listCmd = &cobra.Command{
	Use:     "list --user=<user id>|--all",
	Short:   "Lists the API keys for the specified user, or all platform users",
	PreRunE: cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		all, _ := cmd.Flags().GetBool("all")
		res, err := userauthadmin.ListKeys(userauthadmin.ListKeysParams{
			API:    ecctl.Get().API,
			All:    all,
			UserID: cmd.Flag("user").Value.String(),
		})
		if err != nil {
			return err
		}

		return ecctl.Get().Formatter.Format("user/keys-details", res)
	},
}

func init() {
	listCmd.Flags().String("user", "", "user id for the specified action")
	listCmd.Flags().Bool("all", false, "If specified, lists all API keys for all platform users")
}
