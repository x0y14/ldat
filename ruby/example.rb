# typed: true

module Example
  class Name
    extend T::Sig

    sig {params(first: String, middle: String, last: String).void}
    def initialize(first:, middle:, last:)
      @first = first
      @middle = middle
      @last = last
    end
  end
end
