package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"outplay/x/outplay/types"
)

var _ = strconv.Itoa(0)

func CmdCreateChallenge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-challenge [opponent] [stake]",
		Short: "Broadcast message create-challenge",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOpponent := args[0]
			argStake := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateChallenge(
				clientCtx.GetFromAddress().String(),
				argOpponent,
				argStake,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
