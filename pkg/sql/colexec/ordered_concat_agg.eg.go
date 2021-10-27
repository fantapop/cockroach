// Code generated by execgen; DO NOT EDIT.
// Copyright 2020 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"unsafe"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
)

func newConcatOrderedAggAlloc(allocator *colmem.Allocator, allocSize int64) aggregateFuncAlloc {
	return &concatOrderedAggAlloc{aggAllocBase: aggAllocBase{
		allocator: allocator,
		allocSize: allocSize,
	}}
}

type concatOrderedAgg struct {
	orderedAggregateFuncBase
	allocator *colmem.Allocator
	// curAgg holds the running total.
	curAgg []byte
	// col points to the output vector we are updating.
	col *coldata.Bytes
	// vec is the same as col before conversion from coldata.Vec.
	vec coldata.Vec
	// foundNonNullForCurrentGroup tracks if we have seen any non-null values
	// for the group that is currently being aggregated.
	foundNonNullForCurrentGroup bool
}

func (a *concatOrderedAgg) Init(groups []bool, vec coldata.Vec) {
	a.orderedAggregateFuncBase.Init(groups, vec)
	a.vec = vec
	a.col = vec.Bytes()
	a.Reset()
}

func (a *concatOrderedAgg) Reset() {
	a.orderedAggregateFuncBase.Reset()
	a.foundNonNullForCurrentGroup = false
	a.curAgg = zeroBytesValue
}

func (a *concatOrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	oldCurAggSize := len(a.curAgg)
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Bytes(), vec.Nulls()
	a.allocator.PerformOperation(
		[]coldata.Vec{a.vec},
		func() {
			groups := a.groups
			if sel == nil {
				_ = groups[inputLen-1]
				if nulls.MaybeHasNulls() {
					for i := 0; i < inputLen; i++ {

						if groups[i] {
							// If we encounter a new group, and we haven't found any non-nulls for the
							// current group, the output for this group should be null.
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col.Set(a.curIdx, a.curAgg)
							}
							a.curIdx++
							a.curAgg = zeroBytesValue

							a.foundNonNullForCurrentGroup = false
						}

						var isNull bool
						isNull = nulls.NullAt(i)
						if !isNull {
							a.curAgg = append(a.curAgg, col.Get(i)...)
							a.foundNonNullForCurrentGroup = true
						}
					}
				} else {
					for i := 0; i < inputLen; i++ {

						if groups[i] {
							// If we encounter a new group, and we haven't found any non-nulls for the
							// current group, the output for this group should be null.
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col.Set(a.curIdx, a.curAgg)
							}
							a.curIdx++
							a.curAgg = zeroBytesValue

						}

						var isNull bool
						isNull = false
						if !isNull {
							a.curAgg = append(a.curAgg, col.Get(i)...)
							a.foundNonNullForCurrentGroup = true
						}
					}
				}
			} else {
				sel = sel[:inputLen]
				if nulls.MaybeHasNulls() {
					for _, i := range sel {

						if groups[i] {
							// If we encounter a new group, and we haven't found any non-nulls for the
							// current group, the output for this group should be null.
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col.Set(a.curIdx, a.curAgg)
							}
							a.curIdx++
							a.curAgg = zeroBytesValue

							a.foundNonNullForCurrentGroup = false
						}

						var isNull bool
						isNull = nulls.NullAt(i)
						if !isNull {
							a.curAgg = append(a.curAgg, col.Get(i)...)
							a.foundNonNullForCurrentGroup = true
						}
					}
				} else {
					for _, i := range sel {

						if groups[i] {
							// If we encounter a new group, and we haven't found any non-nulls for the
							// current group, the output for this group should be null.
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col.Set(a.curIdx, a.curAgg)
							}
							a.curIdx++
							a.curAgg = zeroBytesValue

						}

						var isNull bool
						isNull = false
						if !isNull {
							a.curAgg = append(a.curAgg, col.Get(i)...)
							a.foundNonNullForCurrentGroup = true
						}
					}
				}
			}
		},
	)
	newCurAggSize := len(a.curAgg)
	a.allocator.AdjustMemoryUsage(int64(newCurAggSize - oldCurAggSize))
}

func (a *concatOrderedAgg) Flush(outputIdx int) {
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	if !a.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.col.Set(outputIdx, a.curAgg)
	}
	// Release the reference to curAgg eagerly.
	a.allocator.AdjustMemoryUsage(-int64(len(a.curAgg)))
	a.curAgg = nil
}

type concatOrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []concatOrderedAgg
}

var _ aggregateFuncAlloc = &concatOrderedAggAlloc{}

const sizeOfConcatOrderedAgg = int64(unsafe.Sizeof(concatOrderedAgg{}))
const concatOrderedAggSliceOverhead = int64(unsafe.Sizeof([]concatOrderedAgg{}))

func (a *concatOrderedAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(concatOrderedAggSliceOverhead + sizeOfConcatOrderedAgg*a.allocSize)
		a.aggFuncs = make([]concatOrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	f.allocator = a.allocator
	a.aggFuncs = a.aggFuncs[1:]
	return f
}
