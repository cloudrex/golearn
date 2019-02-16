package util

import (
	"github.com/llir/llvm/ir"
)

// FindAllocaInBlock : Attempt to find a single instruction in a block by it's name. Returns 'nil' if nothing was found.
func FindAllocaInBlock(block *ir.Block, name string) *ir.InstAlloca {
	var target *ir.InstAlloca

	for i := 0; i < len(block.Insts); i++ {
		inst := block.Insts[i].(*ir.InstAlloca)

		if inst.LocalIdent.Name() == name {
			target = inst

			break
		}
	}

	return target
}
