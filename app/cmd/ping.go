package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pingCmd)
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping",
	Run: func(cmd *cobra.Command, args []string) {
		ping(settings.Uri)
	},
}

func ping(rpc string) {
	uri := fmt.Sprintf("%s/ping", rpc)
	fmt.Println("ping uri", uri)
	resp, err := http.Get(uri)
	if err != nil {
		log.Print("no resp")
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Print("parse body error")
		panic(err)
	}

	log.Print(string(body))
}
