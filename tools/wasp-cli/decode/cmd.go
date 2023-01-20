package decode

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"

	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/tools/wasp-cli/log"
	"github.com/iotaledger/wasp/tools/wasp-cli/util"
)

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(initDecodeCmd())
	rootCmd.AddCommand(initDecodeMetadataCmd())
}

func initDecodeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "decode <type> <key> <type> ...",
		Short: "Decode the output of a contract function call",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			d := util.UnmarshalDict()

			if len(args) == 2 {
				ktype := args[0]
				vtype := args[1]

				for key, value := range d {
					skey := util.ValueToString(ktype, []byte(key))
					sval := util.ValueToString(vtype, value)
					log.Printf("%s: %s\n", skey, sval)
				}
				return
			}

			if len(args) < 3 || len(args)%3 != 0 {
				log.Check(cmd.Help())
				return
			}

			for i := 0; i < len(args)/2; i++ {
				ktype := args[i*2]
				skey := args[i*2+1]
				vtype := args[i*2+2]

				key := kv.Key(util.ValueFromString(ktype, skey))
				val := d.MustGet(key)
				if val == nil {
					log.Printf("%s: <nil>\n", skey)
				} else {
					log.Printf("%s: %s\n", skey, util.ValueToString(vtype, val))
				}
			}
		},
	}
}

func initDecodeMetadataCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "decode-metadata <0x...>",
		Short: "Translates metadata from Hex to a humanly-readable format",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			metadata, err := isc.RequestMetadataFromBytes(hexutil.MustDecode(args[0]))
			log.Check(err)
			jsonBytes, err := json.MarshalIndent(metadata, "", "  ")
			log.Check(err)
			log.Printf("%s\n", jsonBytes)
		},
	}
}
