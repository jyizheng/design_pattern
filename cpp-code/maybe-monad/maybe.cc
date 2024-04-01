#include <iostream>
#include <optional>

std::optional<int> doubleValue(std::optional<int> val) {
    if (val.has_value()) {
        return val.value() * 2;
    } else {
        return std::nullopt;
    }
}

int main() {
    std::optional<int> maybeFive = 5;
    auto maybeTen = doubleValue(maybeFive);

    if (maybeTen.has_value()) {
        std::cout << "Optional holds value: " << maybeTen.value() << std::endl;
    } else {
        std::cout << "Optional holds no value." << std::endl;
    }

    return 0;
}


