require 'timeout'

initial_default_delta       = 0.5
initial_max_tries           = [3, 5]
initial_percentage_miss     = 0
initial_percentage_faults   = 0
initial_probability_to_fail = 0
initial_bandwidth           = 20
initial_latency             = 25

nodes     = 10
mutations = ['early', 'ring', 'centralized', 'gossip']

puts 'Start!'

nodes.step(100, 30) do |nodes_number|
  initial_default_delta.step(2, 0.5) do |default_delta|
    initial_max_tries.each do |max_tries|
      initial_percentage_miss.step(8, 4) do |percentage_miss|
        initial_percentage_faults.step(20, 5) do |percentage_fault|
          initial_probability_to_fail.step(10, 5) do |probability_to_fail|
            initial_bandwidth.step(200, 40) do |bandwidth|
              initial_latency.step(250, 50) do |latency|
                mutations.each do |mutation|
                  if not (percentage_fault == 0 and probability_to_fail > 0)
                    command = "./bin/simulation #{mutation} #{nodes_number} #{default_delta} #{max_tries} #{percentage_miss} #{latency} #{bandwidth} #{percentage_fault} #{probability_to_fail} false"

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
                end
              end
            end
          end
        end
      end
    end
  end
end

puts 'End.'