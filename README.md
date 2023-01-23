### redfile, task runner / build tool.
Clone of Taskfile, used for playing with Cue.

## Example:
```
version: 3.0

env:
  Name: John
  Age: 57
  Class: Warrior
  Region: US

tasks:
  taskone:
    description: a few prints
    cmds:
      - date
      - echo 2
      - echo $GREETING
      - echo $Region
    env:
      GREETING: Hey, there!
  tasktwo:
    cmds:
      - echo 3
      - echo 4
      - echo 5
```

# Commands
To install taskfile CLI, run command in root folder:<br>
`go install .`
```
redfile -h

Usage:
  redfile [command]

Available Commands:
  cmds        A brief description of your command
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  run         A brief description of your command
  validate    Validate redfile config file against the scheme
```

```
taskfile cmds:

Tasks:
taskone: a few prints
tasktwo: <nil>
```

```
taskfile validate:

env.Age: conflicting values 57 and string (mismatched types int and string)
There were no validation error in your redfile!
```

```
redfile run -c taskone

Executing taskone:
Sun Jan 22 10:28:04 IST 2023
2
```

### TODOs:
[x] env not currently working, but supported