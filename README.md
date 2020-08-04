### This is an example project of connecting two distinct applications over gRPC.  


The gRPC server is written in ruby, and can be ran from the main directory, `ruby ./server.rb`.  This starts a gRPC server on port `50051`.

```ruby
class CacheHubServer
  class << self
    def start
      start_grpc_server
    end

    private
    def start_grpc_server
      puts 'cache hub server up'
      @server = GRPC::RpcServer.new
      @server.add_http2_port('0.0.0.0:50052', :this_port_is_insecure)
      @server.handle(CacheHubService)
      @server.run_till_terminated
    end
  end
end
```
---

There are currently only 2 methods defined in the `cache.proto` file:

### StatusRequest
- Used for basic health check on the server

### Upload
- Used as the main file transfer stream method

```bash
service CacheHub {
  rpc GetStatus(StatusRequest) returns (StatusResponse) {}
  rpc Upload(stream Chunk) returns (UploadStatus) {}
}
```
---
### Client
The gRPC client in written in Go, and can be found in `/golang_client/`.  You can interact with the client from the command line:

```
> go run client.go -h
  -a string
        action to preform.  status, get, store
  -f string
        file name
exit status 2
```

### To store a file in memcache:

`> go run client.go -a=store -f=/path/to/file/myfile.pdf`


### To get a file from the cache:

`go run client.go -a=get -f=myfile.pdf`

`File Location: /tmp/myfile.pdf -- md5: 61e4b830ec5baea0fa9ff430e8ddcfa7% `
 
Note: this returns a file path location from the local system, for learning purposes only.


---

### The Ruby FileHandler and CacheBroker classes
...coming sooon...




---
### Misc Notes:

To generate go client files:

`protoc -I ../protos ../protos/cache.proto --go_out=plugins=grpc:cache`


To generate ruby server files:

`grpc_tools_ruby_protoc -I protos --ruby_out=lib --grpc_out=lib protos/cache.proto`


To start ruby gRPC server:

`ruby ./server.rb`
