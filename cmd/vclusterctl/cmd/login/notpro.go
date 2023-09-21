//go:build !pro
// +build !pro

package login

import (
	"github.com/loft-sh/vcluster/cmd/vclusterctl/flags"
	"github.com/loft-sh/vcluster/pkg/constants"
	"github.com/spf13/cobra"
)

func NewLoginCmd(*flags.GlobalFlags) (*cobra.Command, error) {
	return nil, constants.ErrOnlyInPro
}
