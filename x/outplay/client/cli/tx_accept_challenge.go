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

func CmdAcceptChallenge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept-challenge [challenge-id]",
		Short: "Broadcast message accept-challenge",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChallengeId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAcceptChallenge(
				clientCtx.GetFromAddress().String(),
				argChallengeId,
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