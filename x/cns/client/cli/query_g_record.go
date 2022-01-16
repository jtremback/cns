package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jtremback/cns/x/cns/types"
	"github.com/spf13/cobra"
)

func CmdListGRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-g-record",
		Short: "list all GRecord",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGRecordRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GRecordAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowGRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-g-record [key]",
		Short: "shows a GRecord",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argKey := args[0]

			params := &types.QueryGetGRecordRequest{
				Key: argKey,
			}

			res, err := queryClient.GRecord(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
