class FoodChainSong

  def initialize()
    @song = String.new

    @verse_start = Array.new
    @accum = Array.new

    @verse_start[1] = "I know an old lady who swallowed a fly.\n"
    @verse_start[2] = "I know an old lady who swallowed a spider.\n" +
                 "It wriggled and jiggled and tickled inside her.\n"
    @verse_start[3] = "I know an old lady who swallowed a bird.\n" +
                 "How absurd to swallow a bird!\n"
    @verse_start[4] = "I know an old lady who swallowed a cat.\n" +
                 "Imagine that, to swallow a cat!\n"
    @verse_start[5] = "I know an old lady who swallowed a dog.\n" +
                 "What a hog, to swallow a dog!\n"
    @verse_start[6] = "I know an old lady who swallowed a goat.\n" +
                 "Just opened her throat and swallowed a goat!\n"
    @verse_start[7] = "I know an old lady who swallowed a cow.\n" +
                 "I don't know how she swallowed a cow!\n"
    @verse_start[8] = "I know an old lady who swallowed a horse.\n" +
                 "She's dead, of course!\n"

    @accum[1] = ""
    @accum[2] = "She swallowed the spider to catch the fly.\n"
    @accum[3] = "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.\n"
    @accum[4] = "She swallowed the cat to catch the bird.\n"
    @accum[5] = "She swallowed the dog to catch the cat.\n"
    @accum[6] = "She swallowed the goat to catch the dog.\n"
    @accum[7] = "She swallowed the cow to catch the goat.\n"
  end

  def verse(number_of_verses)
    @song << @verse_start[number_of_verses]

    return @song if number_of_verses == 8

    (1..number_of_verses).reverse_each { |n| @song << @accum[n] }

    @song << "I don't know why she swallowed the fly. Perhaps she'll die.\n"

    return @song
  end
end
