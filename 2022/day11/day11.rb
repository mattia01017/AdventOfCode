$LOAD_PATH << '.'

require 'monkeys'
require 'parser'

parser = MonkeyParser.new

$/ = 'END'
input = gets

monkeys1 = parser.read(input)
monkeys2 = parser.read(input, big=true)

20.times do
    monkeys1.map { |m| m.inspect }
end

10000.times do
    monkeys2.map { |m| m.inspect }
end

puts "First part monkeys: " + monkeys1.map { |m| m.insp_counter }.to_s
puts "Second part monkeys: " + monkeys2.map { |m| m.insp_counter }.to_s