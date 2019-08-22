# rmslack
rmslack is simple tool to remove slack comment

## Installation

slackrm or upgrade slackrm with this command.

```bash
go get -u github.com/laqiiz/slackrm
```


## Usage

```bash
slackrm [options]
```

### Example

Get slack comment url to delete.

![copy link](docs/copy_link.png)

Run slackrm command.

```bash
# Basic
slackrm -url <your slack comment url to delete> -token <your slack api access token>
```

If access token is too long then you can set environment variables.

```bash
# Use environment variables
export SLACK_API_TOKEN=<our slack api access token>
slackrm -url <your slack comment url to delete> 
```

### Options

```bash
>slackrm -help
Usage of slackrm:
  -c string
        -c slack channel name
  -channel string
        slack channel name
  -timestamp string
        timestamp of remove target comment (default "0")
  -tk string
        slack access channel
  -token string
        slack api access token
  -ts string
        timestamp of remove target comment (default "0")
  -u string
        -u delete target slack comment url
  -url string
        url is delete target slack comment url
```

## License

This project is licensed under the Apache License 2.0 License - see the [LICENSE](LICENSE) file for details
