![logo](./assets/logo.png)

## What is godo?

"godo" is a user-friendly CLI-based todo application written in the Go programming language. It allows seamless task management directly from the command line, emphasizing simplicity and efficiency.

## Showcase

```sh
# Add tasks
$ godo add "Learn Go" "Learn Rust" "Learn Zig"
    ☐ 1. Learn Go
    ☐ 2. Learn Rust
    ☐ 3. Learn Zig

# Mark first task as done
$ godo do 1
    ☑ 1. Learn Go
    ☐ 2. Learn Rust
    ☐ 3. Learn Zig

# Remove third task from the list
$ godo rm 3
    ☑ 1. Learn Go
    ☐ 2. Learn Rust

# Remove completed tasks
$ godo clear
    ☐ 1. Learn Rust

# Remove all tasks
$ godo wipe
    Task list is empty
```

## Installation

To build "godo" and use it as a standalone application, follow these steps:

1. Ensure you have Go installed on your system.
2. Run the following commands to build "godo":
```
git clone github.com/axsbee/godo
cd godo
go build
```

## Usage

Here are some useful commands:

- `godo help` - show what commands are available.
- `godo list` - Display all the tasks in your todo list.
- `godo add "Task Description"` - Add a new task to your todo list.
- `godo do <task_index>` - Mark a task as complete.
- `godo rm <task_index>` - Remove a task from your list.

## License

This project is licensed under the [MIT License](LICENSE.md).