#include <string>
#include <iostream>

int main() {
    std::string name = "Bob";
    std::string *pName = &name;

    std::cout << *pName << std::endl;

    return 0;
}
