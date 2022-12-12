class Monkey
    attr_reader :insp_counter, :items

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
            item /= 3
            index = item % @modulo == 0? 0 : 1
            @monkeys[@throw_to[index]].give(item)
            @insp_counter += 1
        end
    end

    def give(item)
        @items.unshift item
    end
end

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