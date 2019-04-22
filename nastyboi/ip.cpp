#pragma once
#include <string>
#include <time.h>
#include <iostream>
#include "stdafx.h"

using std::string;
using std::to_string;

string getiplocal(unsigned int target){
	string localnet = "192.168.1.";
	string targetstring = to_string(target);
	string ret = localnet + targetstring;
	return ret;
}


string getiprand(){
	srand(time(NULL));
	int first = rand() % 255;
	int second = rand() % 255;
	int third = rand() % 255;
	int fourth = rand() % 255;
	string ip = to_string(first) + "." + to_string(second) + "." + to_string(third) + "." + to_string(fourth);
	return ip;
}