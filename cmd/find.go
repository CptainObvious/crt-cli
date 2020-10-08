/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cptainobvious/crt-cli/client"
	"github.com/cptainobvious/crt-cli/model"
	"github.com/cptainobvious/crt-cli/utils"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Retrieve subdomains based on crt.sh",
	Long: `This command retrieve subdomains of a domain based on ssl certificate transparency:

It send a request to https://crt.sh/ with the query %.domainName to retrieve all subdomains certificate
`,
	Run: func(cmd *cobra.Command, args []string) {
		domainName, err := utils.GetDomainName(args[0])
		if err != nil {
			panic(err)
		}
		domain := &model.Domain{Name: domainName}
		c := client.NewCrtClient(model.JsonFormat)
		domains, err := c.GetSubDomains(domain)
		for _, d := range domains {
			fmt.Println(fmt.Sprintf("%s", d.GetName()))
		}
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
