
# Simple speedtest.net scraper

```
$ make get_vendor
$ make run-dev
```

Output format (csv):
- short timestamp (please add 1167606000 to get the real time)
- ping time (ms)
- download speed (MB/s)
- upload speed (MB/s)
- ISP

```
$ cat 2007-04.output
8142000;999:0.49;0.04;TELEKOM
8142000;520:1.12;0.23;NETIA
8142000;849:0.86;0.38;HUGHESNET
8141940;26:1.28;0.21;OPTUS
8142000;65:2.13;0.72;SPECTRUM
8142000;63:2.60;0.35;XFINITY
8141940;74:1.10;1.11;JSC RUSCENTROSVYAZ
8142000;18:5.10;0.65;AT&T INTERNET
8142000;26:9.59;0.97;UPC
8142000;79:3.12;0.10;MIDCO
8142000;82:7.93;0.83;SPECTRUM
8142000;99:0.98;0.21;OPTUS
8142060;30:4.87;0.37;TWC (NOW SPECTRUM)
8142060;28:3.30;1.43;VERIZON FIOS
8142000;60:0.74;0.44;TELUS
8141940;199:0.31;0.35;XFINITY
8142000;465:0.19;0.16;ORION CYBER INTERNET
8141940;69:2.84;0.29;AT&T INTERNET
...
```

:warning: It should take at least 1y and tens of GB, to scrape the entire speedtest.net database :trollface:
