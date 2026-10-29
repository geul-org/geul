package toulmin

// putValue is PutValue as TEXT and SOURCE call it; tests replace it to reach
// the error path (a Value over the 2^32-1-word record limit).
var putValue = PutValue
