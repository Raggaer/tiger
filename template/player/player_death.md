{{- if not .deaths }}
No deaths
{{- end }}
{{ range $i, $element := .deaths }}
{{- $index := sum $i 1 -}}
{{- $deathTime := unixToTime $element.Time -}}
{{- $index }}. Killed by **{{ $element.KilledBy }}** at level **{{ $element.Level }}** - *{{ timeAgoCurrent $deathTime }} ago*
{{ end }}