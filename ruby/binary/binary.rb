class Binary
  attr_reader :binary

  def initialize(value)
    @binary ||= value.to_i.to_s.scan(/\d/)
  end

  def to_decimal()
    total = 0

    binary.reverse.each_with_index do |n, i|
      total += n.to_i * 2**i
    end

    total
  end
end

# If we could just use the sweetness of Ruby
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
