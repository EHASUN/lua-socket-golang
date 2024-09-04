
# About

This project provides socket server using Lua programming to handle multiple connection to ensure the server can handling multiple socket connection in sametime and I have created a golang socket client and interated with lua socket server.

#go_client.go
I have created a golang socket client

#main.lua
I have connected go socket with lua


# Usage
## copas echo server
This implements a basic openresty server via Websockets protocol. Once you are connected with the server, all messages you send will be returned ('echoed') by the server immediately.

```lua
local copas = require'copas'

-- create a openresty webserver and start listening
local server = require "resty.websocket.server"
{
  -- listen on port 8080
  port = 8080,
  -- the protocols field holds
  --   key: protocol name
  --   value: callback on new connection
  protocols = {
    -- this callback is called, whenever a new client connects.
    -- ws is a new websocket instance
    echo = function(ws)
      while true do
        local message = ws:receive()
        if message then
           ws:send(message)
        else
           ws:close()
           return
        end
      end
    end
  }
}

-- use the copas loop
copas.loop()
```


--Once connected with the server, all messages you send will be returned ('echoed') by the server immediately.

```lua
local ev = require'ev'

-- create a openresty webserver and start listening
local server = require "resty.websocket.server".server.ev.listen
{
  -- listen on port 8080
  port = 8080,
  -- the protocols field holds
  --   key: protocol name
  --   value: callback on new connection
  protocols = {
    -- this callback is called, whenever a new client connects.
    -- ws is a new websocket instance
    echo = function(ws)
      ws:on_message(function(ws,message)
          ws:send(message)
        end)

      -- this is optional
      ws:on_close(function()
          ws:close()
        end)
    end
  }
}

-- use the lua-ev loop
ev.Loop.default:loop()

```

## Running test-server examples

The folder test-server contains two re-implementations: test-server.c example.

```shell
cd test-server
lua test-server-ev.lua
```

```shell
cd test-server
lua test-server-copas.lua
```

Connect to the from Javascript (e.g. chrome's debugging console) like this:
```Javascript
var echoWs = new WebSocket('ws://127.0.0.1:8002','echo');
```

# Dependencies

The client and server modules depend on:

  - luasocket
  - luabitop (if not using Lua 5.2 nor luajit)
  - luasec
  - copas (optionally)
  - lua-ev (optionally)

# Install

```shell
$ git clone git://github.com/EHASUN/lua-websockets.git
$ cd lua-websockets
$ luarocks make rockspecs/lua-websockets-scm-1.rockspec
```

The minifed version has be to be installed manually though.


# Tests

Manual Testing

The first run will take A WHILE.
