require "redis"

redis = Redis.new

puts redis.set("mykey", "hello world")
puts redis.get("mykey")
