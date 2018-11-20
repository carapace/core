package cmd

// import (
// 	"fmt"
// 	"github.com/carapace/core/cmd/v0/carapace"
// 	"github.com/spf13/cobra"
// 	"net"
// )
//
// var port string
//
// // serveCmd represents the serve command
// var serveCmd = &cobra.Command{
// 	Use:   "serve",
// 	Short: "Start the carapace node",
// 	Long: `Start the carapace node. The node is self configuring, using environment variables during it's startup cycle
//
// 			`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		app, server, err := carapace.New()
// 		if err != nil {
// 			fmt.Println("internal error occurred while initializing server: ", err.Error())
// 		}
// 		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
// 		if err != nil {
// 			fmt.Printf("unable to listen on port %s\n port: %s: \n", port, err)
// 		}
// 		app.Start()
//
// 		fmt.Printf("serving carapace node on %s\n", port)
// 		err = server.Serve(lis)
// 		fmt.Printf("server shutdown due to: %s", err.Error())
// 	},
// }
//
// func init() {
// 	rootCmd.AddCommand(serveCmd)
// 	serveCmd.Flags().StringVar(&port, "port", "4000", "port to listen on")
// 	// Here you will define your flags and configuration settings.
//
// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")
//
// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
