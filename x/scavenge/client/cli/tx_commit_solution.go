package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"github.com/batphonghan/blog/x/scavenge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCommitSolution() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit-solution [solution-hash] [solution-scavenger-hash]",
		Short: "Broadcast message commit-solution",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			solution := args[0]
			// find a hash of the solution
			solutionHash := sha256.Sum256([]byte(solution))
			// convert the solution hash to string
			solutionHashString := hex.EncodeToString(solutionHash[:])
			// convert a scavenger address to string
			var scavenger = clientCtx.GetFromAddress().String()
			var solutionScavengerHash = sha256.Sum256([]byte(solution + scavenger))
			// convert the hash to string
			var solutionScavengerHashString = hex.EncodeToString(solutionScavengerHash[:])

			msg := types.NewMsgCommitSolution(clientCtx.GetFromAddress().String(), string(solutionHashString), string(solutionScavengerHashString))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
