{{- if not .deaths }}
No deaths
{{- end }}
{{ range $i, $element := .deaths }}
{{- $index := sum $i 1 -}}
{{- $deathTime := unixToTime $element.Time -}}
{{ $index }}. Killed **{{ $element.Player.Name }}** at level **{{ $element.Level }}** - *{{ timeAgoCurrent $deathTime }} ago*
{{ end }}