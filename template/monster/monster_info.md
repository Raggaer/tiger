**You view {{ .monster.Description }}**

- **Experience**: {{ .monster.Experience }}
- **Speed**: {{ .monster.Speed }}
- **Health**: {{ .monster.Health.Now }}

{{ range $index, $element := .monster.Attacks.Attacks }} 
- **Attack**: {{ $element.Name }} ({{ $element.Min}}, {{ $element.Max}}) 
{{- end }}
