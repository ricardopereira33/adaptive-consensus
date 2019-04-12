require 'timeout'

initial_default_delta       = 1
initial_percentage_miss     = 0
initial_percentage_faults   = 0
initial_probability_to_fail = 0
initial_bandwidth           = 100
initial_latency             = 125

max_tries = 3
nodes     = 50
mutations = ['ring', 'old_ring']
# improving ring mutation

puts 'Start!'

nodes.step(200, 50) do |nodes_number|
  initial_default_delta.step(5, 2) do |default_delta|
    initial_percentage_miss.step(16, 8) do |percentage_miss|
      initial_percentage_faults.step(20, 10) do |percentage_fault|
        initial_probability_to_fail.step(10, 5) do |probability_to_fail|
          initial_bandwidth.step(400, 100) do |bandwidth|
            initial_latency.step(400, 125) do |latency|
              mutations.each do |mutation|
                if not (percentage_fault == 0 and probability_to_fail > 0)
                  begin
                    command = "./bin/simulation #{mutation} #{nodes_number} #{default_delta} #{max_tries} #{percentage_miss} #{latency} #{bandwidth} #{percentage_fault} #{probability_to_fail} false"
                    pid     = Process.spawn(command)

                    Timeout.timeout(300) do
                      Process.wait(pid)
                    end
                  rescue Timeout::Error
                    puts "Timeout - #{mutation}"

                    Process.kill('TERM', pid)
                  rescue Expection => e
                    puts "Error: #{e.message}"

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

puts 'End.'