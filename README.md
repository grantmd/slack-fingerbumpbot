slack-fingerbumpbot
=======

A bot for [Slack](https://slack.com) written in Go (golang) that never leaves a finger bump (:point_right:) hanging (:point_left:).

Usage
-----

* Build the code with `go build`

* Start the bot with `./slack-fingerbumpbot` on an internet-accessible server. (Check the output of `./slack-oldbot -h` for configuration options)

* Configure an [Outgoing Webhook](https://my.slack.com/services/new/outgoing-webhook) in your Slack and point it to the place where your bot is running. For example: `http://example.com:8002/`
