# Exercise 02

## Code Explation

The server basically acts like the linux `cat` command, it sends the received message back and to ensure there were no IO interference the code had no logs.
The client was using the `Dial` api which works using the same methods in both protocols and the code again had no logs besides the time log.
A simple bash script was written in order to make testing easier with both servers.

## Results

All data was parsed to ms before any further operation and the data was manipulated using python

### TCP

| Clients | Mean | Standard Deviation |
| ------- |:----:| ------------------:|
| 1       | 1.39 | 0.22               |
| 2       | 1.54 | 0.21               |
| 3       | 1.73 | 0.15               |
| 4       | 1.89 | 0.21               |
| 5       | 2.26 | 0.44               |

### UDP

| Clients | Mean | Standard Deviation |
| ------- |:----:| ------------------:|
| 1       | 1.17 | 0.21               |
| 2       | 1.18 | 0.11               |
| 3       | 1.26 | 0.13               |
| 4       | 1.35 | 0.15               |
| 5       | 1.48 | 0.32               |
