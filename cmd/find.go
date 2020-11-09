package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cptainobvious/crt-cli/client"
	"github.com/cptainobvious/crt-cli/model"
	"github.com/cptainobvious/crt-cli/utils"
)

var (
	blacklist []string
	findCmd   = &cobra.Command{
		Use:   "find domain",
		Short: "Retrieve subdomains based on crt.sh",
		Example: "crt-cli find domain.net",
		Long: `This command retrieve subdomains of a domain based on ssl certificate transparency:
It send a request to https://crt.sh/ with the query %.domainName to retrieve all subdomains certificate
`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			domainName, err := utils.GetDomainName(args[0])
			if err != nil {
				panic(err)
			}
			domain := &model.Domain{Name: domainName}

			c := client.NewCrtClient(model.JsonFormat, client.NewHttpClient())
			blacklist, err := model.NewBlacklist(blacklist)
			if err != nil {
				panic(err)
			}
			domains, err := c.WithBlacklist(blacklist).GetSubDomains(domain)
			if err != nil {
				panic(err)
			}
			for _, d := range domains {
				fmt.Println(fmt.Sprintf("Name: %s, Alive %t", d.GetName(), d.IsAlive()))
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.Flags().StringSliceVarP(&blacklist, "blacklist", "b", []string{}, "")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
