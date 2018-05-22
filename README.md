# tiger

[![Build status](https://ci.appveyor.com/api/projects/status/lh43r8owobd0g6fv?svg=true)](https://ci.appveyor.com/project/Raggaer/tiger)

Discord bot for Open Tibia Servers

## Features

- Monster lookup
- Vocation lookup

## Installing

To install **tiger** you can grab the latest source and compile or download an already compiled tiger application ([AppVeyor](https://ci.appveyor.com/project/Raggaer/tiger) or [GitHub releases](https://github.com/Raggaer/tiger/releases))

Create a bot token on Discord

`https://discordapp.com/developers/applications/me`

Allow the bot to join your server

`https://discordapp.com/developers/docs/topics/oauth2#bots`

Configure `config.toml.sample` file:

```toml
[server]
path = "" # Your server location 

[discord]
prefix = "/" # Prefix for the bot commands
token = "" # Discord bot token
```

By default **tiger** will try to load your database settings from your server `config.lua` file, you can however use the config file for that too:

```toml
[database]
host = ""
user = ""
password = ""
schema = ""
```

Rename `config.toml.sample` to `config.toml` and execute **tiger**