To generate go client files:

`protoc -I ../protos ../protos/cache.proto --go_out=plugins=grpc:cache`


To generate ruby server files:

`grpc_tools_ruby_protoc -I protos --ruby_out=lib --grpc_out=lib protos/cache.proto`


To start ruby gRPC server:

`ruby ./server.rb`