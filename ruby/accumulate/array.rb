class Array
  def accumulate
    #accumulator = []
    self.each_with_object([]) do |item, accumulator|
      accumulator << yield(item)
    end
    #accumulator
  end
end
