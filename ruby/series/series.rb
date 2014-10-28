class Series

  attr_reader :string

  def initialize(string)
    @string = string
  end

  def slices(count)
    raise ArgumentError if count > string.length
    result = []
    string.chars.each_cons(count) do |i|
      result << i.map{|x| x.to_i}
    end
    return result
  end
end
