# heka-redis

This is a Redis LPUSH ouput plugin for [Heka](http://hekad.readthedocs.org/).

Basically, you'll need to edit the `cmake/plugin_loader.cmake` file and add

```cpp
add_external_plugin(git https://github.com/hejinlong/heka-redis master)
```

and build heka.

Refer to Heka's offical [Building External Plugins](http://hekad.readthedocs.org/en/latest/installing.html#build-include-externals) docs for more details.

>##Motto of Go: “Do more with less code”

Example:

```cpp
[RedisOutput]
address = "127.0.0.1:6379"
key = "heka"
message_matcher = "TRUE"
encoder = "PayloadEncoder"
```