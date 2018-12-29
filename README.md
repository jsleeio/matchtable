# matchtable

## what is this?

`matchtable` reads lines from files and generates a map indicating which lines
are present in which files. This is similar to the standard Unix utility
`column`, except that it accepts any number of files at the input.

## demo

```
$ seq 1 3 > A
$ seq 2 4 > B
$ seq 3 5 > C
$ seq 4 5 > D
$ matchtable [A-D] | column -t
ITEM  A  B  C  D
1     X  -  -  -
2     X  X  -  -
3     X  X  X  -
4     -  X  X  X
5     -  -  X  X
```

## original use case

The first use case for `matchtable` was determining where AWS security groups
were used, as they can be referenced in an ever-expand variety of places,
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

* output sorting
* column aliases
* improved memory efficiency
