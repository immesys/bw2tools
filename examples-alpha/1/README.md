# Example 1

This example shows simple publish / subscribe usage of BOSSWAVE. Note that
for the current version of BOSSWAVE, the chain building procedure will not
contact other routers (to avoid spamming other people). This can be a
problem if your local router is actually unaware of all the DOTs needed
to build a chain. To resolve this, please run these commands before
trying the demo (and after starting the local router):

```
bw2 inspect ex1.key
bw2 buildchain -u "castle.bw2.io/bwtools/*" -x "PC*" -t E0ZqjH7CL1KNWRhS27uLREaz-beuUDs6qIXo0hZTQT4=
```
