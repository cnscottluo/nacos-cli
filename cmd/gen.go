package cmd

import (
	"errors"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/spf13/cobra"
)

const (
	defaultKeyLength    = 16
	defaultValueLength  = 32
	defaultSymbolLength = 1
)

var (
	keyLength    uint8
	valueLength  uint8
	symbolLength uint8
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate secret key, identity key and value",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var genSecretKeyCmd = &cobra.Command{
	Use:   "secret",
	Short: "generate secret key",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := internal.GenerateAESKey()
		internal.CheckErr(err)
		encodeKey := internal.Base64Encode(key)
		internal.Info("Secret key: %s", encodeKey)
	},
}

var genIdentityCmd = &cobra.Command{
	Use:   "identity",
	Short: "generate identity key and value",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if keyLength < defaultKeyLength {
			return errors.New("identity key length must be at least 16")
		}
		if valueLength < defaultValueLength {
			return errors.New("identity value length must be at least 32")
		}
		if symbolLength < defaultSymbolLength {
			return errors.New("identity symbol length must be at least 1")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		key, value, err := internal.GenerateIdentity(keyLength, valueLength, symbolLength)
		internal.CheckErr(err)
		internal.Info("Identity key: %s\nIdentity value: %s", key, value)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.AddCommand(genSecretKeyCmd)
	genCmd.AddCommand(genIdentityCmd)
	genIdentityCmd.Flags().Uint8VarP(&keyLength, "key-len", "k", defaultKeyLength, "Length of identity key.")
	genIdentityCmd.Flags().Uint8VarP(&valueLength, "value-len", "v", defaultValueLength, "Length of identity value.")
	genIdentityCmd.Flags().Uint8VarP(
		&symbolLength, "symbol-len", "s", defaultSymbolLength, "Length of identity symbol.",
	)
}
