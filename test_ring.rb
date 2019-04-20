require 'timeout'

metrics = {
  'ring_without_faults_low_default_delta'=> {
      'default_delta'       => 1.0,
      'max_tries'           => 3,
      'percentage_miss'     => 0.0,
      'percentage_faults'   => 0.0,
      'probability_to_fail' => 0.0,
      'bandwidth'           => 200,
      'latency'             => 125.0
  },
  'ring_without_faults_normal'=> {
      'default_delta'       => 3.0,
      'max_tries'           => 3,
      'percentage_miss'     => 0.0,
      'percentage_faults'   => 0.0,
      'probability_to_fail' => 0.0,
      'bandwidth'           => 100,
      'latency'             => 125.0
  },
  'ring_without_faults_high_latency'=> {
      'default_delta'       => 3.0,
      'max_tries'           => 3,
      'percentage_miss'     => 0.0,
      'percentage_faults'   => 0.0,
      'probability_to_fail' => 0.0,
      'bandwidth'           => 100,
      'latency'             => 375.0
  },
   'ring_without_faults_large_bandwidth'=> {
      'default_delta'       => 3.0,
      'max_tries'           => 3,
      'percentage_miss'     => 0.0,
      'percentage_faults'   => 0.0,
      'probability_to_fail' => 0.0,
      'bandwidth'           => 300,
      'latency'             => 125.0
  },
  'ring_with_faults_low_default_delta'=> {
      'default_delta'       => 1.0,
      'max_tries'           => 3,
      'percentage_miss'     => 8.0,
      'percentage_faults'   => 15.0,
      'probability_to_fail' => 15.0,
      'bandwidth'           => 200,
      'latency'             => 125.0
  },
  'ring_with_faults_normal'=> {
      'default_delta'       => 3.0,
      'max_tries'           => 3,
      'percentage_miss'     => 8.0,
      'percentage_faults'   => 15.0,
      'probability_to_fail' => 15.0,
      'bandwidth'           => 100,
      'latency'             => 125.0
  },
  'ring_with_a_lot_of_faults'=> {
      'default_delta'       => 3.0,
      'max_tries'           => 3,
      'percentage_miss'     => 16.0,
      'percentage_faults'   => 30.0,
      'probability_to_fail' => 30.0,
      'bandwidth'           => 200,
      'latency'             => 125.0
  },
  'ring_with_faults_high_latency'=> {
      'default_delta'       => 3.0,
      'max_tries'           => 3,
      'percentage_miss'     => 8.0,
      'percentage_faults'   => 15.0,
      'probability_to_fail' => 15.0,
      'bandwidth'           => 100,
      'latency'             => 375.0
  },
  'ring_with_faults_large_bandwidth'=> {
      'default_delta'       => 3.0,
      'max_tries'           => 3,
      'percentage_miss'     => 8.0,
      'percentage_faults'   => 15.0,
      'probability_to_fail' => 15.0,
      'bandwidth'           => 300,
      'latency'             => 125.0
  }
}

nodes     = 50
mutations = ['ring', 'old_ring']
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