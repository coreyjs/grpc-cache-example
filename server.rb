require 'rubygems'
require_relative  'lib/cache_services_pb'
require_relative 'cache_hub_service'

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

CacheHubServer.start
