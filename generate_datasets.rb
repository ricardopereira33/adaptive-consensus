require 'timeout'

def run(command)
  pid = Process.spawn(command)
  begin

    Timeout.timeout(300) do
      Process.wait(pid)
    end
  rescue Timeout::Error
    puts "Timeout - #{mutation}"
    Process.kill('TERM', pid)
  end
end

def execute(mutation, nodes, default_delta, max_tries, percentage_miss, percentage_faults, probability_to_fail, bandwidth, latency)
  nodes.each do |nodes_number|
    default_delta.each do |delta|
      max_tries.each do |max|
        percentage_miss.each do |percentage|
          percentage_faults.each do |percentage_fault|
            probability_to_fail.each do |probability|
              bandwidth.each do |bandw|
                latency.each do |lat|
                  if !(percentage_fault == 0 && probability > 0) || !(nodes_number > 50 && delta < 2 && mutation == 'ring')
                    run("./bin/simulation #{mutation} #{nodes_number} #{delta} #{max} #{percentage} #{lat} #{bandw} #{percentage_fault} #{probability} false")
                  end
                end
              end
            end
          end
        end
      end
    end
  end
end

puts 'Start!'

nodes               = [30, 50, 100, 200]
default_delta       = [1, 2, 5]
max_tries           = [3, 8]
percentage_miss     = [0, 10, 30]
percentage_faults   = [0, 10, 30]
probability_to_fail = [0, 10, 30]
bandwidth           = [50, 150, 300]
latency             = [50, 200, 400]

puts 'Early'

execute('early', nodes, default_delta, max_tries, percentage_miss, percentage_faults, probability_to_fail, bandwidth, latency)

puts 'Ring'

execute('ring', nodes, default_delta, max_tries, percentage_miss, percentage_faults, probability_to_fail, bandwidth, latency)

puts 'Old Ring'

execute('old_ring', nodes, default_delta, max_tries, percentage_miss, percentage_faults, probability_to_fail, bandwidth, latency)

puts 'Centralized'

execute('centralized', nodes, default_delta, max_tries, percentage_miss, percentage_faults, probability_to_fail, bandwidth, latency)

puts 'Gossip'

execute('gossip', nodes, default_delta, max_tries, percentage_miss, percentage_faults, probability_to_fail, bandwidth, latency)

puts 'End.'