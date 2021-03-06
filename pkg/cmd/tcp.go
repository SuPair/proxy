package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"proxy/pkg/provider"
)

func NewTcpCmd()*cobra.Command{
	var pro = provider.NewTcpProvider()
	o:=provider.TCPArgs{}
	// httpCmd represents the http command
	tcpCmd := &cobra.Command{
		Use:   "tcp",
		Short: "run a tcp proxy server",
		Long: `run a tcp proxy server`,
		Run: func(cmd *cobra.Command, args []string) {
			pro.Start(o)
			fmt.Println("tcp proxy stop")
		},
	}

	o.Args = baseArgs

	tcpCmd.Flags().StringVarP(&o.ParentType,"parent-type","T","tcp","parent type use tls | tcp")
	tcpCmd.Flags().IntVar(&o.Timeout,"timeout",2000,"tcp timeout milliseconds when connect to real server or parent proxy")
	tcpCmd.Flags().BoolVar(&o.IsTLS,"tls",false,"if not tls proxy")
	tcpCmd.Flags().IntVarP(&o.PoolSize,"pool-size","L",20,"conn pool size , which connect to parent proxy, zero: means turn off pool")
	tcpCmd.Flags().IntVarP(&o.CheckParentInterval,"check-parent-interval","I",3,"check if proxy is okay every interval seconds,zero: means no check")

	return tcpCmd
}
