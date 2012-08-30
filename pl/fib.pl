#!/usr/bin/perl

sub fib {
    my $n = $_[0];
    my $a = 0;
    my $b = 1;

    for ($i = 0; $i < $n; $i++) {
        ($a, $b) = ($a+$b, $a);
    }
    return $b;
}

print fib(1);
print fib(2);
print fib(3);
print fib(4);
print fib(5);
print fib(6);
print fib(7);
print fib(8);
print fib(9);

