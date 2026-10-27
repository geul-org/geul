// Package page turns an interpreted GEUL document (inspect.Document) into the
// viewer page: one row per HEAD and BODY record with its fields, Edge Id links,
// back-references, source span and words, rendered as one self-contained HTML
// page. It decodes nothing itself; every reading comes from geul/go.
package page
