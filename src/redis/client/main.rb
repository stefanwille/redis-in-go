require "redis"

redis = Redis.new

100_000.times do
redis.set("mykey", "hello world")
end
# puts redis.set("yourkey", "goodbye world")
# puts redis.get("mykey")
# puts "keys:", redis.keys("*")
