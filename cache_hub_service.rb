require 'grpc'
require 'byebug'
require_relative 'lib/cache_services_pb'

class CacheHubService < Cache::CacheHub::Service
  def get_status(req, _unused_call)
    puts 'Received get status request'
    puts req.name
    Cache::StatusResponse.new(message: 'Systems are operational, nothing to see here')
  end

  def upload(call)
    puts 'got something'

    started, elapsed_time = 0,0
    file_chunks = []
    call.each_remote_read do |r|
      file_chunks.append(r.Content)
    end

    data = file_chunks.join('')
    File.open("out.pdf", "w") { |f| f.write(data)}
    byebug

    Cache::UploadStatus.new(Message: 'Test', Code: :Ok)
  end
end