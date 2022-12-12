class BigMonkey
    attr_reader :insp_counter, :items, :modulo
    attr_writer :max_val

    def initialize(monkeys, items, incrementer, modulo, throw_to)
        @monkeys = monkeys
        @items = items
        @incrementer = incrementer
        @modulo = modulo
        @throw_to = throw_to
        @insp_counter = 0
    end

    def inspect
        until @items.empty? do
            item = @items.pop
            item = @incrementer.call(item)
            item %= @max_val
            index = item % @modulo == 0? 0 : 1
            @monkeys[@throw_to[index]].give(item)
            @insp_counter += 1
        end
    end

    def give(item)
        @items.unshift item
    end
end

class BigMonkeyParser
    def read(input)
        monkeys = []
        input.split("\n\n").each do |monkey_str|
            lines = monkey_str.split("\n")

            items = parse_items(lines[1])
            incrementer = parse_op(lines[2])
            modulo = parse_modulo(lines[3])
            to_throw = [
                lines[4][-1].to_i,
                lines[5][-1].to_i,
            ]

            monkeys.append BigMonkey.new(
                monkeys, items, incrementer, modulo, to_throw
            )
        end
        max_val = 1
        monkeys.each do |m|
            max_val *= m.modulo
        end
        monkeys.each do |m|
            m.max_val = max_val
        end
    end

    private
    def parse_items(line)
        items = line[line.index(":")+1..].split(", ")
        items = items.map { |x| x.to_i }
    end

    private
    def parse_op(line)
        op_str = line[line.index("old")+4..].split(" ")
        if op_str[0] == "*"
            if op_str[1] == "old"
                incrementer = -> (n) { n * n }
            else
                incrementer = -> (n) { n * op_str[1].to_i }
            end
        else
            incrementer = -> (n) { n + op_str[1].to_i }
        end
    end

    private 
    def parse_modulo(line)
        line[line.rindex(" ")+1..].to_i
    end
end