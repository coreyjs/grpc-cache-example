require 'dalli'

class CacheBroker
  def initialize(options = nil)
    options ||= {:namespace => "app_v1", :compress => true}
    @client = Dalli::Client.new('localhost:11211', options)
  end

  def set_key(key, val)
    puts "setting #{key}"
    @client.set(key, val)
  end

  def get_key(key)
    result = @client.get(key)

    puts "Cache miss: #{key}" unless result

    result
  end
end