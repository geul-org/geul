// Package schema checks .geul byte strings for format: the zero-layer rules any
// decoder can apply without a codebook (framing, document structure, Kind
// rules), then codec layouts given as Codecs, and later case frames and
// selectional restrictions.
package schema
