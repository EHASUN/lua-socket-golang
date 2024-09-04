FROM ubuntu:24.04.1
# install autobahn tests suite (python)
RUN apt-get update -y && apt-get install build-essential libssl-dev python -y
# install lua
ENV LUAROCKS_VERSION=2.0.13
ENV LUAROCKS_BASE=luarocks-$LUAROCKS_VERSION
ENV LUA luajit
ENV LUA_DEV libluajit-5.1-dev 
ENV LUA_VER 5.1
ENV LUA_SFX jit
ENV LUA_INCDIR /usr/include/luajit-2.0

ADD . /lua-websockets
WORKDIR /lua-websockets
RUN luarocks make rockspecs/lua-websockets-scm-1.rockspec 
RUN ./test.sh

