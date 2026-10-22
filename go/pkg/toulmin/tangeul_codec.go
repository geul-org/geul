//ff:type feature=toulmin type=model
//ff:what TOULMIN codec table — the direct Edge codes and GEUL ISSUER format numbers of a .geul TANGEUL document
package toulmin

// IssuerGEUL is the 1-word GEUL Issuer (0x0000) under which the TANGEUL
// ISSUER-format records are defined (GEUL t002 TOULMIN 코덱).
const IssuerGEUL uint16 = 0x0000

// Direct TOULMIN Edge codecs (GEUL t002 OVERVIEW Edge `11`). FROZEN.
const (
	CodecCaseDeclare      uint8 = 0x60
	CodecPredicateDeclare uint8 = 0x61
	CodecSpecDeclare      uint8 = 0x62
	CodecRuleDeclare      uint8 = 0x63
	CodecCounterDeclare   uint8 = 0x64
	CodecExceptDeclare    uint8 = 0x65
	CodecAttacks          uint8 = 0x68
	CodecSupersedes       uint8 = 0x69
	CodecRun              uint8 = 0x70
	CodecDo               uint8 = 0x71
	CodecDoOnce           uint8 = 0x72
	CodecUndo             uint8 = 0x73
	CodecChecking         uint8 = 0x74
	CodecEntry            uint8 = 0x75
	CodecTrigger          uint8 = 0x76
)

// TOULMIN-family format numbers under the GEUL Issuer (ISSUER-1WORD 0x7C).
// FROZEN append-only.
const (
	FormatSource       uint16 = 0xC000
	FormatText         uint16 = 0xC001
	FormatSection      uint16 = 0xC002
	FormatCaseBegin    uint16 = 0xC003
	FormatCaseEnd      uint16 = 0xC004
	FormatProse        uint16 = 0xC005
	FormatNamespace    uint16 = 0xC010
	FormatFunc         uint16 = 0xC011
	FormatField        uint16 = 0xC012
	FormatAlias        uint16 = 0xC013
	FormatGroup        uint16 = 0xC014
	FormatRequires     uint16 = 0xC020
	FormatMeans        uint16 = 0xC021
	FormatHas          uint16 = 0xC022
	FormatRuns         uint16 = 0xC023
	FormatEvaluates    uint16 = 0xC024
	FormatTriggers     uint16 = 0xC025
	FormatUntil        uint16 = 0xC026
	FormatSameAs       uint16 = 0xC027
	FormatImportedFrom uint16 = 0xC028
	FormatLabel        uint16 = 0xC029
)
