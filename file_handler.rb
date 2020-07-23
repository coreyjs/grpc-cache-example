require 'digest/md5'

require_relative './cache_broker'


class FileHandler
  # This class handles the storing/retrieving the file as requsted

  def initialize(cache_broker = nil)
    @mode = "rb"
    @chunk_size = 999
    cache_broker ||=CacheBroker.new
    @cache_broker = cache_broker
  end

  def store(name, file)
    puts "Storing File #{name}"
    digest = get_md5(file)
    puts digest

    # begin file chunking to fit into memcache
    chunks = chunks(file)
    puts "Total File Chunks: #{chunks.length}"
    set_in_cache(chunks: chunks, digest: digest, key: name)
  end

  def retrieve(name, output_file)
    puts "Retrieving file: #{name}"
    data, checksum = get_from_cache(name, output_file)

    checksum_status = validate_cache(original_checksum: checksum, data: data)
    puts "File Integrity Check: #{checksum_status}"
  end

  private
  def validate_cache(original_checksum:, data:)
    retrieved_data_digest = Digest::MD5.hexdigest(data)
    retrieved_data_digest == original_checksum
  end

  def get_file_data(file)
    # Open file with 'rb', readonly binary mode,
    # :param file:
    # :return: return binary contents of file
    file = File.open(file, @mode)
    data = file.read
    puts "File Size: #{File.size(file)}"
    file.close
    data
  end

  def get_md5(file)
    data = get_file_data(file)
    Digest::MD5.hexdigest(data)
  end

  def chunks(file)
    puts '.chunks()'
    chunks = []
    #f = File.open(file, @mode)
    File.open(file, @mode).each(nil, @chunk_size) do |chunk|
      chunks.append(chunk)
    end
    chunks
  end

  def set_in_cache(chunks:, digest:, key:)
    # We store the digest for the file as  key:digest
    # We store the chunk size for the file as key:chunks
    # Each chunk will be stored in the cache as key:i, where i
    # is the index of the chunk
    puts '.set_in_cache()'
    @cache_broker.set_key("#{key}:digest", digest)
    @cache_broker.set_key("#{key}:chunks", chunks.length)
    chunks.each_with_index do |c, i|
      @cache_broker.set_key("#{key}:#{i}", c)
    end
  end

  def get_from_cache(key, output_file)
    # get chunk size from the cache
    original_file_chunk_length = @cache_broker.get_key("#{key}:chunks").to_i
    checksum = @cache_broker.get_key("#{key}:digest")
    puts "Checksum: #{checksum}.  Chunk Length: #{original_file_chunk_length}"

    chunks = []
    (0..original_file_chunk_length-1).each do |i|
      chunk = @cache_broker.get_key("#{key}:#{i}")
      chunks.append(chunk)
    end

    cache_data = chunks.join('')
    [cache_data, checksum]
  end
end