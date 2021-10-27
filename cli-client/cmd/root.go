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
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	ori "gitlab.com/Prosp3r/ori/pb"
	"google.golang.org/grpc"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

//client context
var requestCtx context.Context

//ServerLink ...for localDev
//var ServerLink = "localhost:8080"

//ServerLink for production
var ServerLink = "oriserver:8080"

//Client ...
var Client ori.ORIServiceClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oriclient (sum | multiply | subtract | divide) [numberOne numberTwo]",
	Short: "A Command line client for accessing Ori grpc math server ",
	Long: `A Command line client for accessing Ori grpc math server.
	Oriclient is a command line interface for performing simple limited number 
	of arithmetic operations eg. sum, subtract, divide, multiply.
	(sum | multiply | subtract | divide) [numberOne numberTwo]`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oriclient.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//INitialize client
	//fmt.Println("Starting Ori math client")
	//OriColorize()
	//timeout if no response
	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	conn, err := grpc.Dial(ServerLink, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server with error: %v", err)
	}
	defer conn.Close()

	Client = ori.NewORIServiceClient(conn)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".oriclient" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".oriclient")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

/*
///AESTETICS

//Color type for color coded text messages for aestetics
type Color string

const (
	//ColorBlack ..
	ColorBlack Color = "\u001b[30m"
	//ColorRed ...for errors
	ColorRed = "\u001b[31m"
	//ColorGreen ...for successes
	ColorGreen = "\u001b[32m"
	//ColorYellow ...for warnings
	ColorYellow = "\u001b[33m"
	//ColorBlue ...for instructions
	ColorBlue = "\u001b[34m"
	//ColorReset ...
	ColorReset = "\u001b[0m"
)

//colorize ...for aestetics
func colorize(color Color, message string) {
	fmt.Print(string(color), message, string(ColorReset))
	//return string(color)
}

//OriColorize applies Ori Colors to string
func OriColorize(color []Color, message []string) error {
	for i, v := range color {
		colorize(v, message[i])
		//msg = append(msg, colorize(v, message[i]))
	}
	return nil
}*/
