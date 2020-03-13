package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"

	"github.com/squat/kilo/pkg/k8s"
)

func findKey(c *kubernetes.Interface) *cobra.Command {
	return &cobra.Command{
		Use:   "findkey",
		Short: "Find the WireGuard public key corresponding to an endpoint.",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			endpoint := args[0]
			b := k8s.New(*c, nil, nil)
			stop := make(chan struct{})
			defer close(stop)
			if err := b.Nodes().Init(stop); err != nil {
				return fmt.Errorf("initializing backend: %w", err)
			}
			nodes, err := b.Nodes().List()
			if err != nil {
				return fmt.Errorf("listing nodes: %w", err)
			}
			for _, n := range nodes {
				if n.Endpoint.String() == endpoint {
					fmt.Printf("%s", n.Key)
					return nil
				}
			}
			return fmt.Errorf("could not find public key for endpoint %q", endpoint)
		},
	}
}
