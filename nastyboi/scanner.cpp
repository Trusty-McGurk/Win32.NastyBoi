#include <icmpapi.h>
#include <WinSock2.h>
#include <iphlpapi.h>
#include "stdafx.h"
#include "getip.h"

int scan(){
    WSAData wsadata;
    WORD DllVersion = MAKEWORD(2, 1);
    if(WSAStartu(DllVersion, &wsadata) != 0){
        std::cout << "BIG FAIL" << std::endl;
    }

    int i;
    for(i = 0; i < 255; i++){
        //if trying to scan the host's ip, continue
        SOCKADDR_IN addr;
        int addrlen = sizeof(addr);
        addr.sin_family = AF_INET;
        addr.sin_port = htons(445);
        addr.sin_addr.s_addr = inet_addr("127.0.0.1");//replace with ip to scan

        SOCKET sock = socket(AF_INET, SOCK_STREAM, NULL);
        int succ = bind(sock, (SOCKADDR*)&addr, sizeof(addr));
        if(succ){
            //send ip to main thread for exploitation
        }

        close(sock);
    }
}