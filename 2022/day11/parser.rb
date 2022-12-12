require 'monkeys'
class MonkeyParser
    def read(input, big=false)
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

            if big
                monkeys.append BigMonkey.new(
                  monkeys, items, incrementer, modulo, to_throw
                )
            else
                monkeys.append Monkey.new(
                  monkeys, items, incrementer, modulo, to_throw
                )
            end
        end
        if big
            max_val = 1
            monkeys.each { |m| max_val *= m.modulo }
            monkeys.each { |m| m.max_val = max_val }
        end
        monkeys
    end

    def parse_items(line)
        items = line[line.index(":")+1..].split(", ")
        items = items.map { |x| x.to_i }
    end

    def parse_op(line)
        op_str = line[line.index("old")+4..].split(" ")
        if op_str[0] == "*"
            if op_str[1] == "old"
                return -> (n) { n * n }
            end
            -> (n) { n * op_str[1].to_i }
        else
            -> (n) { n + op_str[1].to_i }
        end
    end

    def parse_modulo(line)
        line[line.rindex(" ")+1..].to_i
    end

    private :parse_items, :parse_op, :parse_items
end