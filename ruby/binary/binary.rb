class Binary
  attr_reader :binary

  def initialize(value)
    @binary = [0]
    @binary = value.scan(/\d/).map(&:to_i) if value =~ /^[0|1]*$/
  end

  def to_decimal()
    total = 0

    binary.reverse.each_with_index do |n, i|
      total += n * 2**i
    end

    total
  end
end

# If we could just use the sweetness of Ruby...
#
# Class Binary
#   attr_reader :binary
#
#   def initialize(value)
#     @binary = value
#   end
#
#   def to_decimal
#     @binary.to_i(2)
#   end
# end
