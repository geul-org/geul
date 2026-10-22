// Package geulmd renders a GEUL document (.geul) as GEULMD — standard Markdown
// (media type text/vnd.geulmd, extension .md) whose first fenced block, info
// string `geul`, holds the HEAD and whose second holds the BODY, one record per
// line. A line is written in structured form only when re-encoding it gives back
// the original words; otherwise it falls back to raw words, so nothing is lost.
package geulmd
