// Code generated by ent, DO NOT EDIT.

package queue

import (
	"queuetable/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Queue {
	return predicate.Queue(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Queue {
	return predicate.Queue(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Queue {
	return predicate.Queue(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Queue {
	return predicate.Queue(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Queue {
	return predicate.Queue(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Queue {
	return predicate.Queue(sql.FieldLTE(FieldID, id))
}

// Instruction applies equality check predicate on the "instruction" field. It's identical to InstructionEQ.
func Instruction(v string) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldInstruction, v))
}

// ObjID applies equality check predicate on the "obj_id" field. It's identical to ObjIDEQ.
func ObjID(v uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldObjID, v))
}

// InstructionEQ applies the EQ predicate on the "instruction" field.
func InstructionEQ(v string) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldInstruction, v))
}

// InstructionNEQ applies the NEQ predicate on the "instruction" field.
func InstructionNEQ(v string) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldInstruction, v))
}

// InstructionIn applies the In predicate on the "instruction" field.
func InstructionIn(vs ...string) predicate.Queue {
	return predicate.Queue(sql.FieldIn(FieldInstruction, vs...))
}

// InstructionNotIn applies the NotIn predicate on the "instruction" field.
func InstructionNotIn(vs ...string) predicate.Queue {
	return predicate.Queue(sql.FieldNotIn(FieldInstruction, vs...))
}

// InstructionGT applies the GT predicate on the "instruction" field.
func InstructionGT(v string) predicate.Queue {
	return predicate.Queue(sql.FieldGT(FieldInstruction, v))
}

// InstructionGTE applies the GTE predicate on the "instruction" field.
func InstructionGTE(v string) predicate.Queue {
	return predicate.Queue(sql.FieldGTE(FieldInstruction, v))
}

// InstructionLT applies the LT predicate on the "instruction" field.
func InstructionLT(v string) predicate.Queue {
	return predicate.Queue(sql.FieldLT(FieldInstruction, v))
}

// InstructionLTE applies the LTE predicate on the "instruction" field.
func InstructionLTE(v string) predicate.Queue {
	return predicate.Queue(sql.FieldLTE(FieldInstruction, v))
}

// InstructionContains applies the Contains predicate on the "instruction" field.
func InstructionContains(v string) predicate.Queue {
	return predicate.Queue(sql.FieldContains(FieldInstruction, v))
}

// InstructionHasPrefix applies the HasPrefix predicate on the "instruction" field.
func InstructionHasPrefix(v string) predicate.Queue {
	return predicate.Queue(sql.FieldHasPrefix(FieldInstruction, v))
}

// InstructionHasSuffix applies the HasSuffix predicate on the "instruction" field.
func InstructionHasSuffix(v string) predicate.Queue {
	return predicate.Queue(sql.FieldHasSuffix(FieldInstruction, v))
}

// InstructionEqualFold applies the EqualFold predicate on the "instruction" field.
func InstructionEqualFold(v string) predicate.Queue {
	return predicate.Queue(sql.FieldEqualFold(FieldInstruction, v))
}

// InstructionContainsFold applies the ContainsFold predicate on the "instruction" field.
func InstructionContainsFold(v string) predicate.Queue {
	return predicate.Queue(sql.FieldContainsFold(FieldInstruction, v))
}

// ObjTableEQ applies the EQ predicate on the "obj_table" field.
func ObjTableEQ(v ObjTable) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldObjTable, v))
}

// ObjTableNEQ applies the NEQ predicate on the "obj_table" field.
func ObjTableNEQ(v ObjTable) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldObjTable, v))
}

// ObjTableIn applies the In predicate on the "obj_table" field.
func ObjTableIn(vs ...ObjTable) predicate.Queue {
	return predicate.Queue(sql.FieldIn(FieldObjTable, vs...))
}

// ObjTableNotIn applies the NotIn predicate on the "obj_table" field.
func ObjTableNotIn(vs ...ObjTable) predicate.Queue {
	return predicate.Queue(sql.FieldNotIn(FieldObjTable, vs...))
}

// ObjIDEQ applies the EQ predicate on the "obj_id" field.
func ObjIDEQ(v uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldEQ(FieldObjID, v))
}

// ObjIDNEQ applies the NEQ predicate on the "obj_id" field.
func ObjIDNEQ(v uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldNEQ(FieldObjID, v))
}

// ObjIDIn applies the In predicate on the "obj_id" field.
func ObjIDIn(vs ...uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldIn(FieldObjID, vs...))
}

// ObjIDNotIn applies the NotIn predicate on the "obj_id" field.
func ObjIDNotIn(vs ...uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldNotIn(FieldObjID, vs...))
}

// ObjIDGT applies the GT predicate on the "obj_id" field.
func ObjIDGT(v uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldGT(FieldObjID, v))
}

// ObjIDGTE applies the GTE predicate on the "obj_id" field.
func ObjIDGTE(v uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldGTE(FieldObjID, v))
}

// ObjIDLT applies the LT predicate on the "obj_id" field.
func ObjIDLT(v uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldLT(FieldObjID, v))
}

// ObjIDLTE applies the LTE predicate on the "obj_id" field.
func ObjIDLTE(v uuid.UUID) predicate.Queue {
	return predicate.Queue(sql.FieldLTE(FieldObjID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Queue) predicate.Queue {
	return predicate.Queue(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Queue) predicate.Queue {
	return predicate.Queue(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Queue) predicate.Queue {
	return predicate.Queue(sql.NotPredicates(p))
}