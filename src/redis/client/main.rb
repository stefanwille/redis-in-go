require "redis"

redis = Redis.new

puts redis.set("mykey", "hello world")
puts redis.set("yourkey", "goodbye world")
puts redis.get("mykey")
puts "keys:", redis.keys("*")
