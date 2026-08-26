# [Engraulis](https://engraulis.mayakovskiyy.space)

A lightweight **Go** monitoring module. <br>

# How to use Engraulis?
`engraulis/...` stands for `github.com/mayakovskiyy/engraulis/...`.<br>

First of all you should choose between `engraulis/sysmon` and `engraulis/client` (OR you can use both of them).<br>

For `engraulis/client` follow these steps: <br>
1. Download this repo as a `.zip` archive <br>
2. Add a `client` folder to use it as a module <br>
3. Add a `{your .mod name}/client` to the `import` line <br>
4. Finished! <br>

For `engraulis/sysmon` follow these steps: <br>
!WARNING `engraulis/sysmon` **DOES NOT** WORK ON WINDOWS! It supports only **macOS** and **Linux** (Distributions) <br>
1. Download this repo as a `.zip` archive <br>
2. Download `sqlite3` <br> 
3. Add a `sysmon` folder to use it as a module <br>
4. Add a `{your .mod name}/sysmon` to the `import` line <br>
5. Finished!<br>

## Usage example

You can look into the [main.go](engraulis/main.go) OR <br>
For `engraulis/client`: <br>
You can just use it as a function: <br>
`res := client.Req(address string, delay int (time.Duration), amount int, logging boolean)` <br>


That's all! <br>

# My Plans

Things i'm *planning to* implement: <br>
1. ~~Logging~~ **DONE✅**<br> 
2. Web Dashboard <br>
3. ~~System Monitoring~~ **DONE✅** (Partly. I'll do another metrics beside of RAM Usage soon) <br>
4. ~~Database for system monitoring~~ **DONE✅** <br>
5. MCP Integration (mostly for beginners, I don't like LLMs) <br>
6. ~~Clean the repo and make it public (not only local files)~~ **DONE✅** <br>
7. Other Monitoring features <br>
8. IaC integration (**XML** and **JSON** configs instead of commands)

# Contributing
Wanna contribute? Nice! Check the [CONTRIBUTING.md](CONTRIBUTING.md) file.

# Footer

Thank you for reading this! <br>
Project is distributing under the `BSD 3-Clause` license. <br>
<br>
Danketsu Studio, 2026
