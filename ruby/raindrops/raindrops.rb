class Raindrops
  def self.convert(value)
    result = ''
    result << 'Pling' if (value % 3).zero?
    result << 'Plang' if value % 5 == 0
    result << 'Plong' if value % 7 == 0

    result.empty? ? value.to_s : result
  end
end
