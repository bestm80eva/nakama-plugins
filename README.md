Nakama Plugins
==============

> A collection of plugins which can be loaded by the Nakama server.

A plugin is a bundle of logic which is executed by [Nakama](https://github.com/heroiclabs/nakama) server. The logic in a plugin can hook into various places in the socket pipeline of a game client connection and react to the events. For example when a user updates their profile a plugin can execute on the update and send it to an analytics service or other external system.

Plugins are the best way for developers to share events in the distributed game server with another server or service. This is an advanced use case which is rare but useful for integration with analytics, monetization, machine learning, and other systems.

All plugins are curated by the Heroic Labs team; please open a PR if you'd like to have yours added.
