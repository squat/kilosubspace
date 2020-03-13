package main

import (
	"crypto/sha256"
	"fmt"
	"net"
	"strings"

	"github.com/spf13/cobra"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/squat/kilo/pkg/k8s/apis/kilo/v1alpha1"
	kiloclient "github.com/squat/kilo/pkg/k8s/clientset/versioned"
)

const (
	removeLength = 4
	setLength    = 5
)

func set(kc *kiloclient.Interface) *cobra.Command {
	return &cobra.Command{
		Use:   "set",
		Short: "Set the WireGuard configuration for a peer.",
		Args:  cobra.RangeArgs(removeLength, setLength),
		RunE: func(_ *cobra.Command, args []string) error {
			key := args[2]
			h := sha256.New()
			if _, err := h.Write([]byte(key)); err != nil {
				return fmt.Errorf("hash key: %w", err)
			}
			name := fmt.Sprintf("%x", h.Sum(nil))

			if args[len(args)-1] == "remove" {
				if err := (*kc).KiloV1alpha1().Peers().Delete(name, &metav1.DeleteOptions{}); err != nil && !kerrors.IsNotFound(err) {
					return err
				}
				return nil
			}
			if len(args) != setLength {
				return fmt.Errorf("expected %d args, got %d", setLength, len(args))
			}

			var aips []string
			for _, a := range strings.Split(args[4], ",") {
				a = strings.TrimSpace(a)
				if _, _, err := net.ParseCIDR(a); err != nil {
					return fmt.Errorf("parse allowed IP %s: %w", a, err)
				}
				aips = append(aips, a)
			}

			p := &v1alpha1.Peer{
				ObjectMeta: metav1.ObjectMeta{
					Name: name,
					Labels: map[string]string{
						"app.kubernetes.io/managed-by": "subspace",
					},
				},
				Spec: v1alpha1.PeerSpec{
					PublicKey:           key,
					AllowedIPs:          aips,
					PersistentKeepalive: 10,
				},
			}

			_, err := (*kc).KiloV1alpha1().Peers().Create(p)
			return err
		},
	}
}
