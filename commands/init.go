package commands

import (
	"fmt"
	"io"
	"os"

	cmds "gx/ipfs/QmVTmXZC2yE38SDKRihn96LXX6KwBWgzAg8aCDZaMirCHm/go-ipfs-cmds"
	cmdkit "gx/ipfs/QmdE4gMduCKCGAcczM2F5ioYDfdeKuPix138wrES1YSr7f/go-ipfs-cmdkit"

	"github.com/filecoin-project/go-filecoin/api"
)

var initCmd = &cmds.Command{
	Helptext: cmdkit.HelpText{
		Tagline: "Initialize a filecoin repo",
	},
	Options: []cmdkit.Option{
		cmdkit.StringOption("walletfile", "wallet data file: contains addresses and private keys").WithDefault(""),
		cmdkit.StringOption("walletaddr", "address to store in nodes backend when '--walletfile' option is passed").WithDefault(""),
		cmdkit.StringOption("genesisfile", "path of file containing archive of genesis block DAG data"),
		cmdkit.BoolOption("testgenesis", "when set, creates a custom genesis block with pre-mined funds"),
	},
	Run: func(req *cmds.Request, re cmds.ResponseEmitter, env cmds.Environment) {
		repoDir := getRepoDir(req)
		re.Emit(fmt.Sprintf("initializing filecoin node at %s\n", repoDir)) // nolint: errcheck

		walletFile, _ := req.Options["walletfile"].(string)
		walletAddr, _ := req.Options["walletaddr"].(string)
		genesisFile, _ := req.Options["genesisfile"].(string)
		customGenesis, _ := req.Options["testgenesis"].(bool)

		err := GetAPI(env).Daemon().Init(
			req.Context,
			api.RepoDir(repoDir),
			api.WalletFile(walletFile),
			api.WalletAddr(walletAddr),
			api.GenesisFile(genesisFile),
			api.UseCustomGenesis(customGenesis),
		)

		if err != nil {
			re.SetError(err, cmdkit.ErrNormal)
			return
		}
	},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeEncoder(initTextEncoder),
	},
}

func initTextEncoder(req *cmds.Request, w io.Writer, val interface{}) error {
	_, err := fmt.Fprintf(w, val.(string))
	return err
}

func getRepoDir(req *cmds.Request) string {
	envdir := os.Getenv("FIL_PATH")

	repodir, ok := req.Options[OptionRepoDir].(string)
	if ok {
		return repodir
	}

	if envdir != "" {
		return envdir
	}

	return "~/.filecoin"
}
