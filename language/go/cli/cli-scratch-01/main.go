package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)
type Topic struct {
	 name string
	 details string
}
var topics []Topic
var (
	get bool
	rootCmd = &cobra.Command{
		Use:   "example",
		Short: "use to get example",
		Long:  "long description",
	}
	topicCmd = &cobra.Command{
		Use: "topic",
		Short: "use to work on topic",
		Long: "use topic to get, search and attach topics to your wishlist",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("topics:", strings.Join(args, " "))
		},
	}

	detailCmd = &cobra.Command{
		Use: "detail",
		Short: "use to get detail on topic",
		Long: "use detail to get detail description of a topics",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("detail:", args)
			fmt.Println(cmd.Flag("name").Value)
			fmt.Println(cmd.Flag("get").Value)
		},
	}

)



func init()  {
	topics = []Topic{{"tech","tech aka future"}, {"health","GE"}  }
	rootCmd.PersistentFlags().BoolVarP(&get, "get", "g", false, "get examples or topic")
	rootCmd.AddCommand(topicCmd)

	detailCmd.Flags().StringP("name", "n", "telco", "name of the topic")
	topicCmd.AddCommand(detailCmd)
}



func main() {
	fmt.Println(topics)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}