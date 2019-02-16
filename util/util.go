package util

import (
	"github.com/llir/llvm/ir"
)

// FindAllocaInBlock : Attempt to find a single instruction in a block by it's name. Returns 'nil' if nothing was found.
func FindAllocaInBlock(block *ir.Block, name string) *ir.InstAlloca {
	var target *ir.InstAlloca

	for _, v := range block.Insts {
		inst, ok := v.(*ir.InstAlloca)

		if !ok { // Conversion failed. Continue.
			continue
		} else if inst.LocalIdent.Name() == name {
			target = inst

			break
		}
	}

	return target
}
