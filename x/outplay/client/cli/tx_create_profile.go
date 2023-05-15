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

func CmdCreateProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-profile [name] [date-of-birth] [playing-hand] [ntrp-rating]",
		Short: "Broadcast message create-profile",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDateOfBirth := args[1]
			argPlayingHand := args[2]
			argNtrpRating := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProfile(
				clientCtx.GetFromAddress().String(),
				argName,
				argDateOfBirth,
				argPlayingHand,
				argNtrpRating,
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
