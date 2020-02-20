# Win32.NastyBoi
malware project for CPRE331

It's a little worm that propagates via a CnC server that launches EternalBlue. The worm itself scans the internet for targets and sends them to the CnC server.

For the EternalBlue exploit, AutoBlue is used:
https://github.com/3ndG4me/AutoBlue-MS17-010
