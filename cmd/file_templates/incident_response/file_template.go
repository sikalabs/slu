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

var IndexFileTemplate = `# Incident Index

| date | level | title | author |
| --- | --- | --- | --- |
{{- range $incident := .Incidents }}
| {{ $incident.Date }} | {{ $incident.Level }} | {{ $incident.Title }} | {{ $incident.Author }} |
{{- end}}
`
