# matchtable

## what is this?

`matchtable` reads lines from files and generates a map indicating which lines
are present in which files. This is similar to the standard Unix utility
`comm`, except that it accepts any number of files at the input.

## demo

```
$ cat testdata.sh
#!/bin/sh
seq 1 3 > A
seq 2 4 > B
seq 3 5 > C
seq 4 5 > D
seq 5 5 > E
./matchtable [A-E]


$ ./testdata.sh
ITEM A B C D E
1 X - - - -
2 X X - - -
3 X X X - -
4 - X X X -
5 - - X X X

## options

```
$ ./matchtable -h
Usage of ./matchtable:
  -no-value string
    	string used to indicate an item was NOT present in a column (default "-")
  -separator string
    	string used to separate rendered columns in the output (default " ")
  -sort
    	sort the superset of items lexicographically (default true)
  -yes-value string
    	string used to indicate an item was present in a column (default "X")
```

## original use case

The first use case for `matchtable` was determining where AWS security groups
were used, as they can be referenced in an ever-expanding variety of places,
including but very likely not limited to the below:

* security group rules
* EC2 instances
* elastic network interfaces
* RDS database security groups
* RDS instances
* RDS clusters
* VPC endpoints
* Elasticache clusters
* Lambda functions
* Fargate clusters

## future improvements

* column aliases
* improved memory efficiency
