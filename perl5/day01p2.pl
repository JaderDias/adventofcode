#!/usr/bin/perl
use strict;
use warnings;

open my $file, "<", "../day01.input"
	or die $!;

my $floor = 0;
my $position = 1;
while(my $line = <$file>) {
	chomp $line;
	for my $char (split //, $line) {
		$floor++ if $char eq "(";
		$floor-- if $char eq ")";
		if ($floor == -1) {
			print "$position\n"; 
		}

		$position++;
	}
}
