require 'dalli'
require 'optparse'
require_relative 'file_handler'

def store(name, infile)
  FileHandler.new.store(name, infile)
end

def retrieve(name, outfile)
  FileHandler.new.retrieve(name, outfile)
end

# options = {:namespace => "app_v1", :compress => true}
# dc = Dalli::Client.new('localhost:11211', options)
# dc.set('abc', 123)
# value = dc.get('abc')
# puts value
#

p ARGV

if ARGV.length != 3
  puts 'Not enough cli arguments'
end

if ARGV[0] == "store"
  store(ARGV[1], ARGV[2])
elsif ARGV[0] == "retrieve"
  retrieve(ARGV[1], ARGV[2])
end