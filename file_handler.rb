require 'digest/md5'
require_relative './cache_broker'


class FileHandler
  # This class handles the storing/retrieving the file as requested

  def initialize(file_obj = nil, cache_broker = nil)
    @mode = "rb"
    @chunk_size = 999
    cache_broker ||=CacheBroker.new
    @cache_broker = cache_broker
  end

  # The primary entry point into this class, for storing a file in the cache.
  # This will chunk up the given file, generate a checksum, and store it all
  # in memcache
  #
  # == Parameters:
  # key::
  #   The prefix of all the keys associated with the storing
  #   and chunking of the file
  #
  # file_name::
  #   The ...name..of...the...file...to...chunkerize
  #
  # == Returns
  #   nil
  def store(key, file_name)
    puts "Storing File #{key}"
    digest = get_md5(file_name)
    puts digest

    # begin file chunking to fit into memcache
    chunks = chunks(file_name)
    puts "Total File Chunks: #{chunks.length}"
    set_in_cache(chunks: chunks, digest: digest, key: key)
  end

  # The primary entry point into this class, for retrieving a file in the cache.
  # This will load the file, chunk by chunk, rebuild the file, generate
  # a checksum to ensure data consistency, and return
  #
  # == Parameters:
  # name::
  #   The prefix of all the keys associated with the storing
  #   and chunking of the file
  #
  # output_file::
  #   The tmp file to store
  #
  # == Returns
  #   A tuple of data, [all_cached_data_in_file, checksum]
  #   #  If a cachemiss occurs this will return [nil, nil]
  def retrieve(name, output_file)
    # retrieve a file from the cache
    puts "Retrieving file: #{name}"
    data, checksum = get_from_cache(name, output_file)
    return [false, nil] unless data && checksum

    checksum_status = validate_cache(original_checksum: checksum, data: data)
    puts "File Integrity Check: #{checksum_status}"
    [checksum_status, checksum]
  end

  private
  def validate_cache(original_checksum:, data:)
    retrieved_data_digest = Digest::MD5.hexdigest(data)
    retrieved_data_digest == original_checksum
  end

  # Open file with 'rb', readonly binary mode
  #
  # == Parameters:
  # file::
  #   Name of the file to read in data from
  # == Returns:
  #  data - binary contents of file
  def get_file_data(file)
    file = File.open(file, @mode)
    data = file.read
    puts "File Size: #{File.size(file)}"
    file.close
    data
  end

  # Generate an md5 digest from the file data
  #
  # == Parameters:
  # file::
  #   Name of the file on disk
  #
  # == Returns:
  #   Digest::MD5.hexdigest
  def get_md5(file)
    data = get_file_data(file)
    Digest::MD5.hexdigest(data)
  end

  # Chunks up a file into pieces
  #
  # == Parameters:
  # file::
  #   name of the file to chunk up
  #
  # == Returns:
  #  [chunks] - An array of pieces of binary data from the file, based on
  #   @chunk_size
  def chunks(file)
    puts '.chunks()'
    chunks = []
    File.open(file, @mode).each(nil, @chunk_size) do |chunk|
      chunks.append(chunk)
    end
    chunks
  end

  # We store the digest for the file as  key:digest.
  # We store the chunk size for the file as key:chunks.
  # Each chunk will be stored in the cache as key:i, where i
  # is the index of the chunk
  #
  # == Parameters:
  # chunks::
  #   An array of binary data chunks
  # digest:
  #   The md5 checksum
  # key::
  #   The name of the key to store the chunk value in
  #
  # == Returns:
  #   nil
  def set_in_cache(chunks:, digest:, key:)
    puts '.set_in_cache()'
    @cache_broker.set_key("#{key}:digest", digest)
    @cache_broker.set_key("#{key}:chunks", chunks.length)
    chunks.each_with_index do |c, i|
      @cache_broker.set_key("#{key}:#{i}", c)
    end
  end

  # This accesses the file from the cache, chunk by fabulous chunk
  #
  # == Parameters:
  # key::
  #   The key to prefix all keys
  # output_file::
  #   The tmp file that we will store
  #
  # == Returns:
  #  A tuple of data, [all_cached_data_in_file, checksum]
  #  If a cachemiss occurs this will return [nil, nil]
  def get_from_cache(key, output_file)
    # get chunk size from the cache
    original_file_chunk_length = @cache_broker.get_key("#{key}:chunks").to_i
    checksum = @cache_broker.get_key("#{key}:digest")
    puts "Checksum: #{checksum}.  Chunk Length: #{original_file_chunk_length}"

    return [nil, nil] unless original_file_chunk_length && checksum

    chunks = []
    (0..original_file_chunk_length-1).each do |i|
      chunk = @cache_broker.get_key("#{key}:#{i}")
      chunks.append(chunk)
    end

    cache_data = chunks.join('')
    File.open(output_file, "w+") { |f| f.write(cache_data)}
    [cache_data, checksum]
  end
end