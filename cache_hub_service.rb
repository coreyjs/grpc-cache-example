require 'grpc'
require_relative 'lib/cache_services_pb'

class CacheHubService < Cache::CacheHub::Service
  def get_status(req, _unused_call)
    puts 'Received get status request'
    puts req.name
    Cache::StatusResponse.new(message: 'Systems are operational, nothing to see here')
  end

  def upload(req)
    puts 'got something'
    puts req
    Cache::UploadStatus.new(message: 'Test', code: :Ok)
  end
end