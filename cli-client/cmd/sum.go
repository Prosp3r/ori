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

//var sumRequest *ori.ORISumRequest

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "Sums two or more numbers separated by space",
	Long: `Adds two or more numbers separated by space and prints the result. For example:
	sum 10 40 50 30`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Args received %v", args)
		//INITIATE CONNECTION
		conn, err := grpc.Dial(ServerLink, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect to server with error: %v", err)
		}
		defer conn.Close()
		Client = ori.NewORIServiceClient(conn)

		result := float64(0)
		for x := 0; x < len(args); x++ {
			valB, err := strconv.ParseFloat(args[x], 64)
			if err != nil {
				log.Fatalf("Error converting string to int %v \n", err)
			}
			//valB6 := int64(valB)
			res := ArithmeticRequest(Client, result, valB)
			result = res
		}

		fmt.Printf("Result of the sum is: %v \n", result)

		//fmt.Println("sum called")
	},
}

//ArithmeticRequest will return sum operations of two numbers from the grpc server
func ArithmeticRequest(clientConn ori.ORIServiceClient, valA, valB float64) float64 {
	req := &ori.ORISumRequest{
		A: valA,
		B: valB,
	}

	r, err := clientConn.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not calculate due to error %v \n", err)
	}

	return r.Result
}

func init() {
	rootCmd.AddCommand(sumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
