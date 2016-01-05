require "redis"

redis = Redis.new

# 100_000.times do
# redis.set("mykey", "hello world")
# end
puts redis.set("mykey", "hello world")
puts redis.get("mykey")
puts "keys:", redis.keys("*")
# puts redis.del("mykey")
# puts "keys:", redis.keys("*")
redis.save
puts redis.dbsize
