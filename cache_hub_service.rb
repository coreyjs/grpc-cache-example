require 'grpc'
require_relative 'lib/cache_services_pb'

class CacheHubService < Cache::CacheHub::Service
  def get_status(req, _unused_call)
    puts 'Received get status request'
    Cache::StatusResponse.new
  end
end