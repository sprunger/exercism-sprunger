class Phrase

  def initialize(phrase)
    @phrase = phrase.
      gsub(/[,]/i, ' ').
      gsub(/[^0-9a-z ']/i, '').downcase
  end

  def word_count
    words = Hash.new(0)
    @phrase.split.each { |w| words[w] += 1 }
    words
  end

end
