require 'timeout'

metrics = {
  "edge_case_latency"=>{
    "default_delta"       => 2.0,
    "max_tries"           => 5,
    "percentage_miss"     => 0.0,
    "percentage_faults"   => 0.0,
    "probability_to_fail" => 0.0,
    "bandwidth"           => 250,
    "latency"             => 650.0
  },
  "edge_case_bandwidth" => {
    "default_delta"       => 2.0,
    "max_tries"           => 5,
    "percentage_miss"     => 0.0,
    "percentage_faults"   => 0.0,
    "probability_to_fail" => 0.0,
    "bandwidth"           => 500,
    "latency"             => 100.0
  },
  "edge_case_faults" => {
    "default_delta"       => 2.0,
    "max_tries"           => 5,
    "percentage_miss"     => 8.0,
    "percentage_faults"   => 40.0,
    "probability_to_fail" => 40.0,
    "bandwidth"           => 250,
    "latency"             => 100.0
  },
  "edge_case_lost_messages_and_faults" => {
    "default_delta"       => 2.0,
    "max_tries"           => 5,
    "percentage_miss"     => 60.0,
    "percentage_faults"   => 30.0,
    "probability_to_fail" => 30.0,
    "bandwidth"           => 250,
    "latency"             => 100.0
  },
  "edge_case_high_default_delta" => {
    "default_delta"       => 5.0,
    "max_tries"           => 5,
    "percentage_miss"     => 0.0,
    "percentage_faults"   => 0.0,
    "probability_to_fail" => 0.0,
    "bandwidth"           => 250,
    "latency"             => 100.0
  },
  "edge_case_low_default_delta" => {
    "default_delta"       => 1.0,
    "max_tries"           => 5,
    "percentage_miss"     => 0.0,
    "percentage_faults"   => 0.0,
    "probability_to_fail" => 0.0,
    "bandwidth"           => 250,
    "latency"             => 100.0
  }
}

nodes     = 50
mutations = ['early', 'centralized', 'gossip', 'ring', 'old_ring']
# improving ring mutation

puts 'Start!'

nodes.step(250, 50) do |nodes_number|
  metrics.each do |key, value|
    mutations.each do |mutation|
      begin
        command = "./bin/simulation #{mutation} #{nodes_number} #{value['default_delta']} #{value['max_tries']} #{value['percentage_miss']} #{value['latency']} #{value['bandwidth']} #{value['percentage_faults']} #{value['probability_to_fail']} false"
        pid     = Process.spawn(command)

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

puts 'End.'
