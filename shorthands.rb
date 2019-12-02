#!/usr/bin/env ruby

TOPICS = ["error", "info", "debug"]

puts "package log"
puts
puts 'import (
        "context"
        "github.com/novakit/log/labels"
      )
'
puts
TOPICS.each do |topic|
puts "// #{topic.capitalize} shorthand for Log with topic #{topic}"
puts "func #{topic.capitalize}(ctx context.Context, message string) {"
puts "  Log(ctx, \"#{topic}\", message)"
puts "}"
puts
puts "// #{topic.capitalize}f shorthand for Logf with topic #{topic}"
puts "func #{topic.capitalize}f(ctx context.Context, format string, items ...interface{}) {"
puts "  Logf(ctx, \"#{topic}\", format, items...)"
puts "}"
puts
puts "// #{topic.capitalize}l shorthand for Logl with topic #{topic}"
puts "func #{topic.capitalize}l(ctx context.Context, addLabels labels.Labels) {"
puts "  Logl(ctx, \"#{topic}\", addLabels)"
puts "}"
puts
puts "// #{topic.capitalize}lf shorthand for Loglf with topic #{topic}"
puts "func #{topic.capitalize}lf(ctx context.Context, addLabels labels.Labels, format string, items ...interface{}) {"
puts "  Loglf(ctx, \"#{topic}\", addLabels, format, items...)"
puts "}"
puts
end