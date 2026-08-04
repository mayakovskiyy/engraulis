# About
Engraulis is a *lightweight* monitoring utility for websites and OS. Project is distributing under the `BSD 3 Clause` license.

# Where should I start?
First of all you should download the source code of the project (obviously). <br>
After you should set up your workspace: <br>
1. Install Go (1.25+).
2. Install `sqlite3` (or you should have a C compiler like `gcc`, `clang` or `zig cc` installed).<br>
3. Use the `go mod tidy` to set up the `sqlite3` in the [db.go](engraulis/sysmon/db.go) <br>

That's all! You're able to use any editor or IDE (Zed, Sublime Text, Vim, Nano, etc.).

# Text formatting
For files you must use the snake case. Example: `monitoring_darwin.go`. <br>
For functions you must use the camel case (for local functions) and the pascal case (for external functions). Example: `initDb` (This func does *not* exist), `MtrCurrent`

# Code
You should keep your code: <br>
- Understandable <br>
- Clean (not the 300 lines of shrink code)<br>

And also we're going by the KISS-way. So keep your code as simple as you could.

# LLMs Usage

This project **allows** the LLM usage. BUT: <br>
- We're going by the KISS-way. 300 lines of shrink LLM code would be rejected. <br>
- YOU are responsible for your code. Not LLM. Please, check and improve your code before you send a PR. <br>
You don't like the LLM Usage rules? Ok, you're able to fork the project and vibecode as much as you want.

# Footer
Thanks for reading this file. Good luck with the development <br>
Danketsu Studio, 2026