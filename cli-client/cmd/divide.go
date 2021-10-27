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

// divideCmd represents the divide command
var divideCmd = &cobra.Command{
	Use:   "divide",
	Short: "Divide will return the division of two or more space separated numbers",
	Long: `Divide will return the division of two or more space separated numbers. 
	For example:

divide 2 2 will return 1 and divide 12 3 2 is the same expression as (12/3)/2 which will return 2.`,
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
				res := DivideRequest(Client, valA6, valB6)
				result = res
			} else {
				if i < len(argx)-1 {
					valB6 := argx[i+1] //[ 7  2  1  2] = 2
					res := DivideRequest(Client, result, valB6)
					result = res
				}
			}
		}

		fmt.Printf("Result of the division is: %v \n", result)
		//fmt.Println("divide called")
	},
}

//DivideRequest will return division of two numbers from the grpc server
func DivideRequest(clientConn ori.ORIServiceClient, valA, valB float64) float64 {
	req := &ori.ORIDivideRequest{
		A: valA,
		B: valB,
	}

	r, err := clientConn.Divide(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not calculate due to error %v \n", err)
	}

	return r.Result
}

func init() {
	rootCmd.AddCommand(divideCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divideCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divideCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
