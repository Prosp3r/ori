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
	"strconv"

	"github.com/spf13/cobra"
	ori "gitlab.com/Prosp3r/ori/pb"
	"google.golang.org/grpc"
)

// subtractCmd represents the subtract command
var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "Subtracts one or more space separated number(s) from the other",
	Long: `Subtracts one or more space separated number(s) from the other. 
	For example: subtract 8 4 ie. 8-4`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(ServerLink, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect to server with error: %v", err)
		}
		defer conn.Close()
		Client = ori.NewORIServiceClient(conn)

		argx := make([]float64, 0)
		for _, v := range args {
			val, err := strconv.ParseFloat(v, 64)
			if err != nil {
				log.Fatalf("Error converting string to int %v \n", err)
			}
			argx = append(argx, val)
		}

		var result float64
		for i, v := range argx {
			if i < 1 {
				valA6 := v
				valB6 := argx[i+1]
				res := SubtractRequest(Client, valA6, valB6)
				result = res
			} else {
				if i < len(argx)-1 {
					valB6 := argx[i+1] //[ 7  2  1  2] = 2
					res := SubtractRequest(Client, result, valB6)
					result = res
				}
			}
		}

		fmt.Printf("Result of the subtraction is: %v \n", result)
		//fmt.Println("subtract called")

	},
}

//SubtractRequest will return sum operations of two numbers from the grpc server
func SubtractRequest(clientConn ori.ORIServiceClient, valA, valB float64) float64 {

	req := &ori.ORISutractRequest{
		A: valA,
		B: valB,
	}

	r, err := clientConn.Sutract(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not calculate due to error %v \n", err)
	}

	return r.Result
}

func init() {
	rootCmd.AddCommand(subtractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subtractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subtractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
