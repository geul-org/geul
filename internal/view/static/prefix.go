package static

// Prefix is the URL path the embedded scripts are served under. It sits
// inside the public tree path so a proxy that forwards only /.well-known/geul/
// carries them too; `_` never begins an Issuer or document directory name.
const Prefix = "/.well-known/geul/_static/"
