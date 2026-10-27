package server

// wellKnown is the URL path the tree mode serves, mirroring the public address
// space (GEUL-ORG 공표 주소).
const wellKnown = "/.well-known/geul/"

// MediaType is the media type of a .geul document (GEUL-ORG 미디어 타입).
const MediaType = "application/vnd.geul"

// indexJSON is the file name, inside a tree directory, of its listing as JSON
// for the tree browser.
const indexJSON = "index.json"
