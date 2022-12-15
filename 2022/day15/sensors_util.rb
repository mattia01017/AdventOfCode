def parse(line)
  toks = line.split
  sx = toks[2][2..].to_i
  sy = toks[3][2..].to_i
  bx = toks[8][2..].to_i
  by = toks[9][2..].to_i
  Sensor.new(sx, sy, Beacon.new(bx, by))
end

def count_free_spots(sensors, beacons, row, inf_limit = nil, sup_limit = nil, include_beacons = false)
  ranges = sensors.map { |s| s.covered_range(row) }
  ranges = ranges.filter { |r| not r.nil? }
  ranges = ranges.sort

  counter = 0
  last = inf_limit.nil? ? ranges[0][0] - 1 : inf_limit - 1

  ranges.each do |r|
    break if not sup_limit.nil? and r[0] > sup_limit
    if last < r[1]
      s = [last + 1, r[0]].max
      if sup_limit.nil?
        e = r[1]
      else
        e = [r[1], sup_limit].min
      end

      counter += e - s + 1
      last = e
    end
  end
  unless include_beacons
    counter -= beacons.filter { |b| b.coord.ay == row }.length
  end
  counter
end

def find_col_w_spot(sensors, row, to)
  ranges = sensors.map { |s| s.covered_range(row) }
  ranges = ranges.filter { |r| not r.nil? }

  to.times do |i|
    in_range = false
    ranges.each do |r|
      in_range = in_range(r, i)
      break if in_range
    end
    return i unless in_range
  end
end

def find_row_w_spot(sensors, beacons)
  free_y = 0
  (MAX_COORD + 1).times do |i|
    res = count_free_spots(
      sensors, beacons, i + 1, inf_limit = 0,
      sup_limit = MAX_COORD, include_beacons = true
    )
    if res < MAX_COORD + 1
      free_y = i + 1
      break
    end
  end
  free_y
end

def in_range(range, val)
  val >= range[0] and val <= range[1]
end