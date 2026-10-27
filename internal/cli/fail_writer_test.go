//ff:type feature=cli type=model
//ff:what failWriter — test writer whose every write fails
package cli

// failWriter fails every write.
type failWriter struct{}
