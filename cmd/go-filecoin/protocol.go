package commands

import (
	cmdkit "github.com/ipfs/go-ipfs-cmdkit"
	cmds "github.com/ipfs/go-ipfs-cmds"

	"github.com/filecoin-project/go-filecoin/internal/app/go-filecoin/porcelain"
)

var protocolCmd = &cmds.Command{
	Helptext: cmdkit.HelpText{
		Tagline: "Show protocol parameter details",
	},
	Run: func(req *cmds.Request, re cmds.ResponseEmitter, env cmds.Environment) error {
		params, err := GetPorcelainAPI(env).ProtocolParameters(env.Context())
		if err != nil {
			return err
		}
		return re.Emit(params)
	},
	Type: porcelain.ProtocolParams{},
}
