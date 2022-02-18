/*
Package cmd
Copyright © 2022 Mehdi Teymorian

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
	"github.com/mehditeymorian/hermes/internal/cmd/serve"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var cfgFile string

	rootCmd := &cobra.Command{ //nolint:exhaustivestruct
		Use:   "hermes",
		Short: "WebRTC Signaling Server",
		Long: `Hermes is WebRTC signaling server. 
			it provides room management, event handling, and signaling.`,
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file (default is $HOME/.hermes.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(serve.Command(cfgFile))

	cobra.CheckErr(rootCmd.Execute())
}
