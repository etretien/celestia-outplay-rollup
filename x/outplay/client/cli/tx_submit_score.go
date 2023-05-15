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

func CmdSubmitScore() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-score [challenge-id] [score]",
		Short: "Broadcast message submit-score",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChallengeId := args[0]
			argScore := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitScore(
				clientCtx.GetFromAddress().String(),
				argChallengeId,
				argScore,
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
