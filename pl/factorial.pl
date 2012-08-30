#!/usr/bin/perl

sub factorial {
    my $n = $_[0];
    if ($n <= 0) {
        return 1;
    }
    return $n * factorial($n-1)
}

print factorial(1);
print "\n";
print factorial(2);
print "\n";
print factorial(3);
print "\n";
print factorial(4);
print "\n";
print factorial(5);
print "\n";
print factorial(6);
print "\n";
