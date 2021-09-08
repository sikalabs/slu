package incident_response

var IncidentResponseFileTemplate = `---
date: {{.Date}}
level: {{.Level}}
author: {{.Author}}
---

# {{.Date}}: {{.Title}}

## Problem

TODO

## Cause

TODO

## Solution

TODO

## Prevention

TODO
`
