// Package inspect decodes a GEUL document (.geul) into interpreted entries for
// tools that show it: each HEAD and BODY record keeps its words and gains its
// spec name, its decoded content, its Edge Id and the Edge Ids it references.
// Content is decoded only where re-encoding gives back the original words, so a
// shown form never hides a byte; anything else stays a bare record.
package inspect
