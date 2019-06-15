require 'csv'
require 'byebug'
require 'pp'

data   = {}
i      = 0
header = []

CSV.foreach('global_results_3v.csv') do |row|
  if i != 0
    info = {
      first: {
        mutation:          row[9],
        bandwidthExceeded: row[8],
        time:              row[10]
      }
    }
    row_data   = {}
    row_data_2 = {}
    row_data_3 = {}
    row_data_4 = {}
    row_data_5 = {}
    row_data_6 = {}
    row_data_7 = {}

    if data[row[0]] == nil
      row_data[row[7]]   = info
      row_data_2[row[6]] = row_data
      row_data_3[row[5]] = row_data_2
      row_data_4[row[4]] = row_data_3
      row_data_5[row[3]] = row_data_4
      row_data_6[row[2]] = row_data_5
      row_data_7[row[1]] = row_data_6
      data[row[0]]       = row_data_7

    elsif data[row[0]][row[1]] == nil
      row_data[row[7]]     = info
      row_data_2[row[6]]   = row_data
      row_data_3[row[5]]   = row_data_2
      row_data_4[row[4]]   = row_data_3
      row_data_5[row[3]]   = row_data_4
      row_data_6[row[2]]   = row_data_5
      data[row[0]][row[1]] = row_data_6

    elsif data[row[0]][row[1]][row[2]] == nil
      row_data[row[7]]   = info
      row_data_2[row[6]] = row_data
      row_data_3[row[5]] = row_data_2
      row_data_4[row[4]] = row_data_3
      row_data_5[row[3]] = row_data_4

      data[row[0]][row[1]][row[2]] = row_data_5

    elsif data[row[0]][row[1]][row[2]][row[3]] == nil
      row_data[row[7]]   = info
      row_data_2[row[6]] = row_data
      row_data_3[row[5]] = row_data_2
      row_data_4[row[4]] = row_data_3

      data[row[0]][row[1]][row[2]][row[3]] = row_data_4

    elsif data[row[0]][row[1]][row[2]][row[3]][row[4]] == nil
      row_data[row[7]]     = info
      row_data_2[row[6]]   = row_data
      row_data_3[row[5]]   = row_data_2
      data[row[0]][row[1]][row[2]][row[3]][row[4]] = row_data_3

    elsif data[row[0]][row[1]][row[2]][row[3]][row[4]][row[5]] == nil
      row_data[row[7]]   = info
      row_data_2[row[6]] = row_data

      data[row[0]][row[1]][row[2]][row[3]][row[4]][row[5]] = row_data_2

    elsif data[row[0]][row[1]][row[2]][row[3]][row[4]][row[5]][row[6]] == nil
      row_data[row[7]] = info
      data[row[0]][row[1]][row[2]][row[3]][row[4]][row[5]][row[6]] = row_data

    elsif data[row[0]][row[1]][row[2]][row[3]][row[4]][row[5]][row[6]][row[7]] == nil
      data[row[0]][row[1]][row[2]][row[3]][row[4]][row[5]][row[6]][row[7]] = info
    elsif
      current_info = data[row[0]][row[1]][row[2]][row[3]][row[4]][row[5]][row[6]][row[7]]

      if current_info[:first][:time].to_f > info[:first][:time].to_f
        data[row[0]][row[1]][row[2]][row[3]][row[4]][row[5]][row[6]][row[7]] = info

        if current_info[:first][:mutation] != info[:first][:mutation]
          data[row[0]][row[1]][row[2]][row[3]][row[4]][row[5]][row[6]][row[7]][:second] = {
            mutation:          current_info[:first][:mutation],
            bandwidthExceeded: current_info[:first][:bandwidthExceeded],
            time:              current_info[:first][:time]
          }
        end
      end
    end
  else
    header = row
  end

  i += 1
end

c = 0

CSV.open('best_results.csv', 'wb') do |csv|
  csv << header

  data.each do |nodes, values_nodes|
    values_nodes.each do |defaultDelta, values_defaultDelta|
      values_defaultDelta.each do |maxTries, values_maxTries|
        values_maxTries.each do |percentageMiss, values_percentageMiss|
          values_percentageMiss.each do |percentageFaults, values_percentageFaults|
            values_percentageFaults.each do |probabilityToFail, values_probabilityToFail|
              values_probabilityToFail.each do |latency, values_latency|
                values_latency.each do |bandwidth, values_bandwidth|
                  val  = values_bandwidth[:first]

                  if values_bandwidth[:second]
                    val2 = values_bandwidth[:second]

                    # puts "(#{val2[:time].to_f - val[:time].to_f}) <= #{(val[:time].to_f * 0.1)} = #{(val2[:time].to_f - val[:time].to_f) <= (val[:time].to_f * 0.1)}"

                    c += 1 if val2[:mutation] == 'centralized'

                    if (val2[:time].to_f - val[:time].to_f) <= (val[:time].to_f * 0.15)
                      val = val2
                    end
                  end

                  csv << [nodes, defaultDelta, maxTries, percentageMiss, percentageFaults, probabilityToFail, latency, bandwidth, val[:bandwidthExceeded], val[:mutation], val[:time]]
                end
              end
            end
          end
        end
      end
    end
  end
end

puts c