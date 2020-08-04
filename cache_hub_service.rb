require 'grpc'
require 'byebug'
require_relative 'lib/cache_services_pb'
require_relative 'file_handler'


class CacheHubService < Cache::CacheHub::Service
  # Returns a health status, or heartbeat of the gRPC server
  #
  # == Parameters:
  # req::
  #   The gRPC request object
  # call::
  #   Unused call object.
  #
  # == Returns
  #   An instance of Cache::StatusResponse
  def get_status(req, _unused_call)
    puts 'Received get status request'
    Cache::StatusResponse.new(message: 'Systems are operational, nothing to see here')
  end

  # Uploads a file stream from the calling gRPC client.  This
  # will then rebuild the file, store it in /tmp/, and chunk
  # up the file into smaller sizes to store into memcache.
  #
  # == Parameters:
  # req::
  #   The gRPC request object
  # call::
  #   Unused call object.
  #
  # == Returns
  #   An instance of Cache::UploadStatus
  def upload(call)
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

    # cleanup
    delete_file(temp_file_name)

    Cache::UploadStatus.new(Message: 'Test', Code: :Ok)
  end

  # Retrieves a file from the cache, if it exists
  #
  # == Parameters:
  # req::
  #   The gRPC request object
  # call::
  #   Unused call object.
  #
  # == Returns
  #   An instance of Cache::FileStatusResponse.
  #   This will contain:
  #     code: :CacheMiss or :Exists
  #     location: The tmp file location
  #     digest: The md5 digest to ensure data validity
  def get_file(req, call)
    temp_file_name = "/tmp/#{req.filename}"
    status, checksum = retrieve(req.filename, temp_file_name)

    Cache::FileStatusResponse.new(code: :CacheMiss, location: nil, digest: nil) unless status

    Cache::FileStatusResponse.new(code: :Exists, location: temp_file_name, digest: checksum)
  end

  private
  # Stores/Caches a file into the memcache instance
  #
  # == Parameters
  # key::
  #   The primary key prefix to store the file.  This is usually
  #   just the filename, i.e. myfile.txt
  # file_name::
  #   The name of the file that has the data we need to cache.  This is
  #   going to be in the /tmp/ directory from when we rebuilt the file.
  #
  # == Returns:
  #   nil
  def store(key, file_name)
    FileHandler.new.store(key, file_name)
  end

  # Retrieves a file from the cache
  #
  # == Parameters:
  # name::
  #   The base name, or key, of the file to get.  This should
  #   not be a fully qualified path, just the filename.txt
  # outfile::
  #   The name of the cached file to be written to.  This will be
  #   stored in /tmp/ for now
  #
  # == Returns
  #   A tuple of the status, and the md5 checksum
  #   [true, 394343sjfkjd930]
  def retrieve(name, outfile)
    FileHandler.new.retrieve(name, outfile)
  end


  # Deletes any tmp file we created when rebuilding the data
  #
  # == Parameters:
  # file_name::
  #   The name of the file that is to be deleted.
  def delete_file(file_name)
    begin
      File.open(file_name, 'r') do |f|
        File.delete(f)
      end
    end
  end
end