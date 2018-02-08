Melter takes a command line consisting of an optional -variant parameter and
any number of numerical strings representing Snowflake IDs, and prints out
the time, worker ID, process ID, and increment corresponding to each provided
Snowflake.

If the variant is not specified, the Snowflake is parsed and displayed for
every supported variant.

Current supported variants are Twitter and Discord.

```
~/gocode/bin$ ./melter -variant=Discord 381898139189116930
381898139189116930 Discord: 2017-11-19 20:06:51.707 Z, w 0, p 0, i 2
~/gocode/bin$
```
