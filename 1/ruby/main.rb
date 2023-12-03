input_file = File.open("./../input")

def first_and_last_digits(string)
    digits = string.chars.map{|a| a.to_i}.select{|a| a != 0}
    return [digits.first, digits.last]
end

puts input_file.read.split("\n")
        .map{|a| first_and_last_digits(a) }
        .map{|a| (a[0].to_s + a[1].to_s).to_i }
        .sum