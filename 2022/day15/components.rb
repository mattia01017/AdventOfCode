require "set"

class Coord
  attr_accessor :ax, :ay

  def initialize(x, y)
    @ax = x
    @ay = y
  end

  def hash
    @ax.hash ^ @ay.hash
  end

  def ==(other)
    (@ax == other.ax and @ay == other.ay)
  end

  alias_method :eql?, :==
end

class Beacon
  attr_reader :coord

  def initialize(x, y)
    @coord = Coord.new(x, y)
  end

  def hash
    @coord.hash
  end

  def ==(other)
    @coord == other.coord
  end

  alias_method :eql?, :==
end

class Sensor
  attr_reader :beacon

  def initialize(x, y, beacon)
    @coord = Coord.new(x, y)
    @beacon = beacon
  end

  def covered_range(y)
    d = distance
    diff = (@coord.ay - y).abs
    if diff <= d
      [@coord.ax - (d - diff), @coord.ax + (d - diff)]
    else
      nil
    end
  end

  private

  def distance
    (@coord.ax - @beacon.coord.ax).abs + (@coord.ay - @beacon.coord.ay).abs
  end
end