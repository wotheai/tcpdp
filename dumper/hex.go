package dumper

import (
	"encoding/hex"
	"fmt"

	"go.uber.org/zap"
)

// HexDumper ...
type HexDumper struct {
	Logger *zap.Logger
}

// Dump TCP
func (h *HexDumper) Dump(in []byte) error {
	h.Logger.Info(fmt.Sprintf("\n%s", hex.Dump(in)))
	return nil
}
