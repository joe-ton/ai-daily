#include <iostream>

int main() {
    std::string name = "Joe";

    for (int i = 0; i < name.size(); i++) {
        std::cout << name[i] << std::endl;
    }

    for (char c : name) {
        std::printf("Hello, World\n");
    }

    return 0;
}
