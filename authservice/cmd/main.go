package main

import (
	"github.com/spf13/cobra"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/app"
	_ "github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent/runtime"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "authservice",
		Short: "Start the Auth Service",
		Run: func(cmd *cobra.Command, args []string) {
			app.Run()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
