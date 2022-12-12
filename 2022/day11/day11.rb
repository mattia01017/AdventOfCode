$LOAD_PATH << '.'

require 'Part1'
require 'Part2'

first_parser = MonkeyParser.new
second_parser = BigMonkeyParser.new

$/ = 'END'
input = gets

monkeys1 = first_parser.read(input)
monkeys2 = second_parser.read(input)

20.times do
    monkeys1.map { |m| m.inspect }
end

10000.times do
    monkeys2.map { |m| m.inspect }
end

puts "First monkeys: " + monkeys1.map { |m| m.insp_counter }.to_s
puts "Second monkeys: " + monkeys2.map { |m| m.insp_counter }.to_s