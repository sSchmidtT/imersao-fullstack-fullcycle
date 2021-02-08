/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/sSchmidtT/imersao-fullstack-fullcycle/codepix/domain/model"
	"github.com/sSchmidtT/imersao-fullstack-fullcycle/codepix/infrastructure/db"
	"github.com/sSchmidtT/imersao-fullstack-fullcycle/codepix/infrastructure/repository"
	"github.com/spf13/cobra"
)

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "Run fixtures for fake data generation",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		fmt.Println("initialize db instance")
		defer database.Close()
		pixRepository := repository.PixKeyRepositoryDb{Db: database}
		fmt.Println("initialize pix repository")

		bankBBX, err := model.NewBank("001", "BBX")
		if err != nil {
			fmt.Println("error created bank: ", err.Error())
		}

		bankCTER, err := model.NewBank("002", "CTER")
		if err != nil {
			fmt.Println("error created bank: ", err.Error())
		}

		pixRepository.AddBank(bankBBX)
		pixRepository.AddBank(bankCTER)

		account1, err := model.NewAccount(bankBBX, "1111", "User BBX 1")
		if err != nil {
			fmt.Println("error created account: ", err.Error())
		}
		fmt.Println("create account ", account1)
		account1.ID = "6e4635ce-88d1-4e58-9597-d13fc446ee47"
		err = pixRepository.AddAccount(account1)
		if err != nil {
			fmt.Println("error add account ", err.Error())
		}

		account2, err := model.NewAccount(bankBBX, "2222", "User BBX 2")
		if err != nil {
			fmt.Println("error created account: ", err.Error())
		}
		fmt.Println("create account ", account2)
		account2.ID = "51a720b2-5144-4d7f-921d-57023b1e24c1"
		err = pixRepository.AddAccount(account2)
		if err != nil {
			fmt.Println("error add account ", err.Error())
		}

		account3, err := model.NewAccount(bankCTER, "3333", "User CTER 1")
		if err != nil {
			fmt.Println("error created account: ", err.Error())
		}
		fmt.Println("create account ", account3)
		account3.ID = "103cc632-78e7-4476-ab63-d5ad3a562d26"
		err = pixRepository.AddAccount(account3)
		if err != nil {
			fmt.Println("error add account ", err.Error())
		}

		account4, err := model.NewAccount(bankCTER, "4444", "User CTER 2")
		if err != nil {
			fmt.Println("error created account: ", err.Error())
		}
		fmt.Println("create account ", account4)
		account4.ID = "463b1b2a-b5fa-4b88-9c31-e5c894a20ae3"
		err = pixRepository.AddAccount(account4)
		if err != nil {
			fmt.Println("error add account ", err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixturesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixturesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
