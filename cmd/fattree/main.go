package fattree

import (
	"github.com/reinnet/topology/cmd/common"
	"github.com/reinnet/topology/fattree"
	"github.com/spf13/cobra"
)

func main(k int) error {
	bld, err := fattree.New(k)
	if err != nil {
		return err
	}

	if err := common.Write(bld.Build()); err != nil {
		return err
	}

	return nil
}

// Register fattree command.
func Register(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "fattree",
		Short: "Build fattree topology",
		RunE: func(cmd *cobra.Command, args []string) error {
			k, err := cmd.Flags().GetInt("k")
			if err != nil {
				return err
			}

			return main(k)
		},
	}

	cmd.Flags().IntP("k", "k", 4, "k of k-FatTree")

	root.AddCommand(cmd)
}
