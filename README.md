# Bunny

A BSD licenced, go-powered CLI tool for publishing and subscribing to
RabbitMQ, loosely based off https://github.com/algesten/rabbiteer-rs
which doesn't support TLS without major surgery.

## Install

Install bunny:

    $ go get -v github.com/skunkwerks/bunny

## Command format

The command has two "modes", `publish` and `subscribe`.

    $ bunny [general opts] publish/subscribe [specific opts]

### General options


    $ bunny --help
    ...
    OPTIONS:
        -h, --host <host>            RabbitMQ host [default: 127.0.0.1]
        -p, --password <password>    Password to authenticate with [default: guest]
        -P, --port <port>            Port to connect to [default: 5672]
        -U, --url <url>              AMQP connection url (amqp://user:pass@host:port/vhost)
        -u, --user <user>            User to authenticate with [default: guest]
        -v, --vhost <vhost>          Virtual host [default: ]

### Environment variables

#### AMQP_URL

The connection string can be set using the `AMQP_URL` environment
variable.

```
$ export AMQP_URL="amqp://joe:secret@myspecialhost/somevhost"
$ bunny subscribe -e ttninjs-batch
```

#### CONF file

The connection can be specified in a JSON file pointed out by `CONF`
environment variable.

```
$ cat conf-localhost.json
{
    "amqp": {
        "connection": {
            "host": "localhost",
            "vhost": "docker",
            "login": "admin",
            "password": "admin"
        }
    }
}
$ export CONF=conf-localhost.json
$ bunny subscribe -e ttninjs-batch
```

## Publish

Publishing pushes data from stdin or a file to an exchange.


    $ bunny publish --help
    ...
    FLAGS:
        --rpc        Publish as RPC with replyTo and wait for reply.
    OPTIONS:
        -c, --content-type <content_type>    Content type such as application/json. Inferred from filename if
                                             possible.
        -e, --exchange <exchange>            Exchange to publish to [default ""]
        -f, --file <file>                    Filename (- is stdin) [default: -]
        -H, --header <header>...             Header on the form "My-Header: Value"
        -r, --routing-key <routing_key>      Routing key [default: ]
        -z, --priority <priority>            Priority
        -t, --rpctimeout <rpctimeout>        RPC timeout in milliseconds

### Example

#### From stdin

Provide the content-type as arg.

    $ echo "{\"panda\":true}" | \
        bunny -u admin -p admin -v prod publish -e myexchange -c application/json

#### From a file

Content-type is inferred if possible.

    $ bunny -u admin -p admin -v prod publish -e myexchange -H "batch: true" -f ./foo.json

#### Supports RabbitMQ style RPC

Using the `replyTo` header.

    $ CONF=conf.json bunny publish -e myservice -r somecall --rpc -f ./foo.json

Calls `myservice/somecall` using the contents of file `foo.json` and sets up
a `replyTo` header and waits the the rpc reply. The reply will be printed
to stdout.

## Subscribe

Subscribing binds an anonymous queue to an exchange and listens to
messages pushed to that exchange. The messages can either be dumped to
stdout or as files to a directory.

    $ bunny subscribe --help
    ...
    FLAGS:
        -i, --info       Include delivery info (and headers).
    OPTIONS:
        -e, --exchange <exchange>          Exchange to subscribe to [default ""]
        -o, --output <output>              Output directory (- is stdout) [default: -]
        -r, --routing-key <routing_key>    Routing key [default: #]


### Example

#### To stdout

Outputs the body of each message to stdout. Makes no conversion
of the body. If the body is binary, you will see gibberish in the
terminal.


    $ bunny -u admin -p admin -v prod subscribe -e myexchange
    ...


##### With delivery info

`-i` means we make a JSON wrapping of the message so we can include
some basic information about the delivery.

This always produces a JSON structure. When the body is binary (as
indicated by `content_type`), the data is base64 encoded and when the
type is `text/...`, we output it as a JSON string.

The format is:

```json
{
    "deliver":{
        "consumer_tag":"amq.ctag-Tjxx5Qy5zAr0o_yiLOsNEg",
        "delivery_tag":1,
        "redelivered":false,
        "exchange":"myexchange",
        "routing_key":"text"
    },
    "props":{
        "content_type":"application/json",
        "headers":{
            "backendUpdate":false,
            "batch":false,
            "index":"dist-text",
            "oper":"index"
        }
    },
    "data": "body"
}
```


    $ bunny -u admin -p admin -v prod subscribe -e myexchange -i
    ...


#### To a directory

With `-o` the body of each individual message is output to a separate
file. The header `fileName` can be supplied by the sender, in which case
that file is (over-)written.


    $ bunny -u admin -p admin -v prod subscribe -e myexchange -o /tmp
    ...
