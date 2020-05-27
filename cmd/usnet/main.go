package usnet

import (
	"github.com/reinnet/topology/cmd/common"
	"github.com/reinnet/topology/usnet"
	"github.com/spf13/cobra"
)

func main() error {
	bld, err := usnet.New()
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
		Use:   "usnet",
		Short: "Build usnet topology",
		RunE: func(cmd *cobra.Command, args []string) error {
			return main()
		},
	}

	root.AddCommand(cmd)
}
