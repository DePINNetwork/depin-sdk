package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	govutils "cosmossdk.io/x/gov/client/utils"
	"cosmossdk.io/x/gov/types"
	v1 "cosmossdk.io/x/gov/types/v1"
	"cosmossdk.io/x/gov/types/v1beta1"

	"github.com/depinnetwork/depin-sdk/client"
	"github.com/depinnetwork/depin-sdk/client/flags"
	"github.com/depinnetwork/depin-sdk/client/tx"
	sdk "github.com/depinnetwork/depin-sdk/types"
	"github.com/depinnetwork/depin-sdk/version"
)

// Proposal flags
const (
	FlagTitle     = "title"
	FlagDeposit   = "deposit"
	flagVoter     = "voter"
	flagDepositor = "depositor"
	flagStatus    = "status"
	FlagMetadata  = "metadata"
	FlagSummary   = "summary"
	FlagExpedited = "expedited"

	// Deprecated: only used for v1beta1 legacy proposals.
	FlagProposal = "proposal"
	// Deprecated: only used for v1beta1 legacy proposals.
	FlagDescription = "description"
	// Deprecated: only used for v1beta1 legacy proposals.
	FlagProposalType = "type"
)

// ProposalFlags defines the core required fields of a legacy proposal. It is used to
// verify that these values are not provided in conjunction with a JSON proposal
// file.
var ProposalFlags = []string{
	FlagTitle,
	FlagDescription,
	FlagProposalType,
	FlagDeposit,
}

// NewTxCmd returns the transaction commands for this module
// governance ModuleClient is slightly different from other ModuleClients in that
// it contains a slice of legacy "proposal" child commands. These commands are respective
// to the proposal type handlers that are implemented in other modules but are mounted
// under the governance CLI (eg. parameter change proposals).
func NewTxCmd(legacyPropCmds []*cobra.Command) *cobra.Command {
	govTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Governance transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmdSubmitLegacyProp := NewCmdSubmitLegacyProposal()
	for _, propCmd := range legacyPropCmds {
		flags.AddTxFlagsToCmd(propCmd)
		cmdSubmitLegacyProp.AddCommand(propCmd)
	}

	govTxCmd.AddCommand(
		NewCmdWeightedVote(),
		NewCmdSubmitProposal(),
		NewCmdDraftProposal(),

		// Deprecated
		cmdSubmitLegacyProp,
	)

	return govTxCmd
}

// NewCmdSubmitProposal implements submitting a proposal transaction command.
func NewCmdSubmitProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proposal <path/to/proposal.json>",
		Short: "Submit a proposal along with some messages, metadata and deposit",
		Args:  cobra.ExactArgs(1),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a proposal along with some messages, metadata and deposit.
They should be defined in a JSON file.

Example:
$ %s tx gov submit-proposal path/to/proposal.json

Where proposal.json contains:

{
  // array of proto-JSON-encoded sdk.Msgs
  "messages": [
    {
      "@type": "/cosmos.bank.v1beta1.MsgSend",
      "from_address": "cosmos1...",
      "to_address": "cosmos1...",
      "amount":[{"denom": "stake","amount": "10"}]
    }
  ],
  // metadata can be any of base64 encoded, raw text, stringified json, IPFS link to json
  // see below for example metadata
  "metadata": "4pIMOgIGx1vZGU=",
  "deposit": "10stake",
  "title": "My proposal",
  "summary": "A short summary of my proposal",
  "proposal_type": "standard",
}

metadata example: 
{
	"title": "",
	"authors": [""],
	"summary": "",
	"details": "", 
	"proposal_forum_url": "",
	"vote_option_context": "",
}
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposal, msgs, deposit, err := parseSubmitProposal(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}

			addr, err := clientCtx.AddressCodec.BytesToString(clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			msg, err := v1.NewMsgSubmitProposal(msgs, deposit, addr, proposal.Metadata, proposal.Title, proposal.Summary, proposal.proposalType)
			if err != nil {
				return fmt.Errorf("invalid message: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCmdSubmitLegacyProposal implements submitting a proposal transaction command.
// Deprecated: please use NewCmdSubmitProposal instead.
func NewCmdSubmitLegacyProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "submit-legacy-proposal",
		Aliases: []string{"exec-legacy-content"},
		Short:   "Submit a legacy proposal along with an initial deposit",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a legacy proposal along with an initial deposit.
Proposal title, description, type and deposit can be given directly or through a proposal JSON file.

Example:
$ %s tx gov submit-legacy-proposal --proposal="path/to/proposal.json" --from mykey

Where proposal.json contains:

{
  "title": "Test Proposal",
  "description": "My awesome proposal",
  "type": "Text",
  "deposit": "10test"
}

Which is equivalent to:

$ %s tx gov submit-legacy-proposal --title="Test Proposal" --description="My awesome proposal" --type="Text" --deposit="10test" --from mykey
`,
				version.AppName, version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposal, err := parseSubmitLegacyProposal(cmd.Flags())
			if err != nil {
				return fmt.Errorf("failed to parse proposal: %w", err)
			}

			amount, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			content, ok := v1beta1.ContentFromProposalType(proposal.Title, proposal.Description, proposal.Type)
			if !ok {
				return fmt.Errorf("failed to create proposal content: unknown proposal type %s", proposal.Type)
			}

			proposer, err := clientCtx.AddressCodec.BytesToString(clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			msg, err := v1beta1.NewMsgSubmitProposal(content, amount, proposer)
			if err != nil {
				return fmt.Errorf("invalid message: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagTitle, "", "The proposal title")
	cmd.Flags().String(FlagDescription, "", "The proposal description")
	cmd.Flags().String(FlagProposalType, "", "The proposal Type")
	cmd.Flags().String(FlagDeposit, "", "The proposal deposit")
	cmd.Flags().String(FlagProposal, "", "Proposal file path (if this path is given, other proposal flags are ignored)")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCmdWeightedVote implements creating a new weighted vote command.
// TODO(@julienrbrt): remove this once AutoCLI can flatten nested structs.
func NewCmdWeightedVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "weighted-vote <proposal-id> <weighted-options>",
		Aliases: []string{"vote-weighted"},
		Args:    cobra.ExactArgs(2),
		Short:   "Vote for an active proposal, options: yes/no/no-with-veto/abstain",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a vote for an active proposal. You can
find the proposal-id by running "%s query gov proposals".

Example:
$ %s tx gov weighted-vote 1 yes=0.6,no=0.3,abstain=0.05,no-with-veto=0.05 --from mykey
`,
				version.AppName, version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Get voter address
			from, err := clientCtx.AddressCodec.BytesToString(clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			// validate that the proposal id is a uint
			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("proposal-id %s not a valid int, please input a valid proposal-id", args[0])
			}

			// Figure out which vote options user chose
			options, err := v1.WeightedVoteOptionsFromString(govutils.NormalizeWeightedVoteOptions(args[1]))
			if err != nil {
				return err
			}

			metadata, err := cmd.Flags().GetString(FlagMetadata)
			if err != nil {
				return err
			}

			// Build vote message and run basic validation
			msg := v1.NewMsgVoteWeighted(from, proposalID, options, metadata)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagMetadata, "", "Specify metadata of the weighted vote")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
