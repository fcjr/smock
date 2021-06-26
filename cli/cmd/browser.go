/*
Copyright © 2021 Left Shift Logical, LLC. <support@leftshift.com>

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
	"image"
	"os"

	_ "image/jpeg"
	"image/png"
	_ "image/png"

	"github.com/fcjr/smock"
	"github.com/spf13/cobra"
)

var browserCmd = &cobra.Command{
	Use:   "browser IN-FILE OUT-FILE",
	Short: "Create a Chrome-esque mockups",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		verticalPadding, err := cmd.Flags().GetInt("vertical-padding")
		cobra.CheckErr(err)
		horizontalPadding, err := cmd.Flags().GetInt("horizontal-padding")
		cobra.CheckErr(err)
		bgColorStr, err := cmd.Flags().GetString("background-color")
		cobra.CheckErr(err)
		bgAlpha, err := cmd.Flags().GetUint8("background-alpha")
		cobra.CheckErr(err)

		backgroundColor, err := smock.ParseHexColor(bgColorStr)
		cobra.CheckErr(err)
		backgroundColor.A = bgAlpha

		in, err := os.Open(args[0])
		cobra.CheckErr(err)
		defer in.Close()

		inImg, _, err := image.Decode(in)
		cobra.CheckErr(err)

		browserImg, err := smock.Browser(
			inImg,
			&smock.BrowserOptions{
				VerticalPadding:   verticalPadding,
				HorizontalPadding: horizontalPadding,
				BackgroundColor:   backgroundColor,
			},
		)
		cobra.CheckErr(err)

		outFile, err := os.OpenFile(args[1], os.O_RDWR|os.O_CREATE, 0640)
		cobra.CheckErr(err)
		defer outFile.Close()

		err = png.Encode(outFile, browserImg)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(browserCmd)

	browserCmd.PersistentFlags().Int("vertical-padding", 75, "vertical padding in px")
	browserCmd.PersistentFlags().Int("horizontal-padding", 100, "horizontal padding in px")
	browserCmd.PersistentFlags().String("background-color", "FFFFFF", "background color in hex")
	browserCmd.PersistentFlags().Uint8("background-alpha", 255, "background alpha")
}
