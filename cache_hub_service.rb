require 'grpc'
require 'byebug'
require_relative 'lib/cache_services_pb'
require_relative 'file_handler'


class CacheHubService < Cache::CacheHub::Service
  def get_status(req, _unused_call)
    puts 'Received get status request'
    Cache::StatusResponse.new(message: 'Systems are operational, nothing to see here')
  end

  def upload(call)
    started, elapsed_time = 0,0
    file_chunks = []
    file_name = nil
    call.each_remote_read do |r|
      file_chunks.append(r.Content)
      file_name ||= r.Identifier
    end

    # TODO:  add failure handling and proper setting on UploadStatus
    data = file_chunks.join('')
    temp_file_name = "/tmp/#{file_name}"
    File.open(temp_file_name, "w+") { |f| f.write(data)}
    store(file_name, temp_file_name)
    Cache::UploadStatus.new(Message: 'Test', Code: :Ok)
  end

  private
  def store(key, file_name)
    FileHandler.new.store(key, file_name)
  end

  def retrieve(name, outfile)
    FileHandler.new.retrieve(name, outfile)
  end
end