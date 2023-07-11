# Runner

<img alt="interactive TUI" src="https://github.com/amonks/runner/blob/main/screenshots/tui.gif?raw=true" />

```toml

```

Runner runs a collection of programs specified in tasks.toml files, and
provides a UI for inspecting their execution. Runner's interactive UI for
long-lived programs has full mouse support.

<img alt="noninteractive printed output" src="https://github.com/amonks/runner/blob/main/screenshots/printer.gif?raw=true" />

Runner also works well for short-lived processes, and its interleaved output
can be sent to a file.

```go
package main

import "github.com/amonks/runner"

func main() {
	tasks, err := runner.Load(".")
	if err != nil {
		log.Fatal(err)
	}

	run := runner.RunTask(tasks, "dev")

	ui := runner.NewTUI()
	ui.Start(os.Stdin, os.Stdout, run.IDs())
	run.Start(ui)
}
```

Runner can be used and extended programatically through its Go API. See [the
godoc][godoc]

[godoc]: https://amonks.github.io/runner

## Installation

Runner is a single binary, which you can download from from the [releases
page][releases].

Alternately, if you already use go, you can install Runner with the go command
line tool:

    $ go install github.com/amonks/runner@latest

[releases]: https://github.com/amonks/runner/releases

## Task Files

Task files are called "tasks.toml". They specify one or more tasks.

```toml
[[task]]
  id = "dev"
  type = "long"
  dependencies = ["simulate-coding"]
  triggers = ["build-css", "build-js"]
  watch = ["server-config.json"]
  cmd = """
    echo "dev-server running at http://localhost:3000"
    while true; do sleep 1; done
  """
```

Let's go through the fields that can be specified on tasks.

### ID

ID identifies a task, for example,

- for command line invocation, as in `$ runner <id>`
- in the TUI's task list.

### Type

Type specifies how we manage a task.

If the Type is "long",

- We will restart the task if it exits.
- If the long task A is a dependency or trigger of task B, we will begin B as
  soon as A starts.

If the Type is "short",

- If the task exits 0, we will consider it done.
- If the task exits !0, we will wait 1 second and rerun it.
- If the short task A is a dependency or trigger of task B, we will wait for A
  to complete before starting B.

Any Type besides "long" or "short" is invalid. There is no default type: every
task must specify its type.

### Dependencies

Dependencies are other tasks IDs which should always run alongside this task.
If a task A lists B as a dependency, running A will first run B.

Dependencies do not set up an invalidation relationship: if long task A lists
short task B as a dependency, and B reruns because a watched file is changed,
we will not restart A, assuming that A has its own mechanism for detecting file
changes. If A does not have such a mechanhism, use a trigger rather than a
dependency.

Dependencies can be task IDs from child directories. For example, the
dependency "css/build" specifies the task with ID "build" in the tasks file
"./css/tasks.toml".

### Triggers

Triggers are other task IDs which should always be run alongside this task, and
whose success should cause this task to re-execute. If a task A lists B as a
dependency, and both A and B are running, successful execution of B will always
trigger an execution of A.

Triggers can be task IDs from child directories. For example, the trigger
"css/build" specifies the task with ID "build" in the tasks file
"./css/tasks.toml".

### Watch

Watch specifies file paths where, if a change to the file path is detected, we
should restart the task. Recursive paths are specified with the suffix "/...".

For example,

- `"."` watches for changes to the working directory only, but not changes
  within subdirectories.
- `"./..."` watches for changes at any level within the working directory.
- `"./some/path/file.txt"` watches for changes to the file, which may or may
  not already exist.

### CMD

CMD is the command to run. It runs in a new bash process, as in,

    $ bash -c "$CMD"

CMD can have many lines.

## CLI Usage

```
$ runner dev
```

Runner takes one argument: the task ID to run. Runner looks for a task file in
the current directory.

### User Interfaces

Runner has two UIs that it deploys in different circumstances, a TUI and a
Printer. You can force Runner to use a particular UI by passing the 'ui' flag,
as in,

    $ runner -ui=printer dev

#### Interactive TUI

<img alt="interactive TUI" src="https://github.com/amonks/runner/blob/main/screenshots/tui.gif?raw=true" />

The Interactive TUI is used whenever both,

1. stdout is a tty (eg Runner is _not_ being piped to a file), and,
2. any running task is "long" (eg an ongoing "dev server" process rather than a
   one-shot "build" procedure).

For example, when running a dev server or test executor that stays running
while you make changes.

#### Non-Interactive Printer UI

|------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------|
| <img alt="non-interactive output" src="https://github.com/amonks/runner/blob/main/screenshots/printer.gif?raw=true" /> | <img alt="redirected output" src="https://github.com/amonks/runner/blob/main/screenshots/nontty.gif?raw=true" /> |
|------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------|

Runner simply prints its output if either,

1. runner is not a tty (eg Runner is being piped to a file), or,
2. no tasks are "long" (eg a one-shot "build" procedure, rather than an ongoing
   "dev server").

## Programmatic Use

Runner can be used and extended programatically through its Go API. For more
information, including a conceptual overview of the architecture, example code,
and reference documentation, see [the godoc][godoc].

[godoc]: https://amonks.github.io/runner