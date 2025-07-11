#include <iostream>

int main() {
    std::string name = "Joe";

    std::string *pName = &name;

    std::cout << *pName;
}
