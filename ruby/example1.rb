# typed: true

extend T::Sig

sig {params(name: String).returns(String)}
def say_hello(name)
  return "hello, " + name
end


