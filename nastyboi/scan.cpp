#include "stdafx.h"
#include <WS2tcpip.h>
#pragma comment(lib, "ws2_32.lib")

using namespace std;

int scan(string ip) {//returns 0 on successful scan

	int remoteport = 445;

	SOCKET sock = socket(AF_INET, SOCK_STREAM, 0);
	if (sock == INVALID_SOCKET) {
		cerr << "Can't create socket, " << WSAGetLastError() << endl;
		return 1;
	}

	//string ipaddr = getiplocal(target);
	cout << "Trying to connect to " + ip << endl;

	sockaddr_in hint;
	hint.sin_family = AF_INET;
	hint.sin_port = htons(remoteport);
	inet_pton(AF_INET, ip.c_str(), &hint.sin_addr);


	int connResult = connect(sock, (sockaddr*)&hint, sizeof(hint));
	if (connResult == SOCKET_ERROR) {
		cerr << "Could not connect: " << WSAGetLastError() << endl;
		closesocket(sock);
		return 1;
	}
	cout << "successfully connected to " + ip << endl;
	closesocket(sock);
	return 0;
}