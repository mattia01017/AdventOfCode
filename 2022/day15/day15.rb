# brute force solution, VERY slow (about 45s on my PC)
require './sensors_util.rb'
require './components.rb'

TO_WATCH = 2000000
MAX_COORD = 4000000

sensors = []

STDIN.each_line do |line|
  sensors.append parse(line)

end
beacons = sensors.map { |s| s.beacon }
beacons = beacons.to_set

counter = count_free_spots(sensors, beacons, TO_WATCH)

puts counter

free_y = find_row_w_spot(sensors, beacons)
free_x = find_col_w_spot(sensors, free_y, MAX_COORD)

puts MAX_COORD * free_x + free_y