//ff:type feature=core type=model
//ff:what Absent — the one-word record that fills an optional field left out (TAG WORD 0x0000)
package core

// Absent is the TAG WORD 0x0000 written as a whole one-word record in an
// optional field that is left out (GEUL t002 docs/edge/CODEC 선택 필드).
// Reference Codec 0000 is reserved and LLLL 0000 makes it one word, so it
// is no other record. It is valid only in a field whose CODEC list names the
// target 0000; as a HEAD or BODY record it is still a reserved Reference.
const Absent Word = 0x0000
