# memcached-flusher

Utility to send **flush** command to all servers every 20ms.

## Install

    $ go get -u github.com/flozano/memcached-flusher

## Usage

    $ memcached-flusher 127.0.0.1:11211

or

    $ memcached-flusher 192.168.0.1:11211 192.168.0.2:11211 192.168.0.3:11211
