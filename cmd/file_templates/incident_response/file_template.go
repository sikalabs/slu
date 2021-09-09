package incident_response

var IncidentResponseFileTemplate = `---
title: {{.Title}}
level: {{.Level}}
date: {{.Date}}
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
