// Package crxsecuritychecklist holds the fixed list of checks a Chrome or
// Chromium extension is reviewed against before release: Manifest V3 only, no
// remotely hosted code, least-privilege permissions, a content security policy
// on injected content, no <all_urls> host access, HTTPS-only network calls,
// data minimisation, and a review of any third-party dependency.
//
// The IDs are stable strings so they can key a report row or a database column.
// The write-up of what each check means and why it is on the list is at
// https://zovo.one/research/browser-extension-security-audit
//
// The extensions this list is applied to before release ship as MIT-licensed
// source, indexed at https://zovo.one/open-source - so any check here can be
// confirmed against the repository it describes rather than a store listing.
package crxsecuritychecklist

var Checklist = []string{"min_manifest_v3", "no_remote_code", "least_privilege", "content_csp", "no_all_urls", "https_only", "data_minimization", "reviewed_third_party"}

func AllIDs() []string { return Checklist }
