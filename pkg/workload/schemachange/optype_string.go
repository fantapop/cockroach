// Code generated by "stringer -type=opType"; DO NOT EDIT.

package schemachange

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[addColumn-0]
	_ = x[addConstraint-1]
	_ = x[createIndex-2]
	_ = x[createSequence-3]
	_ = x[createTable-4]
	_ = x[createTableAs-5]
	_ = x[createView-6]
	_ = x[createEnum-7]
	_ = x[createSchema-8]
	_ = x[dropColumn-9]
	_ = x[dropColumnDefault-10]
	_ = x[dropColumnNotNull-11]
	_ = x[dropColumnStored-12]
	_ = x[dropConstraint-13]
	_ = x[dropIndex-14]
	_ = x[dropSequence-15]
	_ = x[dropTable-16]
	_ = x[dropView-17]
	_ = x[dropSchema-18]
	_ = x[renameColumn-19]
	_ = x[renameIndex-20]
	_ = x[renameSequence-21]
	_ = x[renameTable-22]
	_ = x[renameView-23]
	_ = x[setColumnDefault-24]
	_ = x[setColumnNotNull-25]
	_ = x[setColumnType-26]
	_ = x[insertRow-27]
	_ = x[validate-28]
}

const _opType_name = "addColumnaddConstraintcreateIndexcreateSequencecreateTablecreateTableAscreateViewcreateEnumcreateSchemadropColumndropColumnDefaultdropColumnNotNulldropColumnStoreddropConstraintdropIndexdropSequencedropTabledropViewdropSchemarenameColumnrenameIndexrenameSequencerenameTablerenameViewsetColumnDefaultsetColumnNotNullsetColumnTypeinsertRowvalidate"

var _opType_index = [...]uint16{0, 9, 22, 33, 47, 58, 71, 81, 91, 103, 113, 130, 147, 163, 177, 186, 198, 207, 215, 225, 237, 248, 262, 273, 283, 299, 315, 328, 337, 345}

func (i opType) String() string {
	if i < 0 || i >= opType(len(_opType_index)-1) {
		return "opType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _opType_name[_opType_index[i]:_opType_index[i+1]]
}
