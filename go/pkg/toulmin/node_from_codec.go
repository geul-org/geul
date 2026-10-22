//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what NodeFromCodec — the node type (and, for rules, the role) a record codec declares
package toulmin

// nodeTypes lists every node type NodeFromCodec searches, in codebook order.
var nodeTypes = []NodeType{NodeNamespace, NodeCase, NodeRule, NodeFunc, NodeEndpoint, NodeTrigger, NodeTerm, NodeField, NodeAlias, NodeGroup}

// NodeFromCodec reports the node type ref declares. For the three rule codecs it
// also reports the role the codec fixes (ruleRole true); other node codecs leave
// the role to the header word. ok is false when ref is not a node codec.
func NodeFromCodec(ref CodecRef) (t NodeType, role RoleCode, ruleRole bool, ok bool) {
	for _, r := range []RoleCode{RoleGeneral, RoleCounter, RoleExcept} {
		if c, _ := ruleCodec(r); c == ref {
			return NodeRule, r, true, true
		}
	}
	for _, nt := range nodeTypes {
		if c, err := NodeCodec(nt, RoleGeneral); err == nil && c == ref && nt != NodeRule {
			return nt, 0, false, true
		}
	}
	return 0, 0, false, false
}
