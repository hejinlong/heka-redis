# heka-redis
A redis output plugin for heka

This is a redis LPUSH ouput plugin for <a href="http://hekad.readthedocs.org/"> Heka</a>

Basically, you'll need to edit the cmake/plugin_loader.cmake file and add

```cpp
add_external_plugin(git https://github.com/hejinlong/heka-redis master)
```

And build heka.

Refer to Heka's offical <a href="http://hekad.readthedocs.org/en/latest/installing.html#build-include-externals">Building External Plugins</a> docs for more details.

Example:

```cpp
[RedisOutput]
address = "127.0.0.1:6379"
key = "heka"
message_matcher = "TRUE"
encoder = "PayloadEncoder"
```