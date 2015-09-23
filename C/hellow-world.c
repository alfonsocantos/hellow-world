/*Hello World program*/

#include <stdio.h>

#define DEFAULT "World"
int main (int argc, char *argv[]){
	
	char* string;
	
	if (argc <= 1){
    		string = DEFAULT;
	} else {
	    	string = argv[1];
	}
	
	printf("hello %s!\n", string);
	return 0;
}
