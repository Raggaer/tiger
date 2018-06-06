# tiger

[![Build status](https://ci.appveyor.com/api/projects/status/lh43r8owobd0g6fv?svg=true)](https://ci.appveyor.com/project/Raggaer/tiger)

Discord bot for Open Tibia Servers (TFS 1.+)

## Features

- Monster lookup (information, loot, players killed)
- Vocation lookup (information)
- Player lookup (information)
- Latest server deaths
- Highscores
- Solid template system
- Death broadcasts

You can check all the bot commands on the [tiger website](https://tigerbot.org/commands.html)

## Installing

To install **tiger** you can grab the latest source and compile or download an already compiled tiger application ([AppVeyor](https://ci.appveyor.com/project/Raggaer/tiger) or [GitHub releases](https://github.com/Raggaer/tiger/releases))

Create a bot token on Discord

`https://discordapp.com/developers/applications/me`

Allow the bot to join your server

`https://discordapp.com/developers/docs/topics/oauth2#bots`

Configure `config.toml.sample` file:

```toml
[template]
directory = "template/" # Directory where the bot will search for templates
extension = ".tiger" # The exntesion of your templates

[server]
path = "" # Your server location 
address = "" # Your server address with port included

[discord]
prefix = "/" # Prefix for the bot commands
token = "" # Discord bot token
status = "" # Bot status message (Playing...)
channels = ["test", "bot-test"] # Channels where the bot will listen to commands
deathChannels = ["test", "bot-test"] # Channels where the bot will broadcast player deaths
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

## Documentation

You can all the documentation on the [Tiger website](https://tigerbot.org)

## License

**Tiger** is licensed under the **MIT** license