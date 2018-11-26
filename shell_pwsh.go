package main

import (
	"encoding/json"
)

// pwshShell is not a real shell
type pwsh struct{}

// PowerShell instance
var PowerShell Shell = pwsh{}

const powershellHook = `
# Run before the prompt
function Direnv-Hook {
  $diff = & "{{.SelfPath}}" export json | ConvertFrom-Json

  # No output, return
  if (!$diff) {
    return
  }

  # Apply the diff
  $diff.psobject.properties | ForEach {
    $o = $_
    if ($o.Value -Or $o.Value -eq "") {
      #Write-Output "Setting $($o.Name) to $(o.Value)"
      Set-Item -path "Env:$($o.Name)" -value $o.Value
    } else {
      #Write-Output "Deleting $($o.Name)"
      Remove-Item -path "Env:$($o.Name)"
    }
  }
}

$_direnv_prompt_old = Get-Command prompt
function prompt {
	Direnv-Hook
	return & $_direnv_prompt_old
}
`

func (sh pwsh) Hook() (string, error) {
	return powershellHook, nil
}

func (sh pwsh) Export(e ShellExport) string {
	out, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		// Should never happen
		panic(err)
	}
	return string(out)
}

func (sh pwsh) Dump(env Env) string {
	out, err := json.MarshalIndent(env, "", "  ")
	if err != nil {
		// Should never happen
		panic(err)
	}
	return string(out)
}
