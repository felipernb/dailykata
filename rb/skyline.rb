#http://acm.uva.es/p/v1/105.html

buildings = [[1,11,5], [2,6,7], [3,13,9], [12,7,16], [14,3,25], [19,18,22], [23,13,29], [24,4,28]]

def skyline(buildings)
	#initialize the skyline setting 0 in all points
	line_size = 0
	buildings.each {|b| line_size = [line_size, b[2]].max}
	line = Array.new(line_size+1).fill(0)
	buildings.each {|b| (b[0]..b[2]-1).each {|x| line[x] = [line[x], b[1]].max}}
	skyline = []
	curr_height = 0
	line.each_with_index do |y, x|
		if y != curr_height
			skyline << [x,y]
			curr_height = y
		end
	end
	return skyline
end

puts skyline(buildings).inspect
