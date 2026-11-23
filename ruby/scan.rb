# Send one document image to the doc.cheap API and print the main fields.
# Standard library only: ruby scan.rb path/to/image.jpg
require "base64"
require "json"
require "net/http"

API = URI("https://api.doc.cheap/v1/scans")
# The public sandbox key from the docs. Set DOC_CHEAP_API_KEY to use your own.
KEY = ENV.fetch("DOC_CHEAP_API_KEY", "sk_sandbox_public")

abort "usage: ruby scan.rb <image.jpg|image.png>" unless ARGV.length == 1

image = Base64.strict_encode64(File.binread(ARGV[0]))
request = Net::HTTP::Post.new(API, "Authorization" => "Bearer #{KEY}",
                                   "Content-Type" => "application/json")
request.body = JSON.generate(image: image)
response = Net::HTTP.start(API.host, API.port, use_ssl: true, read_timeout: 60) do |http|
  http.request(request)
end
result = JSON.parse(response.body)

unless response.is_a?(Net::HTTPSuccess)
  # Every error body carries a stable code; branch on it, not on the status.
  error = result["error"]
  abort "#{response.code} #{error['code']}: #{error['message']} (request #{error['request_id']})"
end

# A photo that holds no readable document is a 200 too: meta.status says so.
meta = result["meta"]
document = result["document"] || {}
holder = result["holder"] || {}
puts "status:      #{meta['status']}"
puts "billed:      #{meta['billed']}"
puts "kind:        #{document['kind'].inspect}"
puts "country:     #{document['country'].inspect}"
puts "number:      #{document['number'].inspect}"
puts "expiry_date: #{document['expiry_date'].inspect}"
puts "surname:     #{holder['surname'].inspect}"
puts "given_names: #{holder['given_names'].inspect}"
puts "birth_date:  #{holder['birth_date'].inspect}"
puts "mrz:         #{result['mrz']['status']}"
