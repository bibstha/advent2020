i = File.read("input")
a = i.split("\n\n").map! do |x|
  x.gsub("\n", " ") # .tap { |a| pp a }
    .split(" ")
    .map!{ |x| x.split(":") }.to_h
end

k = %w(byr iyr eyr hgt hcl ecl pid)
c = a.count do |hash|
  x = (k - hash.keys).none?
  if !x
    puts "#{hash["hcl"]} #{hash["pid"]} INVALID"
  end
  x
end
puts c

