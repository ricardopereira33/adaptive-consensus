filename = 'to_remove.txt'
file     = File.open(filename, "r")

file.each_line do |line|
  mutation, delay, max_tries, percentage_miss =
    line.match(/(.*) - (\d+) - (\d+) - (\d+).00/).captures
  `rm results-with-faults/csv-datasets/#{mutation}_#{delay}_#{max_tries}_#{percentage_miss}.00_.csv`
end

file.close