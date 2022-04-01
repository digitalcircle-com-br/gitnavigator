# gitnavigator
Simple tool to better manage / navigate in multiple repos

## Config

Add ~/.gitnavigator like:

```yaml
addr: :19999 #Port to listen
root: ~/projects #Git repos root dir
favorites: [] # Not used yet
global_cmds: # Global Cms
  # - label: pwd
  #   cmd: [pwd]
  # - label: ls
  #   cmd: [ls]
repo_cmds: # Per repo cms
  - label: Open Terminal
    cmd: [open, -a, Iterm.app, $repo]
  - label: Open Folder
    cmd: [open, $repo]
  - label: Open VSCode
    cmd: [code, $repo]
  - label: Open Goland
    cmd: [open, -a, Goland.app, $repo]

```