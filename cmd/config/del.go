package config

import (
	"bytes"
	"encoding/json"
	"slices"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var allowedDel = []string{"author"}

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Return to default for specific option",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		isAttrAllowed := slices.IndexFunc(allowedDel, func(el string) bool { return el == key }) != -1

		if !isAttrAllowed {
			return
		}

		// Yes, Viper is awful
		// https://github.com/spf13/viper/issues/632#issuecomment-493339494

		configMap := viper.AllSettings()
		delete(configMap, key)
		encodedConfig, _ := json.Marshal(configMap)
		viper.ReadConfig(bytes.NewReader(encodedConfig))
		viper.WriteConfig()
	},
}
