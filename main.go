package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	kiloclient "github.com/squat/kilo/pkg/k8s/clientset/versioned"
)

func main() {
	cmd := &cobra.Command{
		Use:   "kilosubspace",
		Args:  cobra.ArbitraryArgs,
		Short: "Use Kilo as a backend for Subspace",
		Long:  "",
	}
	var kubeconfig string
	cmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", os.Getenv("KUBECONFIG"), "Path to kubeconfig.")
	var bin string
	cmd.PersistentFlags().StringVar(&bin, "wg", os.Getenv("KILOSUBSPACE_WG"), "Path to wg binary.")

	var c kubernetes.Interface
	var kc kiloclient.Interface
	cmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return fmt.Errorf("failed to create Kubernetes config: %v", err)
		}
		c = kubernetes.NewForConfigOrDie(config)
		kc = kiloclient.NewForConfigOrDie(config)

		return nil
	}
	cmd.RunE = wg(&bin)
	for _, subCmd := range []*cobra.Command{
		findKey(&c),
		set(&kc),
	} {
		cmd.AddCommand(subCmd)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
