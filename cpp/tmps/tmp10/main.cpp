#include <iostream>

std::string getValue(std::string size) {
    std::string firstTwo = size.substr(0, 2);
    return firstTwo;
}

int main() {
    std::string size;

    std::cout << "Size: ";
    std::cin >> size;

    auto result = getValue(size);

    std::cout << "Size Returned: " << result << std::endl; 
}
