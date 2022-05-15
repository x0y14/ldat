if ARGV.size == 0 then
    pp "not found : analyze target (*.rb)"
    exit(1)
end

if ARGV[0].end_with?(".rb") == false then
    pp "target suffix is not '.rb'"
    exit(1)
end

pp RubyVM::AbstractSyntaxTree.parse_file(ARGV[0])
exit(0)
