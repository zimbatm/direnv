# Install with:
#
#   new-item -itemtype file -path $profile -force
#   cp direnv.ps1 $profile
#
#   # maybe
#   Set-ExecutionPolicy RemoteSigned
#
#

# Run before the prompt
function Direnv-Hook {
  # Use -AsHashtable so that mixed case keys are unique
  # Only available in powershell 6.0
  # TODO: verify this claim ^^^
  $diff = & direnv export json | ConvertFrom-Json -AsHashtable

  # No output, return
  if (!$diff) {
    $host.ui.WriteErrorLine('direnv: no diff')
    return
  }

  # Apply the diff
  $diff.psobject.properties | ForEach {
    $o = $_
    if ($o.Value -Or $o.Value -eq "") {
      $host.ui.WriteErrorLine("Setting $($o.Name) to $($o.Value)")
      Set-Item -path "Env:$($o.Name)" -value $o.Value
    } else {
      $host.ui.WriteErrorLine("Deleting $($o.Name)")
      Remove-Item -path "Env:$($o.Name)"
    }
  }
}

Rename-Item -Path Function:Global:prompt -NewName Function:Direnv-Prompt-Backup

function Global:prompt {
  Direnv-Hook
  return Direnv-Prompt-Backup
}
