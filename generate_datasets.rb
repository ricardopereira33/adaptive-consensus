require 'timeout'

initial_default_delta       = 1
initial_max_tries           = 1
initial_percentage_miss     = 1
initial_percentage_faults   = 0
initial_probability_to_fail = 0
initial_latency             = 10

nodes     = 10
mutations = ['early', 'ring', 'centralized', 'gossip']
faults    = [true, false]

puts 'Start!'

nodes.step(100, 20) do |nodes_number|
  initial_default_delta.step(5, 0.5) do |default_delta|
    initial_max_tries.step(5, 1) do |max_tries|
      initial_percentage_miss.step(10, 2) do |percentage_miss|
        initial_percentage_faults.step(30, 5) do |percentage_fault|
          initial_probability_to_fail.step(12, 2) do |probability_to_fail|
            initial_latency.step(200, 20) do |latency|
              mutations.each do |mutation|
                `./bin/simulation #{mutation} #{nodes_number} #{default_delta} #{max_tries} #{percentage_miss} #{is_faulty} #{latency} #{percentage_fault} #{probability_to_fail} false`
              end
            end
          end
        end
      end
    end
  end
end

puts 'End.'