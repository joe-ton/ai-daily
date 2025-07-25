#include <iostream>

std::string getNumber(std::string size) {
    std::string firstTwo = size.substr(0, 2);
    return firstTwo;
}

int main() {
    std::string size;

    std::cout << "Size: ";
    std::cin >> size;

    auto result = getNumber(size);

    std::cout << "Size typed: " << result << std::endl;
}
