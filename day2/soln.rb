def soln1
  c = input.count do |min, max, char, string|
    n = string.count(char)
    min <= n && n <= max
  end
  puts "Soln1: Count = #{c}"
end

def soln2
  c = input.count do |min, max, char, string|
    (string[min-1] == char) ^ (string[max-1] == char)
  end
  puts "Soln2: Count = #{c}"
end

def input
  re = Regexp.new("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")
  File.readlines("input").map! do |str|
    matches = re.match(str)
    [matches[1].to_i, matches[2].to_i, matches[3], matches[4]]
  end
end

soln1
soln2
